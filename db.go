package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var (
	schema = []string{
		"CREATE TABLE `hashes`(`id` Integer NOT NULL PRIMARY KEY, `hash` Text NOT NULL, `divider` Real NOT NULL, CONSTRAINT `unique_id` UNIQUE (`id`), CONSTRAINT `unique_hash_divider` UNIQUE (`hash` COLLATE NOCASE, `divider`))",
		"CREATE UNIQUE INDEX `hash_divider` ON `hashes`(`hash` COLLATE NOCASE, `divider`)",
		"CREATE TABLE `pools`(`id` Integer NOT NULL PRIMARY KEY AUTOINCREMENT, `host` Text NOT NULL COLLATE NOCASE, `port` Integer NOT NULL, `hash_id` Integer NOT NULL, CONSTRAINT `unique_id` UNIQUE (`id`), CONSTRAINT `unique_host_port` UNIQUE (`host` COLLATE NOCASE, `port`), CONSTRAINT `lnk_hashes_pools` FOREIGN KEY (`hash_id`) REFERENCES `hashes`(`id`) ON DELETE Cascade ON UPDATE Cascade)",
		"CREATE UNIQUE INDEX `host_port` ON `pools`(`host` COLLATE NOCASE, `port`)",
		"CREATE TABLE `users`(`name` Text NOT NULL PRIMARY KEY COLLATE NOCASE, `pool_id` Integer NOT NULL, `user` Text NOT NULL, `password` Text NOT NULL, CONSTRAINT `unique_name` UNIQUE (`name` COLLATE NOCASE), CONSTRAINT `unique_pool_user_password` UNIQUE (`pool_id`, `user` COLLATE NOCASE, `password` COLLATE NOCASE), CONSTRAINT `lnk_pools_users` FOREIGN KEY (`pool_id`) REFERENCES `pools`(`id`) ON DELETE Cascade ON UPDATE Cascade)",
		"CREATE UNIQUE INDEX `pool_user_password` ON `users`(`pool_id`, `user` COLLATE NOCASE, `password` COLLATE NOCASE)",
	}
	data = []string{
		"INSERT INTO `hashes` (`hash`, `divider`) VALUES ('sha256', 1.0)",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('cn.ss.btc.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('cn.ss.btc.com', 1800, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('121.29.19.24', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sz.ss.btc.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sz.ss.btc.com', 1800, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('us.ss.btc.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('us.ss.btc.com', 1800, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('eu.ss.btc.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('eu.ss.btc.com', 1800, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc.f2pool.com', 1314, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc.f2pool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc-us.f2pool.com', 1314, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc-us.f2pool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc-eu.f2pool.com', 1314, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc-eu.f2pool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc-bj.ss.poolin.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc-bj.ss.poolin.com', 1883, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc-va.ss.poolin.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc-va.ss.poolin.com', 1883, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('stratum.antpool.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('stratum.antpool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc.viabtc.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc.viabtc.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('stratum.poolhb.com', 8888, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bak.poolhb.com', 8888, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('hk.huobipool.com', 8888, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bn.huobipool.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bn.huobipool.com', 1800, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bn.huobipool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bs.huobipool.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bs.huobipool.com', 1800, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bs.huobipool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bm.huobipool.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bm.huobipool.com', 1800, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bm.huobipool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bu.huobipool.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bu.huobipool.com', 1800, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bu.huobipool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('cn.stratum.slushpool.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('cn.stratum.slushpool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('cn02.stratum.slushpool.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('cn02.stratum.slushpool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('cn03.stratum.slushpool.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('cn03.stratum.slushpool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('us-east.stratum.slushpool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('eu.stratum.slushpool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('ca.stratum.slushpool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sg.stratum.slushpool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('jp.stratum.slushpool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.hk.nicehash.com', 3334, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.usa.nicehash.com', 3334, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.eu.nicehash.com', 3334, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.jp.nicehash.com', 3334, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.in.nicehash.com', 3334, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.br.nicehash.com', 3334, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('cn.emcd.io', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('us.emcd.io', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('eu.emcd.io', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('eu2.emcd.io', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('gate.emcd.io', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('ir.emcd.io', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.sea.mine.zpool.ca', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.na.mine.zpool.ca', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.eu.mine.zpool.ca', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.jp.mine.zpool.ca', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
		"INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('mining.dev.pool.titan.io', 4242, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'))",
	}
)

/*
Db - the database of proxy.
*/
type Db struct {
	handle *sql.DB
}

/*
Init - initializing of the database. Validating of the existence of the database.
If the database is not exist we are creating a database and filling it of default values.
If the database is exist and correct we are opening connect to it.

@return bool if initializing successfull.
*/
func (d *Db) Init() bool {
	//LogInfo("proxy : check database file on path %s", "", dbPath)
	_, err := os.Stat(dbPath)
	if os.IsPermission(err) {
		LogError("proxy : access denied to %s", "", dbPath)
		return false
	}
	if os.IsNotExist(err) {
		//LogInfo("proxy : database file not exist", "")
		d.Create()
	}

	//LogInfo("proxy : opening database file on path %s", "", dbPath)
	d.handle, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		LogError("proxy : error opening database on path %s: %s", "", dbPath, err.Error())
		return false
	}

	return true
}

/*
Create - the creating the database and filling it of default values.

@return error
*/
func (d *Db) Create() error {
	//LogInfo("proxy : creating database on path %s", "", dbPath)
	td, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		LogError("proxy : error creating database on path %s: %s", "", dbPath, err.Error())
		return err
	}
	defer td.Close()
	t, err := td.Begin()
	if err != nil {
		LogError("proxy : error starting transaction: %s", "", err.Error())
		return err
	}
	//LogInfo("proxy : creating schema", "")
	for _, v := range schema {
		if _, err := td.Exec(v); err != nil {
			LogError("proxy : error executing query %s: %s", "", v, err.Error())
			t.Rollback()
			return err
		}
	}
	//LogInfo("proxy : loading default data", "")
	for _, v := range data {
		if _, err := td.Exec(v); err != nil {
			LogError("proxy : error executing query %s: %s", "", v, err.Error())
			t.Rollback()
			return err
		}
	}
	t.Commit()

	return nil
}

/*
Close - the closing of the database.
*/
func (d *Db) Close() {
	//LogInfo("proxy : closing database file on path %s", "", dbPath)
	_ = d.handle.Close()
}

/*
AddUser - the addding of the user to the database.

@param *User user pointer to User struct.

@return error
*/
func (d *Db) AddUser(user *User) error {
	name := user.GetName()

	if !ValidateHexString(name) {
		return fmt.Errorf("invalid format user.name = %s", name)
	}

	id, err := d.GetPool(user.pool)
	if err != nil {
		return err
	}
	if id == 0 {
		return fmt.Errorf("pool %s not found", user.pool)
	}
	if _, err := d.handle.Exec(
		"INSERT INTO `users` (`name`, `pool_id`, `user`, `password`) VALUES ($1, $2, $3, $4);",
		user.name,
		id,
		user.user,
		user.password,
	); err != nil {
		return err
	}

	return nil
}

/*
GetUser - the getting of the user from the database.

@param string name user name.

@return *User the pointer to the founded user.
        error
*/
func (d *Db) GetUser(name string) (*User, error) {
	var host string
	var port int

	if !ValidateHexString(name) {
		return nil, fmt.Errorf("invalid format name = %s", name)
	}
	user := new(User)

	row := d.handle.QueryRow("SELECT `u`.`name`, `p`.`host`, `p`.`port`, `u`.`user`, `u`.`password`, `h`.`hash`, `h`.`divider` FROM `users` AS `u` INNER JOIN `pools` AS `p` ON `u`.`pool_id` = `p`.`id` INNER JOIN `hashes` AS `h` ON `p`.`hash_id` = `h`.`id` WHERE `u`.`name` = $1;", name)
	err := row.Scan(&user.name, &host, &port, &user.user, &user.password, &user.hash, &user.divider)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user with name = %s not found", name)
	}
	if err != nil {
		return nil, err
	}
	user.pool = host + ":" + strconv.Itoa(port)

	return user, nil
}

/*
GetUserByPool - the getting of the user by the pool and the user name of the pool.

@param string pool the user pool in format addr:port.
@param string user the name of the user.

@return *User the pointer to the founded user.
        error
*/
func (d *Db) GetUserByPool(pool string, user string) (*User, error) {
	id, err := d.GetPool(pool)
	if err != nil {
		return nil, err
	}
	if id == 0 {
		return nil, errors.New("pool not found")
	}
	if user == "" {
		return nil, errors.New("empty user")
	}

	us := User{pool: pool, user: user}

	row := d.handle.QueryRow("SELECT `u`.`name`, `u`.`password`, `h`.`hash`, `h`.`divider` FROM `users` AS `u` INNER JOIN `pools` AS `p` ON `u`.`pool_id` = `p`.`id` INNER JOIN `hashes` AS `h` ON `p`.`hash_id` = `h`.`id` WHERE `p`.`id` = $1 AND `u`.`user` = $2;", id, user)
	err = row.Scan(&us.name, &us.password, &us.hash, &us.divider)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &us, nil
}

/*
GetPool - the getting of the pool by his data.

@param string pool the pool of the user in format addr:port.

@return uint64 the identifier of the pool.
        error
*/
func (d *Db) GetPool(pool string) (uint64, error) {
	var id uint64

	if !ValidateAddr(pool, true) {
		return 0, fmt.Errorf("invalid format pool = %s", pool)
	}
	parts := strings.Split(pool, ":")
	port, _ := strconv.Atoi(parts[1])

	row := d.handle.QueryRow("SELECT `id` FROM `pools` WHERE `host` = $1 AND `port` = $2;", parts[0], port)
	err := row.Scan(&id)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return id, nil
}
