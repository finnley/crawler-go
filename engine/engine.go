package engine

import (
	"crawler-go/fetcher"
	"log"
)

//seeds为种子页面
func Run(seeds ...Request)  {
	var requests []Request
	//将seeds放到request中
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		//将第一个requests取出来
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)
		//...将slice内容展开追加进去 parserFunc.Requests[0] parserFunc.Requests[1] ...
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %s", item)
		}
	}
}