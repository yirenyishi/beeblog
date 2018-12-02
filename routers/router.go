package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/map", &controllers.MapController{})
	beego.Router("/file/upload", &controllers.FileController{}, "post:Upload")
	beego.Router("/himg/upload", &controllers.FileController{}, "post:HeadImgUpload")
	beego.Include(&controllers.PageController{})
}
