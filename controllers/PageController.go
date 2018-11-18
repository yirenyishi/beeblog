package controllers

import "github.com/astaxie/beego"

type PageController struct {
	beego.Controller
}

// @router /iframe/blog [get]
func (this *PageController) Blog() {
	this.TplName = "iframe/blog.html"
}

// @router /note [get]
func (this *PageController) Note() {
	this.Data["IsNote"] = true
	this.TplName = "note.html"
}
