package routers

import (
	"github.com/astaxie/beego"
	"beeblog/controllers"
)

func init() {
	beego.Router("/comms/save", &controllers.CommentController{}, "post:Save")
	beego.Router("/comms/del/:id([0-9]+)", &controllers.CommentController{}, "get:Del")
}