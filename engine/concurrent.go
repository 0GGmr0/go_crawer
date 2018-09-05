package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ... Request) {
	//读入的url
	in := make(chan Request)
	//读出的数据
	out := make(chan ParseResult)
	//建立通道的意思吧
	e.Scheduler.ConfigureMasterWorkerChan(in)
	//把读入的url送入这个调度地
	for i := 0; i < e.WorkCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	//接收worker输出的结果
	for {
		//从通道拿到数据
		result := <-out
		for _, item := range result.Items {
			fmt.Printf("Got item: %v", item)
		}

		for _, request := range result.Request {
			e.Scheduler.Submit(request)
		}

	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result

		}
	}()
}