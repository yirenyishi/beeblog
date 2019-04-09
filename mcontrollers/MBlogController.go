package mcontrollers

import (
	"beeblog/models"
	"beeblog/service"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type MBlogController struct {
	beego.Controller
}


func (this *MBlogController) Get() {
	blogService := service.BlogService{}
	userService := service.UserService{}
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	blog, err := blogService.GetBlog(id)
	if err != nil {
		this.Data["json"] = models.ReurnServerError(500)
		this.ServeJSON()
		return
	}
	this.Data["json"] = models.ReurnData("",blog)
	this.ServeJSON()
	userService.CountBrows(blog.UserId)
	blogService.EditBlogBrows(id)
	return
}


func (this *MBlogController) BlogsPage() {
	blogService := service.BlogService{}
	num, _ := this.GetInt("num")
	size, _ := this.GetInt("size")
	cat, _ := this.GetInt64("cat")
	flag, _ := this.GetInt("flag")
	if num <= 0 {
		num = 1
	}
	if size < 15 {
		size = 15
	}
	if cat <= 0 {
		cat = -1
	}
	pages, err := blogService.FindBlogs(num, size, cat, flag)
	if err != nil {
		this.Data["json"] = models.ReurnServerError(500)
		this.ServeJSON()
		return
	}
	this.Data["json"] = models.ReurnData("",pages)
	this.ServeJSON()
	return
}


func (this *MBlogController) EditPage() {
	blogService := service.BlogService{}
	uid := this.GetSession("userid")
	if uid == nil {
		this.Redirect("/login", 302)
		return
	}
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	blog, err := blogService.GetBlog(id)
	if err != nil {
		this.Redirect("/404", 302)
		return
	}
	if blog.UserId != uid.(int64) {
		this.Redirect("/403", 302)
		return
	}
	this.Data["Blog"] = blog
	this.TplName = "editblog.html"
}

func (this *MBlogController) Save() {
	blogService := service.BlogService{}
	userService := service.UserService{}
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["json"] = models.ReurnError(401, "")
		this.ServeJSON()
		return
	}
	title := this.GetString("title")
	blogHtml := this.GetString("blogHtml")
	catory := this.GetString("catory")
	catoryId, _ := strconv.ParseInt(catory, 10, 64)
	labels := this.GetStrings("labels[]")
	blog := &models.Blog{Title: title, BlogHtml: blogHtml, CategoryId: catoryId, UserId: uid.(int64)}
	err := blogService.SaveBlog(blog, labels)
	if err == nil {
		this.Data["json"] = models.ReurnData("",blog.Id)
	} else {
		this.Data["json"] = models.ReurnError(500, "保存失败")
	}
	this.ServeJSON()
	userService.CountBlog(uid.(int64))
	return
}

func (this *MBlogController) Edit() {
	userService := service.UserService{}
	blogService := service.BlogService{}
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["json"] = models.ReurnError(401, "")
		this.ServeJSON()
		return
	}
	id,_ := this.GetInt64("id")
	title := this.GetString("title")
	blogHtml := this.GetString("blogHtml")
	catory := this.GetString("catory")
	catoryId, _ := strconv.ParseInt(catory, 10, 64)
	labels := this.GetStrings("labels[]")
	blog,err :=blogService.GetBlog(id)
	if err != nil {
		this.Data["json"] = models.ReurnError(500, "保存失败")
		this.ServeJSON()
		return
	}
	blog.Title = title
	blog.BlogHtml = blogHtml
	blog.CategoryId = catoryId
	blog.Utime = time.Now()
	err = blogService.EditBlog(blog, labels)
	if err == nil {
		this.Data["json"] = models.ReurnSuccess("")
	} else {
		this.Data["json"] = models.ReurnError(500, "保存失败")
	}
	this.ServeJSON()
	userService.CountBlog(uid.(int64))
	return
}

func (this *MBlogController) Del() {
	blogService := service.BlogService{}
	userService := service.UserService{}
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["json"] = models.ReurnError(401, "")
		this.ServeJSON()
		return
	}
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	blog, err := blogService.GetBlog(id)
	if err != nil {
		this.Data["json"] = models.ReurnError(500, "")
		this.ServeJSON()
		return
	}
	if blog.UserId != uid.(int64) {
		this.Data["json"] = models.ReurnError(503, "")
		this.ServeJSON()
		return
	}
	blog.Delflag = 1
	err = blogService.DelBlog(blog)
	if err != nil {
		this.Data["json"] = models.ReurnError(500, "")
		this.ServeJSON()
		return
	}
	this.Data["json"] = models.ReurnSuccess("")
	this.ServeJSON()
	userService.CountBlog(uid.(int64))
	return
}

func (this *MBlogController) New() {
	uid := this.GetSession("userid")
	if uid == nil {
		this.Redirect("/login", 302)
		return
	}
	this.TplName = "newblog.html"
}