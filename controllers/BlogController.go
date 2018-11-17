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
	blog := &models.Blog{Title: "ELK+logback+kafska+nginx 搭建分布式日志分析平台"}
	err := service.SaveBlog(blog)
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
	//this.Data["IsHome"] = true
	this.TplName = "blog.html"
}
