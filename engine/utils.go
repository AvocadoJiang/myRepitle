package engine

import (
	entity2 "myReptile/entity"
	"myReptile/fetcher"
)

func Worker(request entity2.Request) (entity2.ParseResult, error) {
	//请求
	body, err := fetcher.Fetch(request.Url)
	if nil != err {
		return entity2.ParseResult{}, err
	}
	//解析
	return request.ParseFunc(body), nil
}
