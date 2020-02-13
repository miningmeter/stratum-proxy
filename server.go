/*
Stratum-прокси с внешним управлением.
*/

package main

import (
	"flag"
	"log"
	"net"
	"os"
	"regexp"
	"time"

	rpc2 "github.com/miningmeter/rpc2"
	"github.com/miningmeter/rpc2/stratumrpc"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"net/http"
	_ "net/http/pprof"
)

/*
VERSION - версия прокси.
*/
const VERSION = "0.01"

var (
	// Интерфейс обработки команд от майнера и пула.
	mining Mining
	// Воркеры.
	workers Workers
	// База данных.
	db Db
	// Stratum хост и порт.
	stratumAddr = "127.0.0.1:9332"
	// API хост и порт.
	webAddr = "127.0.0.1:8080"
	// Вывод в syslog.
	syslog = false
	// GitCommit - Git commit for build
	GitCommit string
	// Скомпилированная регулярка для проверки шестнадцатиричных строк.
	rHexStr = regexp.MustCompile(`^[\da-fA-F]+$`)
	// Расширения, поддерживаемые прокси.
	sExtensions = []string{
		"subscribe-extranonce",
		"version-rolling",
	}
	// Путь к базе данных SQLite.
	dbPath = "proxy.db"
)

/*
Главная функция.
*/
func main() {
	flag.StringVar(&stratumAddr, "stratum.addr", "127.0.0.1:9332", "Address and port for stratum")
	flag.StringVar(&webAddr, "web.addr", "127.0.0.1:8080", "Address and port for web server and metrics")
	flag.BoolVar(&syslog, "syslog", false, "On true adapt log to out in syslog, hide date and colors")
	flag.StringVar(&dbPath, "db.path", "proxy.db", "Filepath for SQLite database")
	flag.Parse()

	if syslog {
		log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	}
	LogInfo("proxy : version: %s-%s", "", VERSION, GitCommit)

	// Инициализируем базу данных.
	if !db.Init() {
		os.Exit(1)
	}
	defer db.Close()
	// Инициализируем внутренние хранилища.
	workers.Init()

	// Инициализируем API и метрики.
	LogInfo("proxy : web server serve on: %s", "", webAddr)
	// Пользователи.
	http.Handle("/api/v1/users", &API{})
	// Метрики.
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(webAddr, nil)

	InitMinerServer()

	os.Exit(0)
}

/*
InitMinerServer - инициализация сервера для подключения майнеров.
*/
func InitMinerServer() {
	// Запускаем JSON-RPC сервер.
	server := rpc2.NewServer()
	// Подписываем его на необходимые события.
	server.Handle("mining.subscribe", mining.Subscribe)
	server.Handle("mining.authorize", mining.Authorize)
	server.Handle("mining.submit", mining.Submit)
	server.Handle("mining.extranonce.subscribe", mining.ExtranonceSubscribe)
	server.Handle("mining.configure", mining.Configure)

	server.OnDisconnect(Disconnect)

	LogInfo("proxy : listen on: %s", "", stratumAddr)

	// Цикл ожидания подключений.
	link, _ := net.Listen("tcp", stratumAddr)
	for {
		// Ждем подключения.
		conn, err := link.Accept()
		if err != nil {
			LogError("proxy : accept error: %s", "", err.Error())
			break
		}

		go WaitWorker(conn, server)
	}
}

/*
WaitWorker - ожидание инициализации воркера.

@param net.Conn     conn   - соединение.
@param *rpc2.Server server - сервер.
*/
func WaitWorker(conn net.Conn, server *rpc2.Server) {
	addr := conn.RemoteAddr().String()
	LogInfo("%s : try connect to proxy", "", addr)
	// Инициализируем воркер.
	w := &Worker{addr: addr}
	// Связываем JSON-RPC соединение с воркером.
	state := rpc2.NewState()
	state.Set("worker", w)
	// Запускаем обработчик соединения в отдельном потоке.
	go server.ServeCodecWithState(stratumrpc.NewStratumCodec(conn), state)
	// Ждем 3 секунды инициализации воркера, которая начнется при отправке воркером команд.
	<-time.After(3 * time.Second)
	// Если воркер не инициализировался - убиваем соединение.
	if w.GetID() == "" {
		LogInfo("%s : disconnect by silence", "", addr)
		conn.Close()
	}
}

/*
Connect - обработка подключения воркера к прокси.

@param *rpc2.Client client указатель на подключенный клиент
@param *Worker w указатель на подключаемый воркер
*/
func Connect(client *rpc2.Client, w *Worker) {
	wAddr := w.GetAddr()
	if err := w.Init(client); err == nil {
		sID := w.GetID()
		LogInfo("%s : connect to proxy", sID, wAddr)
	} else {
		LogError("%s : error connect to proxy: %s", "", wAddr, err.Error())
		client.Close()
	}
}

/*
Disconnect - обработка отключения майнера от прокси.

@param *rpc2.Client client указатель на отключаемый клиент
*/
func Disconnect(client *rpc2.Client) {
	temp, _ := client.State.Get("worker")
	w := temp.(*Worker)
	w.Disconnect()
}
