package scheduler

import (
	"myReptile/entity"
)

type Scheduler interface {
	//请求提交方法
	Submit(entity.Request)
	//内部通道配置
	ConfigWorkerChannel(chan entity.Request)
}
