package utils

import (
	"io"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"bufio"
	"log"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/net/html/charset"
)

func HtmlParser(r *bufio.Reader) ([]byte, error){
	e := DetermineEncoding(r)
	utf8Reader := transform.NewReader(r, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err!= nil {
		return nil, err
	}
	return all, nil
}


func DetermineEncoding(r io.Reader) encoding.Encoding{
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
