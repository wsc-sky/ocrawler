package parser

import (
	"ocrawler/engine"
	"regexp"
)

const CityRe = `<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>	`


func ParseCity(contents []byte) engine.ParserResult{
	re := regexp.MustCompile(CityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User: "+name)
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				ParserFunc: func(c []byte) engine.ParserResult {
					return ParseProfile(c, name)
				},
			})
	}
	return result
}
