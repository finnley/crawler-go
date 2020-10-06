package parser

import (
	"crawler-go/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	//现在只需要提取城市和url,改造如下
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	//对于每一个url生成一个新的request
	result := engine.ParseResult{}
	for _, m := range matches {
		//fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}

	return result
}
