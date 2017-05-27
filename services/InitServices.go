package services

import "github.com/astaxie/beego/orm"
import _ "github.com/mattn/go-sqlite3"

var (
	OrmerS   orm.Ormer
	ForwardS = new(ForwardService)
)

func init() {

	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	//数据库连接
	orm.RegisterDataBase("default", "sqlite3", "data/data.db")
	//开启DEBUG模式，输出SQL信息
	orm.Debug = true

	OrmerS = orm.NewOrm()
	OrmerS.Using("default")

}
