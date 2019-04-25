package routers

import (
	"myproject/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.MainController{})
	// 实现了自定义的get/post请求, 请求将不会访问默认方法
	// beego.Router("/login", &controllers.MainController{}, "get:ShowLogin") // 假使点击"submit"(Post), 页面将无法访问
	beego.Router("/login", &controllers.MainController{}, "get:ShowLogin;post:HandleLogin")
}
