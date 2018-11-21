package controllers

import (
	"beeblog/models"
	"beeblog/service"
	"github.com/astaxie/beego"
	"strconv"
	"fmt"
)

type BlogController struct {
	beego.Controller
}

func (this *BlogController) Save() {
	uid := this.GetSession("userid")
	if uid == nil{
		this.Data["json"] = models.ReurnError(401, "")
		this.ServeJSON()
		return
	}
	title := this.GetString("title")
	blogHtml := this.GetString("blogHtml")
	catory := this.GetString("catory")
	catoryId, _ := strconv.ParseInt(catory, 10, 64)
	labels := this.GetStrings("labels[]")
	blog := &models.Blog{Title: title, BlogHtml: blogHtml, CategoryId: catoryId, UserId: 1}
	err := service.SaveBlog(blog, labels)
	if err == nil {
		this.Data["json"] = models.ReurnSuccess("")
	} else {
		this.Data["json"] = models.ReurnError(500, "保存失败")
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
	uid := this.GetSession("userid")
	if uid == nil {
		this.Redirect("/login.html",302)
		return
	}
	this.TplName = "newblog.html"
}
func (this *BlogController) Blog1() {
	this.TplName = "blog1.html"
}

func (this *BlogController) BlogsPage() {
	cats,errcat := service.GetCats()
	if errcat != nil {
		this.Redirect("500.html", 302)
		return
	}
	num, _ := this.GetInt("num")
	size, _ := this.GetInt("size")
	cat, _ := this.GetInt64("cat")
	flag, _ := this.GetInt("flag")
	if num <= 0 {
		num = 1
	}
	if size < 5 {
		size = 5
	}
	if cat <= 0 {
		cat = -1
	}
	fmt.Println("nelson page", num, size, cat)
	pages, err := service.FindBlogs(num, size, cat, flag)
	if err != nil {
		this.Redirect("500.html", 302)
		return
	}
	this.Data["Page"] = pages
	this.Data["Cats"] = cats
	this.Data["Cat"] = cat
	this.Data["Flag"] = flag
	this.Data["IsBlog"] = true
	this.TplName = "blogs.html"

}
