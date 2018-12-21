package controllers

import (
	"beeblog/models"
	"beeblog/service"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type BlogController struct {
	beego.Controller
}

func (this *BlogController) EditPage() {
	uid := this.GetSession("userid")
	if uid == nil {
		this.Redirect("/login", 302)
		return
	}
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	blog, err := service.GetBlog(id)
	if err != nil {
		this.Redirect("/500", 302)
		return
	}
	if blog.UserId != uid.(int64) {
		this.Redirect("/403", 302)
		return
	}
	this.Data["Blog"] = blog
	this.TplName = "editblog.html"
}

func (this *BlogController) Save() {
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
	err := service.SaveBlog(blog, labels)
	if err == nil {
		this.Data["json"] = models.ReurnData("",blog.Id)
	} else {
		this.Data["json"] = models.ReurnError(500, "保存失败")
	}
	this.ServeJSON()
	service.CountBlog(uid.(int64))
	return
}

func (this *BlogController) Edit() {
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
	blog,err :=service.GetBlog(id)
	if err != nil {
		this.Data["json"] = models.ReurnError(500, "保存失败")
		this.ServeJSON()
		return
	}
	blog.Title = title
	blog.BlogHtml = blogHtml
	blog.CategoryId = catoryId
	blog.Utime = time.Now()
	err = service.EditBlog(blog, labels)
	if err == nil {
		this.Data["json"] = models.ReurnSuccess("")
	} else {
		this.Data["json"] = models.ReurnError(500, "保存失败")
	}
	this.ServeJSON()
	service.CountBlog(uid.(int64))
	return
}

func (this *BlogController) Get() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	blog, err := service.GetBlog(id)
	if err != nil {
		this.Redirect("/500", 302)
		return
	}
	if uid := this.GetSession("userid"); uid != nil {
		if blog.UserId == uid.(int64) {
			this.Data["IsAuthor"] = true
		}
		if flag, err := service.IsLike(id, uid.(int64)); err == nil {
			this.Data["IsLike"] = flag
		}
	}
	if blogs, err := service.TopBlogByUser(blog.UserId); err == nil {
		this.Data["Top"] = blogs
	}
	this.Data["Blog"] = blog
	this.Data["UserId"] = this.GetSession("userid")
	this.Data["HeadImg"] = this.GetSession("headimg")
	this.Data["NickName"] = this.GetSession("nickname")
	this.Data["IsLogin"] = this.GetSession("nickname") != nil
	this.Data["IsDDesc"] = true
	this.TplName = "blog.html"
	if this.Ctx.Input.GetData("refresh") != true {
		service.CountBrows(blog.UserId)
		service.EditBlogBrows(id)
	}
	return
}

func (this *BlogController) Del() {
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["json"] = models.ReurnError(401, "")
		this.ServeJSON()
		return
	}
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	blog, err := service.GetBlog(id)
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
	err = service.DelBlog(blog)
	if err != nil {
		this.Data["json"] = models.ReurnError(500, "")
		this.ServeJSON()
		return
	}
	this.Data["json"] = models.ReurnSuccess("")
	this.ServeJSON()
	service.CountBlog(uid.(int64))
	return
}

func (this *BlogController) New() {
	uid := this.GetSession("userid")
	if uid == nil {
		this.Redirect("/login", 302)
		return
	}
	this.TplName = "newblog.html"
}

func (this *BlogController) BlogsPage() {
	cats, errcat := service.GetCats()
	if errcat != nil {
		this.Redirect("/500", 302)
		return
	}
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
	pages, err := service.FindBlogs(num, size, cat, flag)
	if err != nil {
		this.Redirect("/500", 302)
		return
	}
	this.Data["UserId"] = this.GetSession("userid")
	this.Data["HeadImg"] = this.GetSession("headimg")
	this.Data["NickName"] = this.GetSession("nickname")
	this.Data["IsLogin"] = this.GetSession("nickname") != nil
	this.Data["Page"] = pages
	this.Data["Cats"] = cats
	this.Data["Cat"] = cat
	this.Data["Flag"] = flag
	this.Data["IsBlog"] = true
	this.TplName = "blogs.html"
}

