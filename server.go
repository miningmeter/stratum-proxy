/*
Stratum-proxy with external manage.
*/

package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"regexp"
	"time"

	"github.com/ambelovsky/gosf"
	ws "github.com/ambelovsky/gosf"
	"github.com/joho/godotenv"
	rpc2 "github.com/miningmeter/rpc2"
	"github.com/miningmeter/rpc2/stratumrpc"
	"gitlab.com/TitanInd/hashrouter/contractmanager"
	"gitlab.com/TitanInd/hashrouter/events"
	"gitlab.com/TitanInd/hashrouter/interfaces"
)

/*
VERSION - proxy version.
*/
const VERSION = "0.01"

var (
	// Processing commangds from worker and pool.
	mining Mining
	// Workers.
	workers Workers
	// Db of users credentials.
	db Db
	// Stratum endpoint.
	stratumAddr = "127.0.0.1:9332"
	// API endpoint.
	webAddr = "127.0.0.1:8080"
	// Pool target
	poolAddr = ""
	// Out to syslog.
	syslog = false
	// GitCommit - Git commit for build
	GitCommit string
	// Compiled regexp for hexademical checks.
	rHexStr = regexp.MustCompile(`^[\da-fA-F]+$`)
	// Extensions that supported by the proxy.
	sExtensions = []string{
		"subscribe-extranonce",
		"version-rolling",
	}
	// SQLite db path.
	dbPath = "proxy.db"
	// Metrics proxy tag.
	tag = ""
	// HashrateContract Address
	hashrateContract string
	// Eth node Address
	ethNodeAddr string
)

func init() {
	flag.StringVar(&stratumAddr, "stratum.addr", "127.0.0.1:9332", "Address and port for stratum")
	flag.StringVar(&webAddr, "web.addr", "127.0.0.1:8080", "Address and port for web server and metrics")
	flag.StringVar(&poolAddr, "pool.addr", "mining.dev.pool.titan.io:4242", "Address and port for mining pool")
	flag.BoolVar(&syslog, "syslog", false, "On true adapt log to out in syslog, hide date and colors")
	flag.StringVar(&dbPath, "db.path", "proxy.db", "Filepath for SQLite database")
	flag.StringVar(&tag, "metrics.tag", stratumAddr, "Prometheus metrics proxy tag")
	flag.StringVar(&hashrateContract, "contract.addr", "", "Address of smart contract that node is servicing")
	flag.StringVar(&ethNodeAddr, "ethNode.addr", "", "Address of Ethereum RPC node to connect to via websocket")
}

/*
Main function.
*/
func main() {

	flag.Parse()

	if syslog {
		log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	}

	// var addr = flag.String("addr", "localhost:8080", "http service address")

	godotenv.Load(".env")

	LogInfo("proxy : version: %s-%s", "", VERSION, GitCommit)

	// Initializing of database.
	// if !db.Init() {
	// 	os.Exit(1)
	// }
	// defer db.Close()
	// Inintializing of internal storage.
	workers.Init(poolAddr, os.Getenv("DEFAULT_POOL_USER"), os.Getenv("DEFAULT_POOL_PASSWORD"))

	// Initializing of API and metrics.
	// LogInfo("proxy : web server serve on: %s", "", webAddr)
	// // Users.
	// http.Handle("/api/v1/users", &API{})
	// // Metrics.
	// http.Handle("/metrics", promhttp.Handler())
	// go http.ListenAndServe(webAddr, nil)

	ws.Startup(map[string]interface{}{"port": 9999})

	eventManager := events.NewEventManager()

	InitContractManager(eventManager, hashrateContract, ethNodeAddr)

	InitWorkerServer(poolAddr)

	os.Exit(0)
}

type ConnectionInfo struct {
	IpAddress     string
	Status        string
	SocketAddress string
	Total         string
	Accepted      string
	Rejected      string
}

func InitSocket(in chan ConnectionInfo) {
	ws.Listen("ws", func(client *ws.Client, request *ws.Request) *ws.Message {
		return ws.NewSuccessMessage()
	})

	for connection := range in {
		ws.Broadcast("", "", &gosf.Message{
			Body: map[string]interface{}{
				"type":        "cxns",
				"connections": []ConnectionInfo{connection},
			},
		})
	}
}

type DestinationUpdateHandler struct{ interfaces.Subscriber }

func (d *DestinationUpdateHandler) Update(message interface{}) {
	destinationMessage := message.(contractmanager.Dest)

	newPoolAddr := destinationMessage.NetUrl

	oldUser := workers.user
	oldPass := workers.password

	//LogInfo("Switching to new pool address: %v", "", newPoolAddr)

	workers.Init(newPoolAddr, os.Getenv("TEST_POOL_USER"), os.Getenv("TEST_POOL_PASSWORD"))

	workers.Reset()

	<-time.After(180 * time.Second)

	log.Printf("Switching back to old pool address: %v", poolAddr)
	workers.Init(poolAddr, oldUser, oldPass)
}

func InitContractManager(eventManager interfaces.IEventManager, hashrateContract string, ethNodeAddr string) {

	ctx := context.Background()

	sellerManager := &contractmanager.SellerContractManager{}
	sellerManager.SetLogger(log.Default())

	handler := &DestinationUpdateHandler{}
	eventManager.Attach(contractmanager.DestMsg, handler)

	contractmanager.Run(&ctx, sellerManager, eventManager, hashrateContract, ethNodeAddr)
}

/*
InitWorkerServer - initializing of server for workers connects.
*/
func InitWorkerServer(poolAddr string) {
	// Launching of JSON-RPC server.
	server := rpc2.NewServer()
	// Subscribing of server to needed handlers.
	server.Handle("mining.subscribe", mining.Subscribe)
	server.Handle("mining.authorize", mining.Authorize)
	server.Handle("mining.submit", mining.Submit)
	server.Handle("mining.extranonce.subscribe", mining.ExtranonceSubscribe)
	server.Handle("mining.configure", mining.Configure)

	server.OnDisconnect(Disconnect)

	LogInfo("proxy : listen on: %s", "", stratumAddr)

	// Waiting of connections.
	link, _ := net.Listen("tcp", stratumAddr)
	for {
		conn, err := link.Accept()
		if err != nil {
			LogError("proxy : accept error: %s", "", err.Error())
			break
		}

		go WaitWorker(conn, server)
	}
}

/*
WaitWorker - waiting of worker init.

@param net.Conn     conn   - connection.
@param *rpc2.Server server - server.
*/
func WaitWorker(conn net.Conn, server *rpc2.Server) {
	addr := conn.RemoteAddr().String()
	//LogInfo("%s : try connect to proxy", "", addr)
	// Initializing of worker.
	w := &Worker{addr: addr}
	// Linking of JSON-RPC connection to worker.
	state := rpc2.NewState()
	state.Set("worker", w)
	// Running of connection handler in goroutine.
	go server.ServeCodecWithState(stratumrpc.NewStratumCodec(conn), state)
	// Waiting 3 seconds of worker initializing, which will begin when the worker sends the commands.
	<-time.After(3 * time.Second)
	// If worker not initialized, we kill connection.
	if w.GetID() == "" {
		//LogInfo("%s : disconnect by silence", "", addr)
		conn.Close()
	}
}

/*
Connect - processing of connecting worker to proxy.

@param *rpc2.Client client pointer to connecting client
@param *Worker w pointer to connecting worker
*/
func Connect(client *rpc2.Client, w *Worker) {
	wAddr := w.GetAddr()
	if err := w.Init(client); err == nil {
		// sID := w.GetID()
		//LogInfo("%s : connect to proxy", sID, wAddr)
	} else {
		LogError("%s : error connect to proxy: %s", "", wAddr, err.Error())
		client.Close()
	}
}

/*
Disconnect - processing of disconnecting worker to proxy.

@param *rpc2.Client client pointer to disconnecting client
*/
func Disconnect(client *rpc2.Client) {
	temp, _ := client.State.Get("worker")
	w := temp.(*Worker)
	w.Disconnect()
}
