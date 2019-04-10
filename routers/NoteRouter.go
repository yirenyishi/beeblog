package routers

import (
	"github.com/astaxie/beego"
	"beeblog/controllers"
)

func init() {
	beego.Router("/api/note/:id([0-9]+)", &controllers.NoteController{}, "get:Get")
	beego.Router("/note", &controllers.NoteController{}, "get:Note")
	beego.Router("/api/notecoll/save", &controllers.NoteController{}, "post:SaveNoteColl")
	beego.Router("/api/notecoll/edit", &controllers.NoteController{}, "post:EditNoteColl")
	beego.Router("/api/note/save", &controllers.NoteController{}, "post:Save")
	beego.Router("/api/note/edit/:id([0-9]+)", &controllers.NoteController{}, "post:Edit")
	beego.Router("/api/note/del/:id([0-9]+)", &controllers.NoteController{}, "post:Delete")
	beego.Router("/api/notecol/del/:id([0-9]+)", &controllers.NoteController{}, "post:DelNoteColl")
}
