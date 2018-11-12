package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.UserController{}, "Get:LoginPage")
	beego.Router("/login", &controllers.UserController{}, "post:Login")

}
