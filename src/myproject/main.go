package main

import (
	_ "myproject/routers"

	"github.com/astaxie/beego"
)

// add view func

func main() {
	beego.AddFuncMap("PrePage", GetPrePage)
	beego.AddFuncMap("NextPage", GetNextPage)
	beego.Run()
	//透明static
}

// GetPrePage ...
func GetPrePage(data int) (preindex int) {
	preindex = data - 1
	return
}

// GetNextPage ...
func GetNextPage(data int) (nextindex int) {
	nextindex = data + 1
	return
}
