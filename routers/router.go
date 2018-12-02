package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/map", &controllers.MapController{})
	beego.Router("/iframe/note", &controllers.PageController{},"get:IframeNote")
	beego.Router("/iframe/user", &controllers.PageController{},"get:IframeUser")
	beego.Router("/iframe/blog", &controllers.PageController{},"get:Blog")
	beego.Router("/file/upload", &controllers.FileController{}, "post:Upload")
	beego.Router("/himg/upload", &controllers.FileController{}, "post:HeadImgUpload")
	beego.Include(&controllers.PageController{})
}
