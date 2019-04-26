package controllers

import (
	"myproject/models"
	"time"

	"path"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
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

// HandleLogin supper post method
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

// ShowLogin supper get methods
func (c *MainController) ShowLogin() {
	c.TplName = "login.html"
}

// ShowAdd get method
func (c *MainController) ShowAdd() {
	c.TplName = "add.html"
}

// HandleAdd supper post method
func (c *MainController) HandleAdd() {
	artiName := c.GetString("articleName")
	artiContent := c.GetString("content")
	artiImg := c.GetString("uploadname")

	logs.Info(artiName, artiContent, artiImg)
	// c.TplName = ""
	f, h, err := c.GetFile("uploadname")
	// upload file
	filext := path.Ext(h.Filename)
	logs.Info(filext)
	if strings.ToLower(filext) != ".jpg" && strings.ToLower(filext) != ".png" {
		logs.Info("file type err")
		return
	}
	// logs.Info(h.Size)
	if h.Size > 5*1024*1024 {
		logs.Info("file size too big")
		return
	}

	filename := time.Now().Format("2006-01-02 15:04:05") + filext

	defer f.Close()
	if err != nil {
		logs.Info("upload file failed")
		return
	}

	if err = c.SaveToFile("uploadname", "./static/img/"+strings.Replace(strings.Replace(filename, " ", "-", 1), ":", "-", 2)); err != nil {
		logs.Info("file save failed", err)
		return
	}

	if artiContent == "" || artiName == "" {
		logs.Info("add content error")
		return
	}

	o := orm.NewOrm()
	arti := models.Article{}
	arti.ArtiName = artiName
	arti.Acontent = artiContent
	arti.Aimg = "/static/img" + h.Filename
	logs.Info(arti)
	if _, err = o.Insert(&arti); err != nil {
		logs.Info("insert failed", err)
	}

	// c.Redirect("/index", 302)
	c.TplName = "index.html"
}

// ShowIndex get method
func (c *MainController) ShowIndex() {
	o := orm.NewOrm()
	var articles []models.Article
	if _, err := o.QueryTable("Article").All(&articles); err != nil {
		logs.Info(err)
		return
	}
	c.Data["articles"] = articles
	c.TplName = "index.html"

}

// ShowContent ...
func (c *MainController) ShowContent() {
	id, err := c.GetInt("id")

	if err != nil {
		logs.Info("get info by id", err)
		return
	}
	o := orm.NewOrm()

	arti := models.Article{Id: id}
	if err = o.Read(&arti); err != nil {
		logs.Info("read failed", err)
		return
	}

	c.Data["article"] = arti

	c.TplName = "content.html"
}
