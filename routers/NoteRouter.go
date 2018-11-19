package routers

import (
	"github.com/astaxie/beego"
	"beeblog/controllers"
)

func init() {
	beego.Router("/note/:id([0-9]+)", &controllers.NoteController{}, "get:Get")
	beego.Router("/notecoll/save", &controllers.NoteController{}, "post:SaveNoteColl")
	beego.Router("/note/save", &controllers.NoteController{}, "post:Save")
}
