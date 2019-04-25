package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
	// "github.com/astaxie/beego/orm"
	// "myproject/models"
	// "fmt"
)

// MainController beego route controller
type MainController struct {
	beego.Controller
}

// Get route
func (c *MainController) Get() {
	// o := orm.NewOrm()
	// user := models.User{}
	// user.Name = "eilinge"
	// user.Pwd = "tester"

	// if _, err := o.Insert(&user); err != nil {
	// 	// beego.Info("error", err)
	// 	panic(err)
	// 	// return
	// }
	
	// o := orm.NewOrm()
	// user := models.User{}
	// user.Name = "eilinge"
	// user.Pwd = "tester"
	
	// if err := o.Read(&user, "Name", "Pwd"); err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("select successfully : %v", user)

	// o := orm.NewOrm()
	// user := models.User{}
	// user.Name = "eilinge"
	// user.Pwd = "tester"
	
	// if err := o.Read(&user, "Name", "Pwd"); err == nil {
	// 	user.Name = "eilin"
	// 	user.Pwd = "duzi"
	// 	if _, err = o.Update(&user); err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("update successfully : %v", user)
	// }

	// o := orm.NewOrm()
	// user := models.User{}
	// user.Id =1
	
	// if _, err := o.Delete(&user); err == nil {
	// 	panic(err)
	// }
	

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
