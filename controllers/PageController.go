package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/service"
	"beeblog/models"
)

type PageController struct {
	beego.Controller
}

// @router /iframe/blog [get]
func (this *PageController) Blog() {
	cats, err := service.GetCats()
	if err != nil {
		this.Redirect("/500", 302)
		return
	}
	this.Data["Cats"] = cats
	this.TplName = "iframe/blog.html"
}

// @router /iframe/user [get]
func (this *PageController) IframeUser() {
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["IsLogin"] = false
	} else {
		this.Data["IsLogin"] = true
		if user, err := service.GetUser(uid.(int64)); err == nil {
			this.Data["User"] = user
		} else {
			this.Data["User"] = &models.User{Id: uid.(int64)}
		}
	}
	this.TplName = "iframe/user.html"
	return
}

// @router /iframe/note [get]
func (this *PageController) IframeNote() {
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["IsLogin"] =  false
	}else {
		this.Data["IsLogin"] = true
		noteColls,err:=service.GetNoteColl(uid.(int64))
		if err== nil {
			this.Data["NoteColl"] = noteColls
		}
	}
	this.TplName = "iframe/note.html"
}

// @router /us
func (this *PageController) UsPage() {
	this.Data["IsUs"] = true
	this.Data["UserId"] = this.GetSession("userid")
	this.Data["HeadImg"] = this.GetSession("headimg")
	this.Data["NickName"] = this.GetSession("nickname")
	this.Data["IsLogin"] = this.GetSession("nickname") != nil
	this.TplName = "us.html"
}

// @router /404 [get]
func (this *PageController) PageNotFound() {
	this.TplName = "404.html"
}

// @router /500 [get]
func (this *PageController) ServerError() {
	this.TplName = "500.html"
}

// @router /403 [get]
func (this *PageController) ServerDemined() {
	this.TplName = "403.html"
}
