package engine

import "myReptile/concurrent/entity"
import "myReptile/concurrent/fetcher"

func Worker(request entity.Request) (entity.ParseResult, error) {
	//请求
	body, err := fetcher.Fetch(request.Url)
	if nil != err {
		return entity.ParseResult{}, err
	}
	//解析
	return request.ParseFunc(body), nil
}
