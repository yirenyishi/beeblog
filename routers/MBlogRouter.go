package routers

import (
	"beeblog/mcontrollers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/blogs", &mcontrollers.MBlogController{}, "get:BlogsPage")
	beego.Router("/api/blog/:id([0-9]+)", &mcontrollers.MBlogController{}, "get:Get")
}
