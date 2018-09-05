package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct {}

//func (e SimpleEngine) Run(seeds ...Request) {
//	var requests []Request
//	//从种子中拿到所有的Request
//	for _, r := range seeds {
//		requests = append(requests, r)
//	}
//	//遍历种子去进行解析
//	for len(requests) > 0 {
//		r := requests[0]
//		requests = requests[1:]
//
//		//去获取这个url的内容，失败则忽略
//		parseResult, err := e.worker(r)
//		if err != nil {
//			continue
//		}
//		//把所有的新的request添加到request中
//		requests = append(requests, parseResult.Request...)
//
//		for _, item := range parseResult.Items {
//			log.Printf("Got item %v", item)
//		}
//	}
//}


func  worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParseFunc(body), nil
}