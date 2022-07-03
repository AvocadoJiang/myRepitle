package main

import (
	"fmt"
	"myReptile/51jobs/parser"
	"myReptile/engine"
	"myReptile/entity"
	"myReptile/scheduler"
)

const seed = "https://search.51job.com/list/000000,000000,0000,00,9,99,Java,2,%d.html"

//https://search.51job.com/list/000000,000000,0000,00,9,99,Java,2,4.html?lang=c&postchannel=0000&workyear=99&cotype=99&degreefrom=99&jobterm=99&companysize=99&ord_field=0&dibiaoid=0&line=&welfare=
func main() {
	var requests []entity.Request
	//构建请求
	for i := 1; i <= 10; i++ {
		// 循环10次，只爬取前10页的数据，i为页码
		request := entity.Request{
			Url:       fmt.Sprintf(seed, i),
			ParseFunc: parser.ParseJobList,
		}
		requests = append(requests, request)
	}
	//并发爬取
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		//线程数
		WorkerCount: 10,
	}
	concurrentEngine.Run(requests)
	//单线程
	//simpleEngine := engine.SimpleEngine{}
	//simpleEngine.Run(requests)
}
