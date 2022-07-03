package parser

import (
	"encoding/json"
	"fmt"
	"myReptile/concurrent/51jobs/dao"
	"myReptile/concurrent/51jobs/jsonEntity"
	entity "myReptile/concurrent/entity"
	"regexp"
)

//岗位列表解析器
func ParseJobList(contents []byte) entity.ParseResult {
	re := regexp.MustCompile(`window.__SEARCH_RESULT__ = (.*?)</script>`)

	matches := re.FindAllSubmatch(contents, -1)
	result := entity.ParseResult{}
	db := dao.ConnectMysql()
	defer db.Close()

	for _, match := range matches {
		//岗位列表（本次解析的结果）
		var jobListResult jsonEntity.JobListResult
		//解析json
		err := json.Unmarshal(match[1], &jobListResult)
		if err != nil {
			fmt.Println(err.Error())
		}

		for _, item := range jobListResult.EngineJds {
			dao.AddRecord(db, item)
			result.Items = append(result.Items, item)
			result.Requests = append(result.Requests, entity.Request{
				//岗位详情url
				Url: item.JobHref,
				//岗位详情解析器
				ParseFunc: entity.NilParser,
			})
		}

	}
	return result

}
