package controllers

import (
	"fmt"
	"math"
	"myproject/models"
	"strconv"
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
	c.TplName = "register.html"
}

// Post supper post method
func (c *MainController) Post() {
	// get data
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	// validate data
	if userName == "" || pwd == "" {
		logs.Info("no data")
		c.Redirect("/register", 302)
		return
	}
	c.Redirect("/login", 302)
}

// ShowLogin supper get methods
func (c *MainController) ShowLogin() {
	c.TplName = "login.html"
}

// HandleLogin supper post method
func (c *MainController) HandleLogin() {
	// get data
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	// validate data
	if userName == "" || pwd == "" {
		logs.Info("no data")
		c.Redirect("/login", 302)
		return
	}

	c.Redirect("/index", 302)
}

// ShowAdd get method
func (c *MainController) ShowAdd() {
	o := orm.NewOrm()
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)
	c.Data["types"] = types
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

	// 上传文件至myproject中
	if err = c.SaveToFile("uploadname", "./static/img/"+strings.Replace(strings.Replace(filename, " ", "-", 1), ":", "-", 2)); err != nil {
		logs.Info("file save failed", err)
		return
	}

	if artiContent == "" || artiName == "" {
		logs.Info("add content error")
		return
	}

	// store data to database
	o := orm.NewOrm()
	arti := models.Article{}
	arti.ArtiName = artiName
	arti.Acontent = artiContent
	arti.Aimg = "/static/img/" + h.Filename
	logs.Info(arti)

	typeName := c.GetString("select")
	if typeName == "" {
		logs.Info("select data error")
		return
	}

	var artiType models.ArticleType
	artiType.TypeName = typeName

	err = o.Read(&artiType, "TypeName")
	if err != nil {
		logs.Info("get type error", err)
		return
	}
	arti.ArticleTypes = &artiType

	if _, err = o.Insert(&arti); err != nil {
		logs.Info("insert failed", err)
	}

	c.Redirect("/index", 302)
}

// ShowIndex get method
func (c *MainController) ShowIndex() {
	o := orm.NewOrm()
	var articles []models.Article
	qs := o.QueryTable("Article")
	// qs.
	// if _, err := qs.All(&articles); err != nil {
	// 	logs.Info(err)
	// 	return
	// }

	count, err := qs.Count()

	if err != nil {
		logs.Info("get count", err)
		return
	}
	c.Data["count"] = count
	countPage := 2
	// int64(3/2) 1
	// math.Ceil(1.2) 2
	// math.Floor(1.2) 1
	pageCount := math.Ceil(float64(count) / float64(countPage))
	pageIndex := 1

	// <a href="/index?id=1">first</a>
	pageIn := c.GetString("id")
	pageIndex, _ = strconv.Atoi(pageIn)

	if err != nil {
		logs.Info("get id failed", err)
		return
	}
	countStart := countPage * (pageIndex - 1)
	if _, err = qs.Limit(countPage, countStart).All(&articles); err != nil {
		logs.Info(err)
		return
	}

	firstPage := false
	if pageIndex == 1 {
		firstPage = true
	}

	lastPage := false
	if pageIndex == int(pageCount) {
		lastPage = true
	}

	artype := orm.NewOrm()
	var types []models.ArticleType
	_, err = artype.QueryTable("ArticleType").All(&types)
	if err != nil {
		logs.Info("check articles type failed", err)
		return
	}

	c.Data["types"] = types
	c.Data["pageCount"] = pageCount
	c.Data["articles"] = articles
	c.Data["current"] = pageIndex
	c.Data["firstPage"] = firstPage
	c.Data["lastPage"] = lastPage
	// 进行数据传输
	c.TplName = "index.html"

}

// HandleIndex ...
func (c *MainController) HandleIndex() {
	typeName := c.GetString("select")
	if typeName == "" {
		logs.Info("select data failed")
		return
	}

	o := orm.NewOrm()
	var articles []models.ArticleType
	// Filter() 过滤器: where
	// Filter("table_field", fieldValue)
	// Orm 多表查询是惰性删除, 需指定外键关联表: RelatedSel
	o.QueryTable("Article").RelatedSel("ArticleType").Filter("ArticleType__TypeName", typeName).All(&articles)
	logs.Info(articles)
	// c.Redirect("/index", 302)
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

// ShowUpdate ...
func (c *MainController) ShowUpdate() {
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
	c.TplName = "update.html"
}

// HandleUpdate ...
func (c *MainController) HandleUpdate() {
	articleName := c.GetString("articleName")
	content := c.GetString("content")

	f, h, err := c.GetFile("uploadname")

	if err != nil {
		logs.Info("update file is not nill", err)
		return
	}
	// upload file
	filext := path.Ext(h.Filename)

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

	// update img to myproject

	if err = c.SaveToFile("uploadname", "./static/img/"+strings.Replace(strings.Replace(filename, " ", "-", 1), ":", "-", 2)); err != nil {
		logs.Info("file save failed", err)
		return
	}

	var id int
	if id, err = c.GetInt("id"); err != nil {
		logs.Info("get id failed", err)
		return
	}

	o := orm.NewOrm()
	arti := models.Article{Id: id}

	err = o.Read(&arti)
	if err != nil {
		logs.Info("Read data failed")
	}

	arti.ArtiName = articleName
	arti.Acontent = content
	arti.Aimg = "/static/img/" + h.Filename
	if _, err = o.Update(&arti, "ArtiName", "Acontent", "Aimg"); err != nil {
		logs.Info("update failed", err)
		return
	}
	fmt.Printf("update successfully : %v", arti)
	c.Redirect("/index", 302)
}

// ShowDelete ...
func (c *MainController) ShowDelete() {
	var id int
	id, err := c.GetInt("id")
	if err != nil {
		logs.Info("get id failed", err)
		return
	}
	o := orm.NewOrm()
	arti := models.Article{Id: id}

	err = o.Read(&arti)
	if err != nil {
		logs.Info("Read data failed")
	}

	_, err = o.Delete(&arti)
	if err != nil {
		panic(err)
	}

	c.Redirect("/index", 302)
}

// ShowaddType ...
func (c *MainController) ShowaddType() {
	o := orm.NewOrm()
	var addTypes []models.ArticleType
	// addTypes = models.ArticleType()
	_, err := o.QueryTable("ArticleType").All(&addTypes)

	if err != nil {
		logs.Info("read failed")
		return
	}
	logs.Info("addTypes", addTypes)
	c.Data["addTypes"] = addTypes
	c.TplName = "addType.html"
}

// HandleaddType ...
func (c *MainController) HandleaddType() {

	typeName := c.GetString("typeName")
	if typeName == "" {
		logs.Info("typename is empty")
		return
	}

	o := orm.NewOrm()
	typena := models.ArticleType{}
	typena.TypeName = typeName
	_, err := o.Insert(&typena)
	if err != nil {
		logs.Info("insert failed", err)
		return
	}
	c.Redirect("/addType", 302)
}

//ShowdeleteType ...
func (c *MainController) ShowdeleteType() {
	id, err := c.GetInt("id")
	if err != nil {
		logs.Info("get type id failed", err)
		return
	}
	deletype := orm.NewOrm()
	dp := models.ArticleType{Id: id}
	_, err = deletype.Delete(&dp)
	if err != nil {
		logs.Info("delete type failed", err)
		return
	}
	c.Redirect("/addType", 302)
}
