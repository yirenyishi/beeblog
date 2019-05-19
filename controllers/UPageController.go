package controllers

import "github.com/astaxie/beego"

type UPageController struct {
	beego.Controller
}

func (this *UPageController) SearchPage() {
	this.TplName = "search.html"
}
