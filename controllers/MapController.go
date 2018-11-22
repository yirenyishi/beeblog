package controllers

import "github.com/astaxie/beego"

type MapController struct {
	beego.Controller
}

func (this *MapController) Get() {
	this.Data["NickName"] = this.GetSession("nickname")
	this.Data["IsLogin"] = this.GetSession("nickname") != nil
	this.Data["IsMap"] = true
	this.TplName = "map.html"
}
