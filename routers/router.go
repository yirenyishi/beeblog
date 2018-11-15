package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/login", &controllers.UserController{}, "get:LoginPage")
	beego.Router("/", &controllers.IndexController{})
	//beego.Router("/login", &controllers.UserController{}, "post:Login")

}
