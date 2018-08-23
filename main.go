package main

import (
	"ocrawler/engine"
	"ocrawler/zhenai/constant"
	"ocrawler/zhenai/parser"
	"ocrawler/scheduler"
	"runtime"
	"ocrawler/persist"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: persist.ItemSaver(),
	}

	//e.Run(engine.Request{
	//	Url:        constant.CityList,
	//	ParserFunc: parser.ParseCityList,
	//})

	e.Run(engine.Request{
		Url:        constant.Shanghai,
		ParserFunc: parser.ParseCity,
	})

}






