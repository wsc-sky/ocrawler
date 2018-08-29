package view

import (
	"os"
	"testing"
	"ocrawler/crawler/frontendfrontend/model"
	"ocrawler/crawler/enginer/engine"
	common "ocrawler/crawler/modeler/model"

)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView(
		"template.html")

	out, err := os.Create("template.test.html")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: common.Profile{
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
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		t.Error(err)
	}

	// TODO: verify contents in template.test.html
}
