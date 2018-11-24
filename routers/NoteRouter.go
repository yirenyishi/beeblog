package routers

import (
	"github.com/astaxie/beego"
	"beeblog/controllers"
)

func init() {
	beego.Router("/note/:id([0-9]+)", &controllers.NoteController{}, "get:Get")
	beego.Router("/note", &controllers.NoteController{}, "get:Note")
	beego.Router("/notecoll/save", &controllers.NoteController{}, "post:SaveNoteColl")
	beego.Router("/notecoll/edit", &controllers.NoteController{}, "post:EditNoteColl")
	beego.Router("/note/save", &controllers.NoteController{}, "post:Save")
	beego.Router("/note/edit/:id([0-9]+)", &controllers.NoteController{}, "post:Edit")
	beego.Router("/note/del/:id([0-9]+)", &controllers.NoteController{}, "post:Delete")
	beego.Router("/notecol/del/:id([0-9]+)", &controllers.NoteController{}, "post:DelNoteColl")
}
