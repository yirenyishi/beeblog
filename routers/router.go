package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/map", &controllers.MapController{})
	beego.Router("/file/upload", &controllers.FileController{}, "post:Upload")
	beego.Include(&controllers.PageController{})
}
