package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/like/:id([0-9]+)", &controllers.LikeController{}, "get:Save")
	beego.Router("/api/unlike/:id([0-9]+)", &controllers.LikeController{}, "get:Delete")
}
