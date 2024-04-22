# LachesisNode

### build 

```
/.../LachesisNode # go build -o build/lachesisnode ./lachesis/main.go
```

### customizable flag

```

		Name:  "datadir",
		Usage: "Directory for the configuration",
		Value: defaultDataDir(),

		Name:  "node_addr",
		Usage: "IP:Port to bind Lachesis",
		Value: "127.0.0.1:1337",


		Name:  "no_client",
		Usage: "Run Lachesis with dummy in-memory App client",


		Name:  "proxy_addr",
		Usage: "IP:Port to bind Proxy Server",
		Value: "127.0.0.1:1338",


		Name:  "client_addr",
		Usage: "IP:Port of Client App",
		Value: "127.0.0.1:1339",


		Name:  "service_addr",
		Usage: "IP:Port of HTTP Service",
		Value: "127.0.0.1:8000",


		Name:  "log_level",
		Usage: "debug, info, warn, error, fatal, panic",
		Value: "debug",


		Name:  "heartbeat",
		Usage: "Heartbeat timer milliseconds (time between gossips)",
		Value: 1000,


		Name:  "max_pool",
		Usage: "Max number of pooled connections",
		Value: 2,


		Name:  "tcp_timeout",
		Usage: "TCP timeout milliseconds",
		Value: 1000,


		Name:  "cache_size",
		Usage: "Number of items in LRU caches",
		Value: 500,


		Name:  "sync_limit",
		Usage: "Max number of events for sync",
		Value: 1000,


		Name:  "store",
		Usage: "badger, inmem",
		Value: "badger",


		Name:  "store_path",
		Usage: "File containing the store database",
		Value: defaultBadgerDir(),

```



### http service

```
/stats
/participants/
/event/
/lasteventfrom/
/events/
/consensusevents/
/round/
/lastround/
/roundwitnesses/
/roundevents/
/root/
/block/
```



### run

./lachesisnode run [--flag "value"]

### transaction
todo: add api to add transaction manually
