package controllers

import "github.com/astaxie/beego"

type MapController struct {
	beego.Controller
}

func (this *MapController) Get() {
	this.Data["IsMap"] = true
	this.TplName = "map.html"
}
