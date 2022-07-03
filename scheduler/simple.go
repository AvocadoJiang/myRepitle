package scheduler

import (
	"myReptile/entity"
)

//简单调度器
//多个goroutine公用一个channel

//简单调度器结构
type SimpleScheduler struct {
	WorkerChannel chan entity.Request
}

//请求提交
func (s *SimpleScheduler) Submit(request entity.Request) {
	s.WorkerChannel <- request
}

//内部配置
func (s *SimpleScheduler) ConfigWorkerChannel(in chan entity.Request) {
	s.WorkerChannel = in
}
