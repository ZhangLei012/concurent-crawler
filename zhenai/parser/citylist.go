package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]+>([^<]+)</a>`

// ParseCityList is the function to parser the content fetched by fetcher
// and return cityList with ParserFunc
func ParseCityList(content []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	match := re.FindAllSubmatch(content, -1)

	result := engine.ParserResult{}
	for _, m := range match {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			URL:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
