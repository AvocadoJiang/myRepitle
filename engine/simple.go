package engine

import (
	"myReptile/entity"
	"myReptile/fetcher"
)

type SimpleEngine struct {
}

//总控模块
//单线程抓取
func (s *SimpleEngine) Run(requests []entity.Request) {
	//从任务中获取req
	for len(requests) > 0 {
		//开始处理第一个请求
		request := requests[0]
		requests = requests[1:]
		//调用fetcher,请求数据
		body, err := fetcher.Fetch(request.Url)
		if nil != err {
			continue
		}
		//解析
		parseResult := request.ParseFunc(body)
		//将解析结果中的请求添加到请求队列中
		requests = append(requests, parseResult.Requests...)
	}
}
