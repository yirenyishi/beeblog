package controllers

import (
	"beeblog/models"
	"beeblog/service"
	"github.com/astaxie/beego"
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
	if err == nil {
		this.Data["json"] = blog
	} else {
		this.Data["json"] = models.ReurnError(500,"保存失败")
	}
	this.ServeJSON()
	return
}

func (this *BlogController) Get() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	blog, err := service.GetBlog(id)
	if err == nil {
		this.Data["Blog"] = blog
	}
	this.TplName = "blog.html"
}

func (this *BlogController) New() {
	this.TplName = "newblog.html"
}
func (this *BlogController) Blog1() {
	this.TplName = "blog1.html"
}

func (this *BlogController) BlogsPage() {
	blogs,_ := service.FindBlogs()
	this.Data["Blogs"] = blogs
	this.Data["IsBlog"] = true
	this.TplName = "blogs.html"

}