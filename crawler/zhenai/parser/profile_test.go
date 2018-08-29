package parser

import (
	"testing"
	"io/ioutil"
	"ocrawler/crawler/model"
	"ocrawler/crawler/engine"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := parseProfile(contents,
		"http://album.zhenai.com/u/108906739",
		"安静的雪")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element, but was %d", len(result.Items))
	}
	actual := result.Items[0]

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

	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}