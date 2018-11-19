package controllers

import "github.com/astaxie/beego"

type PageController struct {
	beego.Controller
}

// @router /iframe/blog [get]
func (this *PageController) Blog() {
	this.TplName = "iframe/blog.html"
}

// @router /iframe/note [get]
func (this *PageController) IframeNote() {
	this.TplName = "iframe/note.html"
}
