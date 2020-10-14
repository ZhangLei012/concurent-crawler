package parser

import (
	"crawler/engine"
	"regexp"
)

const userListRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

//ParseCity is to parse the city page and get the local user's url and name
func ParseCity(content []byte) engine.ParserResult {
	re := regexp.MustCompile(userListRe)
	match := re.FindAllSubmatch(content, -1)
	result := engine.ParserResult{}
	for _, m := range match {
		name := string(m[2])
		result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engine.Request{
			URL: string(m[1]),
			ParserFunc: func(content []byte) engine.ParserResult {
				return ParseProfile(content, name)
			},
		})
	}
	return result
}
