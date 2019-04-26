package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// User ...
type User struct {
	Id   int
	Name string `orm:"unique"`
	Pwd  string
}

// Article ...
type Article struct {
	Id       int       `orm:"pk;auto"`
	ArtiName string    `orm:"size(20)"`
	Atime    time.Time `orm:"auto_now"`
	Acontent string
	Acount   int `orm:"default(0);null"`
	Aimg     string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:tester@tcp(127.0.0.1:3306)/test?charset=utf8")

	orm.RegisterModel(new(User), new(Article))

	orm.RunSyncdb("default", false, true)
	// true: 每次运行程序时, 重新创建表
	// orm.RunSyncdb("default", true, true)
}
