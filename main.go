package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
	"crawler/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		WorkCount: 10,
		Scheduler: &scheduler.SimpleScheduler{},

	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}


