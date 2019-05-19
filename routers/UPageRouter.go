package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/search", &controllers.UPageController{}, "get:SearchPage")
}
