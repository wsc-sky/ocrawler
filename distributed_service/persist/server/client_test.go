package main

import (
	"testing"
	"ocrawler/distributed_service/rpcsupport"
	"ocrawler/crawler/engine"
	"ocrawler/crawler/model"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	// start ItemSaverServer
	go serveRpc(host, "test")

	time.Sleep(time.Second)
	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// call save
	item := engine.Item{
		Url: "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id: "108906739",
		Payload: model.Profile{
			Name:       "安静的雪",
			Gender:     "女",
			Age:        35,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Marriage:   "离异",
			Location:   "新疆阿勒泰",
			Education:  "大学本科",
			Occupation: "人事/行政",
			Xinzuo:     "牡羊座",
			Car:        "未购车",
			House: "已购房",
		}}
	var result string
	err = client.Call("ItemSaverService.Save", item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
