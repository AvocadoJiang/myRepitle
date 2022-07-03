package engine

import (
	"fmt"
	"myReptile/concurrent/entity"
	"myReptile/concurrent/scheduler"
)

type ConcurrentEngine struct {
	WorkerCount int //并发调用worker数量
	Scheduler   scheduler.Scheduler
}

//总控模块
func (e *ConcurrentEngine) Run(requests []entity.Request) {
	//负责传递请求
	in := make(chan entity.Request)
	//配置通大佛
	e.Scheduler.ConfigWorkerChannel(in)
	//负责返回解析结果
	out := make(chan entity.ParseResult)
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}
	for _, request := range requests {
		e.Scheduler.Submit(request)
	}
	for {
		//读取解析结果
		result := <-out
		for _, item := range result.Items {
			fmt.Println("item:%s\n", item)
		}
		//将获取到的新请求进行提交
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
	//接受out传递过来的结果
	//从任务中获取req
	for len(requests) > 0 {
		//开始处理第一个请求
		request := requests[0]
		requests = requests[1:]
		//调用worker
		parseResult, err := Worker(request)
		if err != nil {
			continue
		}
		//将解析结果中的请求添加到请求队列中
		requests = append(requests, parseResult.Requests...)
	}
}

//创建workder 并调用worker进行请求解析
func createWorker(in chan entity.Request, out chan entity.ParseResult) {
	go func() {
		for {
			//循环获取请求
			request := <-in
			result, err := Worker(request)
			if nil != err {
				continue
			}
			//结果写入
			out <- result
		}
	}()
}
