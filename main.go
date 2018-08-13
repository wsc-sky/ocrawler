package main

import (
	"ocrawler/engine"
	"ocrawler/zhenai/constant"
	"ocrawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        constant.CityList,
		ParserFunc: parser.ParseCityList,
	})

}






