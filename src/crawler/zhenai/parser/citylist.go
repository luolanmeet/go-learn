package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	limit := 10

	for _, m := range matchs {
		result.Items = append(result.Items, "City "+string(m[2]))
		request := engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		}
		result.Requests = append(result.Requests, request)

		limit--
		if limit <= 0 {
			break
		}
	}

	return result
}
