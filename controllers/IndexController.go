package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/service"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	timeBlog, _ := service.IndexBlogs(12, 0)
	this.Data["TimeBlog"] = timeBlog
	browsBlog, _ := service.IndexBlogs(12, 1)
	this.Data["BrowsBlog"] = browsBlog
	likeBlog, _ := service.IndexBlogs(12, 2)
	this.Data["LikeBlog"] = likeBlog
	commentBlog, _ := service.IndexBlogs(12, 3)
	this.Data["CommBlog"] = commentBlog

	this.Data["UserId"] = this.GetSession("userid")
	this.Data["HeadImg"] = this.GetSession("headimg")
	this.Data["NickName"] = this.GetSession("nickname")
	this.Data["IsLogin"] = this.GetSession("nickname") != nil
	this.Data["IsHome"] = true
	this.TplName = "index.html"
	return
}
