package parser

import (
	"crawler-go/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	//对于每一个url生成一个新的request
	result := engine.ParseResult{}
	for _, m := range matches {
		//fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
		name := string(m[2])
		result.Items = append(result.Items, "User " + name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}

	return result
}