package fetcher

import (
	"net/http"
	"fmt"
	"ocrawler/crawler/utils"
	"bufio"
	"time"
	"log"
	"ocrawler/crawler/config"
)

var rateLimiter = time.Tick(time.Second / config.Qps)
var  (
	reqHeaderKey = "User-Agent"
	reqHeaderValue = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.137 Safari/537.36 LBBROWSER"
)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	log.Printf("Fetching url %s", url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(reqHeaderKey, reqHeaderValue)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	contents, err := utils.HtmlParser(bodyReader)
	if err != nil {
		panic(err)
	}
	return contents, nil
}




