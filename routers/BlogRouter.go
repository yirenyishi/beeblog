package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/blog/new", &controllers.BlogController{}, "post:Save")
	beego.Router("/blog/:id([0-9]+)", &controllers.BlogController{}, "get:Get")
	beego.Router("/blog/new", &controllers.BlogController{}, "get:New")
	beego.Router("/blog1", &controllers.BlogController{}, "get:Blog1")
}
