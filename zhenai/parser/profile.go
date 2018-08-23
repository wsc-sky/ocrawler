package parser

import (
	"ocrawler/engine"
	"regexp"
	"strconv"
	"ocrawler/model"
)

var genderRe = regexp.MustCompile(
	`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var ageRe = regexp.MustCompile(
	`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe = regexp.MustCompile(
	`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightRe = regexp.MustCompile(
	`td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var incomeRe = regexp.MustCompile(
	`<td><span class="label">月收入：</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(
	`<td><span class="label">婚况：</span>([^<]+)</td>`)
var locationRe = regexp.MustCompile(
	`<td><span class="label">工作地：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(
	`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(
	`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(
	`td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(
	`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var guessRe = regexp.MustCompile(
	`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)

func ParseProfile(contents []byte, name string) engine.ParserResult {
	profile := model.Profile{Name: name}

	profile.Gender = extractString(contents, genderRe)
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}
	profile.Income = extractString(contents, incomeRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Location = extractString(contents, locationRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.Car = extractString(contents, carRe)

	result := engine.ParserResult{
		Items: []interface{}{profile},
	}

	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		name := string(m[2])
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: func(c []byte) engine.ParserResult {
					return ParseProfile(c, name)
				},
			})

	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string  {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}else{
		return ""
	}
}


