package parser

import (
	"testing"
	"io/ioutil"
	"ocrawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "真心找对象")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element, but was %d", len(result.Items))
	}
	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Name:       "真心找对象",
		Gender:     "女",
		Age:        24,
		Height:     158,
		Weight:     46,
		Income:     "3001-5000元",
		Marriage:   "未婚",
		Location:   "广东东莞",
		Education:  "中专",
		Occupation: "销售总监",
		Xinzuo:     "狮子座",
		Car:        "未购车",
	}

	if profile != expected {
		t.Errorf("expected %v, got %v", expected, profile)
	}
}