package main

import (
	"ocrawler/crawler/engine"
	"runtime"
	"ocrawler/crawler/persist"
	"ocrawler/crawler/config"
	"ocrawler/crawler/zhenai/constant"
	"ocrawler/crawler/zhenai/parser"
	"ocrawler/crawler/scheduler"
)

// to start distribute crawler you need to set up:
// 1. elastic search at port:9000
// 2. go run main.go
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	itemChan, err := persist.ItemSaver(config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 500,
		ItemChan: itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url:        constant.Xian,
		Parser: engine.NewFuncParser(parser.ParseCity, config.ParseCity),
	})

}






