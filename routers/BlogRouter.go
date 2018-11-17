package routers

import (
	"github.com/astaxie/beego"
	"beeblog/controllers"
)

func init() {
	beego.Router("/blog/save", &controllers.BlogController{}, "get:Save")
	beego.Router("/blog/:id([0-9]+)", &controllers.BlogController{}, "get:Get")
}
