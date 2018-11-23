package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/service"
)

type PageController struct {
	beego.Controller
}

// @router /iframe/blog [get]
func (this *PageController) Blog() {
	cats, err := service.GetCats()
	if err != nil {
		this.Redirect("500.html", 302)
		return
	}
	this.Data["Cats"] = cats
	this.TplName = "iframe/blog.html"
}

// @router /iframe/note [get]
func (this *PageController) IframeNote() {
	this.TplName = "iframe/note.html"
}
