package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/service"
	"beeblog/models"
)

type PageController struct {
	beego.Controller
}

func (this *PageController) Blog() {
	catService := service.CategoryService{}
	cats, err := catService.GetCats()
	if err != nil {
		this.Redirect("/404", 302)
		return
	}
	this.Data["Cats"] = cats
	this.TplName = "iframe/blog.html"
}

func (this *PageController) IframeUser() {
	userService := service.UserService{}
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["IsLogin"] = false
	} else {
		this.Data["IsLogin"] = true
		if user, err := userService.GetUser(uid.(int64)); err == nil {
			this.Data["User"] = user
		} else {
			this.Data["User"] = &models.User{Id: uid.(int64)}
		}
	}
	this.TplName = "iframe/user.html"
	return
}

func (this *PageController) IframeNote() {
	noteService := service.NoteService{}
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["IsLogin"] =  false
		this.Data["NoteColl"] = []string{}
	}else {
		this.Data["IsLogin"] = true
		noteColls,err:=noteService.GetNoteColl(uid.(int64))
		if err== nil {
			this.Data["NoteColl"] = noteColls
		}
	}
	this.TplName = "iframe/note.html"
}

func (this *PageController) UsPage() {
	this.Data["IsUs"] = true
	this.Data["UserId"] = this.GetSession("userid")
	this.Data["HeadImg"] = this.GetSession("headimg")
	this.Data["NickName"] = this.GetSession("nickname")
	this.Data["IsLogin"] = this.GetSession("nickname") != nil
	this.TplName = "us.html"
}

func (this *PageController) PageNotFound() {
	this.TplName = "404.html"
}

func (this *PageController) ServerError() {
	this.TplName = "500.html"
}

func (this *PageController) ServerDemined() {
	this.TplName = "403.html"
}
