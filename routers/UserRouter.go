package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login", &controllers.UserController{}, "get:LoginPage")
	beego.Router("/logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/regist", &controllers.UserController{}, "get:RegistPage")
	beego.Router("/u/:id([0-9]+)", &controllers.UserController{}, "get:UserInfo")

	beego.Router("/api/user/edit", &controllers.UserController{}, "post:Edit")
	beego.Router("/api/login", &controllers.UserController{}, "post:Login")
	beego.Router("/api/regist", &controllers.UserController{}, "post:Regist")
	beego.Router("/me/blog", &controllers.UserController{}, "get:PersonBlog")
	beego.Router("/me/note", &controllers.UserController{}, "get:PersonNote")
	beego.Router("/me/like", &controllers.UserController{}, "get:PersonLike")
	beego.Router("/me/info", &controllers.UserController{}, "get:PersonInfo")
}
