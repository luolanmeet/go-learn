package parser

import (
	"crawler/engine"
	"regexp"
)

// 只有姓名
//const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
// 有姓名和性别
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>[^性别：]*性别：</span>([^</td>]*)</td>`

func ParseCity(contents []byte) engine.ParseResult {

	re := regexp.MustCompile(cityRe)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matchs {
		name := string(m[2])
		gender := string(m[3])
		result.Items = append(result.Items, "User "+name)
		request := engine.Request{
			Url: string(m[1]),
			// 通过函数式编程， 适配不兼容的接口
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name, gender)
			},
		}
		result.Requests = append(result.Requests, request)
	}

	return result
}
