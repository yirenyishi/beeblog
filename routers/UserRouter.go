package routers

import (
	"github.com/astaxie/beego"
	"beeblog/controllers"
)

func init() {
	beego.Router("/login", &controllers.UserController{}, "get:LoginPage")
	beego.Router("/login", &controllers.UserController{}, "post:Login")
	beego.Router("/regist", &controllers.UserController{}, "post:Regist")
}
