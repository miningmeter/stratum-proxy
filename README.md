# Stratum proxy
* Support for different mining algorithms through one port.
* Support for mining by several workers for 1 account.
* Counters of shares for each worker, user, pool and algorithm.
* Hash rate of each worker.
* Automatic detection of the mining algorithm for the correct calculation of the hashrate.
* Registration in a proxy through the API.
* Metrics in standard Prometheus format.

# Supported pools
The mining algorithm is automatically determined based on the pair <pool_host>:<pool_port>, so the proxy only supports connecting to a specific set of pools stored in the database. An API to expand the list of algorithms and pools is not yet available.

# Management REST API
The REST API is available at the proxy address `http://<web.addr>/api/v1` and now the API has only 1 command for registering credentials for connecting to the pool.

### POST /users
```json
{
    "pool": "<host>:<port>",
    "user": "<username>",
    "password": "<password>"
}
```
Correct answer:
```json
{
    "name": "<name>",
    "error": ""
}
```
The parameter `name` is used for identifing the connection and connects to the right pool with the right account. The proxy connection string will look like this:
```
-o stratum+tcp://<proxy_host>:<proxy_stratum_port> -u <name> -p <any, ignored>
```
Accounts are not deleted (in the future it is planned to do an automatic deletion after a period of inactivity).

# Available metrics
Metrics are available at `http://<web.addr>/metrics` and include a set of standard `Prometheus` metrics and custom metrics for monitoring the work of workers.

## List of custom metrics
* `proxy_worker_up{"proxy"="<proxy_host>:<proxy_port>", "worker"="<worker_host>:<worker_port>", "user"="<name>"}` - the status of the worker. Appears when the worker is connected to the proxy.
* `proxy_pool_up{"proxy"="<proxy_host>:<proxy_port>", "hash"="<hash>", "pool"="<pool_host>:<pool_port>"}` - the status of the pool. Appears when a proxy is connected to a pool.
* `proxy_worker_sended{"proxy"="<proxy_host>:<proxy_port>", "worker"="<worker_host>:<worker_port>", "user"="<name>", "hash"="<hash>", "pool"="<pool_host>:<pool_port>"}` - counter of the shares sent by the miner.
* `proxy_worker_accepted{"proxy"="<proxy_host>:<proxy_port>", "worker"="<worker_host>:<worker_port>", "user"="<name>", "hash"="<hash>", "pool"="<pool_host>:<pool_port>"}` - counter of the shares received by the pool.
* `proxy_worker_speed{"proxy"="<proxy_host>:<proxy_port>", "worker"="<worker_host>:<worker_port>", "user"="<name>", "hash"="<hash>", "pool"="<pool_host>:<pool_port>"}` - worker speed in hashes per second. Hashrate measurement window - 5 minutes, measurement interval - 1 minute.
* `proxy_worker_difficulty{"proxy"="<proxy_host>:<proxy_port>", "worker"="<worker_host>:<worker_port>", "user"="<name>", "hash"="<hash>", "pool"="<pool_host>:<pool_port>"}` - the difficulty set by the pool for the worker.
