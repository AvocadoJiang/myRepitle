package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"myReptile/51jobs/jsonEntity"
)

var (
	userName  string = "root"
	password  string = "123456"
	ipAddrees string = "127.0.0.1"
	port      int    = 3306
	dbName    string = "myReptile"
	charset   string = "utf8"
)

func ConnectMysql() *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)
	Db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	return Db
}

func AddRecord(Db *sqlx.DB, entity jsonEntity.EngineJds) {
	result, err := Db.Exec("insert into job_list  values(?,?,?,?,?,?,?,?,?,?,?,?)",
		entity.Jobid, entity.Coid, entity.JobHref, entity.JobName, entity.JobTitle, entity.CompanyHref, entity.CompanyName,
		entity.ProvidesalaryText, entity.Workarea, entity.WorkareaText, entity.Workyear, entity.Jobwelf)
	if err != nil {
		fmt.Printf("data insert faied, error:[%v]", err.Error())
		return
	}
	id, _ := result.LastInsertId()
	fmt.Printf("insert success, last id:[%d]\n", id)
}
