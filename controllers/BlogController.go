package controllers

import (
	"beeblog/models"
	"beeblog/service"
	"beeblog/utils"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type BlogController struct {
	beego.Controller
}

func (this *BlogController) EditPage() {
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

func (this *BlogController) Save() {
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
	blogDesc := this.GetString("blogDesc")
	catory := this.GetString("catory")
	catoryId, _ := strconv.ParseInt(catory, 10, 64)
	labels := this.GetStrings("labels[]")
	blog := &models.Blog{Title: title, BlogHtml: blogHtml, CategoryId: catoryId, UserId: uid.(int64), BlogDesc: blogDesc}
	err := blogService.SaveBlog(blog, labels)
	if err == nil {
		this.Data["json"] = models.ReurnData("", blog.Id)
		blog.User.Salt = ""
		blog.User.UserPwd = ""
		utils.ESSave(blog)
	} else {
		this.Data["json"] = models.ReurnError(500, "保存失败")
	}
	this.ServeJSON()
	userService.CountBlog(uid.(int64))
	return
}

func (this *BlogController) Edit() {
	blogService := service.BlogService{}
	userService := service.UserService{}
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["json"] = models.ReurnError(401, "")
		this.ServeJSON()
		return
	}
	id, _ := this.GetInt64("id")
	title := this.GetString("title")
	blogHtml := this.GetString("blogHtml")
	blogDesc := this.GetString("blogDesc")
	catory := this.GetString("catory")
	catoryId, _ := strconv.ParseInt(catory, 10, 64)
	labels := this.GetStrings("labels[]")
	blog, err := blogService.GetBlog(id)
	if err != nil {
		this.Data["json"] = models.ReurnError(500, "保存失败")
		this.ServeJSON()
		return
	}
	blog.Title = title
	blog.BlogHtml = blogHtml
	blog.BlogDesc = blogDesc
	blog.CategoryId = catoryId
	blog.Utime = time.Now()
	err = blogService.EditBlog(blog, labels)
	if err == nil {
		this.Data["json"] = models.ReurnSuccess("")
		blog.User.Salt = ""
		blog.User.UserPwd = ""
		utils.ESSave(blog)
	} else {
		this.Data["json"] = models.ReurnError(500, "保存失败")
	}
	this.ServeJSON()
	userService.CountBlog(uid.(int64))
	return
}

func (this *BlogController) Get() {
	blogService := service.BlogService{}
	likeService := service.LikeService{}
	userService := service.UserService{}
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	blog, err := blogService.GetBlog(id)
	blog.User.Salt = ""
	blog.User.UserPwd = ""
	if err != nil {
		this.Redirect("/404", 302)
		return
	}
	if uid := this.GetSession("userid"); uid != nil {
		if blog.UserId == uid.(int64) {
			this.Data["IsAuthor"] = true
		}
		if flag, err := likeService.IsLike(id, uid.(int64)); err == nil {
			this.Data["IsLike"] = flag
		}
	}
	if blogs, err := blogService.TopBlogByUser(blog.UserId); err == nil {
		this.Data["Top"] = blogs
	}
	//utils.Index()
	//utils.ESSave(blog)
	this.Data["Blog"] = blog
	this.Data["UserId"] = this.GetSession("userid")
	this.Data["HeadImg"] = this.GetSession("headimg")
	this.Data["NickName"] = this.GetSession("nickname")
	this.Data["IsLogin"] = this.GetSession("nickname") != nil
	this.Data["IsDDesc"] = true
	this.TplName = "blog.html"
	if this.Ctx.Input.GetData("refresh") != true {
		userService.CountBrows(blog.UserId)
		blogService.EditBlogBrows(id)
	}
	return
}

func (this *BlogController) Del() {
	blogService := service.BlogService{}
	userService := service.UserService{}
	likeService := service.LikeService{}
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["json"] = models.ReurnError(401, "")
		this.ServeJSON()
		return
	}
	idStr := this.Ctx.Input.Param(":id")
	utils.ESDelete(idStr)
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
	likeService.DelLikeByBlog(id)
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
	catService := service.CategoryService{}
	blogService := service.BlogService{}
	cats, errcat := catService.GetCats()
	if errcat != nil {
		this.Redirect("/404", 302)
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
	pages, err := blogService.FindBlogs(num, size, cat, flag)
	if err != nil {
		this.Redirect("/404", 302)
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

func (this *BlogController) Search() {
	blog, _ := utils.Search("elk")
	this.Data["json"] = models.ReurnData("", blog)
	this.ServeJSON()
}
