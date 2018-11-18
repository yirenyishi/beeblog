package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
	"beeblog/service"
	"strconv"
)

type BlogController struct {
	beego.Controller
}

func (this *BlogController) Save() {
	title := this.GetString("title")
	blogHtml := this.GetString("blogHtml")
	catory := this.GetString("catory")
	catoryId, _ := strconv.ParseInt(catory, 10, 64)
	labels := this.GetStrings("labels[]")
	blog := &models.Blog{Title: title, BlogHtml: blogHtml, CategoryId: catoryId, UserId: 1}
	err := service.SaveBlog(blog, labels)
	if (err == nil) {
		this.Data["json"] = blog
	} else {
		this.Data["json"] = models.ReurnError("保存失败")
	}
	this.ServeJSON()
}

func (this *BlogController) Get() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	blog, err := service.GetBlog(id)
	if (err == nil) {
		this.Data["Blog"] = blog
	}
	this.TplName = "blog.html"
}

func (this *BlogController) New() {
	this.TplName = "newblog.html"
}
