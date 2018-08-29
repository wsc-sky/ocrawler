package engine

import "ocrawler/crawler/config"

type ParserFunc func(contents []byte, url string) ParseResult

// interface for Request Parser
type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

// new request to fetch data
type Request struct {
	Url string
	Parser Parser
}

// result of parsing
type ParseResult struct {
	Requests []Request
	Items []Item
}

// fetched data
type Item struct {
	Url 	string
	Id 		string
	Type 	string
	Payload interface{}
}

// temporary NilParser
type NilParser struct{}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return config.NilParser, nil
}

// implement the interface of common parser function
type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

// FuncParser factory
func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}



