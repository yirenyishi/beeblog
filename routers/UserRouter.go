package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login", &controllers.UserController{}, "get:LoginPage")
	beego.Router("/logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/login", &controllers.UserController{}, "post:Login")
	beego.Router("/regist", &controllers.UserController{}, "post:Regist")
	beego.Router("/regist", &controllers.UserController{}, "get:RegistPage")



	beego.Router("/me/blog", &controllers.UserController{}, "get:PersonBlog")

}
