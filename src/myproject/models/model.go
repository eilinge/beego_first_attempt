package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// User ...
type User struct {
	Id int
	Name string
	Pwd string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:tester@tcp(127.0.0.1:3306)/test?charset=utf8")

	orm.RegisterModel(new(User))

	orm.RunSyncdb("default", false, true)
}