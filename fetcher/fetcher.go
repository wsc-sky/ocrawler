package fetcher

import (
	"net/http"
	"fmt"
	"ocrawler/utils"
	"bufio"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
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


