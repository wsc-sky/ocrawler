package persist

import (
	"testing"
	"ocrawler/crawler/model"
	"github.com/olivere/elastic"
	"context"
	"encoding/json"
	"ocrawler/crawler/engine"
)

const index = "dating_test"

func TestSave(t *testing.T) {
	expected := engine.Item{
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

	// TODO: try to start up elastic search
	// here using docker go client
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	err = Save(client, expected, index)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("dating_profile").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, err := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("Got %v, expected %v",actual,expected)
	}

}