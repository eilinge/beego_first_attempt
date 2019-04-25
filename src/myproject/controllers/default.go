package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	// "github.com/astaxie/beego/orm"
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

	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"
	c.TplName = "register.html"
}

// Post supper post method
func (c *MainController) Post() {
	// get data
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	// logs.Info(userName, pwd)
	// validate data
	if userName == "" || pwd == "" {
		logs.Info("no data")
		c.Redirect("/register", 302)
		return
	}
	// store data ro database
	// o := orm.NewOrm()
	// user := models.User{}
	// user.Name = userName
	// user.Pwd = pwd

	// if _, err := o.Insert(&user); err != nil {
	// 	logs.Info("insert failed", err)
	// 	// panic(err)
	// 	return
	// }

	// return register.html
	c.Ctx.WriteString("register successfully")
}

// Post supper post method
func (c *MainController) HandleLogin() {
	// get data
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	// logs.Info(userName, pwd)
	// validate data
	if userName == "" || pwd == "" {
		logs.Info("no data")
		c.Redirect("/login", 302)
		return
	}
	// store data ro database
	// o := orm.NewOrm()
	// user := models.User{}
	// user.Name = userName

	// if _, err := o.Read(&user); err != nil {
	// 	logs.Info("Read failed", err)
	// 	// panic(err)
	// 	return
	// }

	// return login.html
	c.Ctx.WriteString("login successfully")
}

func (c *MainController) ShowLogin() {
	c.TplName = "login.html"
}
