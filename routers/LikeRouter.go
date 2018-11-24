package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/like/:id([0-9]+)", &controllers.LikeController{}, "get:Save")
	beego.Router("/unlike/:id([0-9]+)", &controllers.LikeController{}, "get:Delete")
}
