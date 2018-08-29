package main

import (
	"errors"
	"net/rpc"
	"log"
	"flag"
	"ocrawler/crawler/engine"
	"ocrawler/crawler/scheduler"
	"ocrawler/crawler/zhenai/parser"
	"ocrawler/distributed_service/rpcsupport"
	itemSaver "ocrawler/distributed_service/persist/client"
	"ocrawler/crawler/zhenai/constant"
	"ocrawler/crawler/config"
	worker "ocrawler/distributed_service/worker/client"
	"strings"
)

var (
	itemSaverHost = flag.String(
		"itemsaver_host", "", "itemsaver host")

	workerHosts = flag.String(
		"worker_hosts", "",
		"worker hosts (comma separated)")
)


// to start distribute crawler you need to set up:
// 1. elastic search at port:9000
// 2. fetch service: go run /worker/server/worker.go --port="your_port"
// 3. item saver service: go run /persist/server/itemsaver.go --port="your_port"
// 4. go run main.go --itemsaver_host=":1234" --worker_hosts=":9000,:9001,:9002,:9003"
func main() {
	flag.Parse()

	itemChan, err := itemSaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool, err := createClientPool(strings.Split(*workerHosts, ","))
	if err != nil {
		panic(err)
	}

	processor := worker.CreateProcessor(pool)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 200,
		ItemChan: itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:        constant.CityList,
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}

func createClientPool(hosts []string) (chan *rpc.Client, error) {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}

	if len(clients) == 0 {
		return nil, errors.New("no connections available")
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out, nil
}
