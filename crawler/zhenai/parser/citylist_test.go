package parser

import (
	"testing"
	"io/ioutil"
)

const (
	resultSize = 470
)

var expectedUrls = []string {
	"http://www.zhenai.com/zhenghun/aba",
	"http://www.zhenai.com/zhenghun/akesu",
	"http://www.zhenai.com/zhenghun/alashanmeng",
}
var expectedCities = []string {
	"City: 阿坝","City: 阿克苏", "City: 阿拉善盟",
}



func TestParserCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents, "")

	if len(result.Requests) != resultSize{
		t.Errorf("results should have %d requests; but had %d", resultSize, len(result.Requests, ))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}
}