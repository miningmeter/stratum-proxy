# Database.
## Scheme.
```sqlite
BEGIN;

-- CREATE TABLE "hashes" ---------------------------------------
CREATE TABLE "hashes"(
	"id" Integer NOT NULL PRIMARY KEY,
	"hash" Text NOT NULL,
	"divider" Real NOT NULL,
CONSTRAINT "unique_id" UNIQUE ( "id" ),
CONSTRAINT "unique_hash_divider" UNIQUE ( "hash" COLLATE NOCASE, "divider" ) );
-- -------------------------------------------------------------

-- CREATE INDEX "hash_divider" ---------------------------------
CREATE UNIQUE INDEX "hash_divider" ON "hashes"( "hash" COLLATE NOCASE, "divider" );
-- -------------------------------------------------------------

-- CREATE TABLE "pools" ----------------------------------------
CREATE TABLE "pools"(
	"id" Integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	"host" Text NOT NULL COLLATE NOCASE,
	"port" Integer NOT NULL,
	"hash_id" Integer NOT NULL,
CONSTRAINT "unique_id" UNIQUE ( "id" ),
CONSTRAINT "unique_host_port" UNIQUE ( "host" COLLATE NOCASE, "port" ),
CONSTRAINT "lnk_hashes_pools" FOREIGN KEY ( "hash_id" ) REFERENCES "hashes"( "id" ) ON DELETE Cascade ON UPDATE Cascade );
-- -------------------------------------------------------------

-- CREATE INDEX "host_port" ------------------------------------
CREATE UNIQUE INDEX "host_port" ON "pools"( "host" COLLATE NOCASE, "port" );
-- -------------------------------------------------------------

-- CREATE TABLE "users" ----------------------------------------
CREATE TABLE "users"(
	"name" Text NOT NULL PRIMARY KEY COLLATE NOCASE,
	"pool_id" Integer NOT NULL,
	"user" Text NOT NULL,
	"password" Text NOT NULL,
CONSTRAINT "unique_name" UNIQUE ( "name" COLLATE NOCASE ),
CONSTRAINT "unique_pool_user_password" UNIQUE ( "pool_id", "user" COLLATE NOCASE, "password" COLLATE NOCASE ),
CONSTRAINT "lnk_pools_users" FOREIGN KEY ( "pool_id" ) REFERENCES "pools"( "id" ) ON DELETE Cascade ON UPDATE Cascade );
-- -------------------------------------------------------------

-- CREATE INDEX "pool_user_password" ---------------------------
CREATE UNIQUE INDEX "pool_user_password" ON "users"( "pool_id", "user" COLLATE NOCASE, "password" COLLATE NOCASE );
-- -------------------------------------------------------------

COMMIT;
```
## Data.
```sqlite
BEGIN;

INSERT INTO `hashes` (`hash`, `divider`) VALUES ('sha256', 1.0);

INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('cn.ss.btc.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('cn.ss.btc.com', 1800, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('121.29.19.24', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sz.ss.btc.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sz.ss.btc.com', 1800, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('us.ss.btc.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('us.ss.btc.com', 1800, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('eu.ss.btc.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('eu.ss.btc.com', 1800, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc.f2pool.com', 1314, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc.f2pool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc-us.f2pool.com', 1314, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc-us.f2pool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc-eu.f2pool.com', 1314, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc-eu.f2pool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc-bj.ss.poolin.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc-bj.ss.poolin.com', 1883, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc-va.ss.poolin.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc-va.ss.poolin.com', 1883, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('stratum.antpool.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('stratum.antpool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc.viabtc.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('btc.viabtc.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('stratum.poolhb.com', 8888, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bak.poolhb.com', 8888, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('hk.huobipool.com', 8888, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bn.huobipool.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bn.huobipool.com', 1800, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bn.huobipool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bs.huobipool.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bs.huobipool.com', 1800, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bs.huobipool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bm.huobipool.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bm.huobipool.com', 1800, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bm.huobipool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bu.huobipool.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bu.huobipool.com', 1800, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('bu.huobipool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('cn.stratum.slushpool.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('cn.stratum.slushpool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('cn02.stratum.slushpool.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('cn02.stratum.slushpool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('cn03.stratum.slushpool.com', 443, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('cn03.stratum.slushpool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('us-east.stratum.slushpool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('eu.stratum.slushpool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('ca.stratum.slushpool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sg.stratum.slushpool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('jp.stratum.slushpool.com', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.hk.nicehash.com', 3334, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.usa.nicehash.com', 3334, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.eu.nicehash.com', 3334, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.jp.nicehash.com', 3334, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.in.nicehash.com', 3334, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.br.nicehash.com', 3334, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('cn.emcd.io', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('us.emcd.io', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('eu.emcd.io', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('eu2.emcd.io', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('gate.emcd.io', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('ir.emcd.io', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.sea.mine.zpool.ca', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.na.mine.zpool.ca', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.eu.mine.zpool.ca', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));
INSERT INTO `pools` (`host`, `port`, `hash_id`) VALUES ('sha256.jp.mine.zpool.ca', 3333, (SELECT `id` FROM `hashes` WHERE `hash` = 'sha256'));

COMMIT;
```
