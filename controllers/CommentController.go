package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
	"beeblog/service"
	"strconv"
)

type CommentController struct {
	beego.Controller
}

func (this *CommentController) Save() {
	commentService := service.CommentService{}
	blogService := service.BlogService{}
	userService := service.UserService{}
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["json"] = models.ReurnError(401, "")
		this.ServeJSON()
		return
	}
	blogId, berr := this.GetInt64("blog")
	if blogId == 0 || berr != nil {
		this.Data["json"] = models.ReurnError(403, "")
		this.ServeJSON()
		return
	}
	commVal := this.GetString("commval")
	blog, err := blogService.ReadBlog(blogId)
	if err != nil {
		this.Data["json"] = models.ReurnError(403, "")
		this.ServeJSON()
		return
	}
	comm := &models.Comment{BlogId: blogId, CuserId: uid.(int64), BuserId: blog.UserId, ComVal: commVal}
	if pid, _ := this.GetInt64("pid"); pid != 0 {
		parent := &models.Comment{Id: pid}
		if err := commentService.ReadComment(parent); err == nil {
			comm.BuserId = parent.CuserId
		}
		comm.Pid = pid
	}
	err = commentService.SaveComment(comm)
	if err == nil {
		this.Data["json"] = models.ReurnData("", comm)
	} else {
		this.Data["json"] = models.ReurnError(500, "保存失败")
	}
	this.ServeJSON()
	userService.CountComments(uid.(int64),blogId)
	return
}

func (this *CommentController) Del() {
	commentService := service.CommentService{}
	userService := service.UserService{}
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["json"] = models.ReurnError(401, "")
		this.ServeJSON()
		return
	}
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	comm := &models.Comment{Id: id}
	err := commentService.ReadComment(comm)
	if err != nil {
		this.Data["json"] = models.ReurnError(500, "")
		this.ServeJSON()
		return
	}
	if comm.CuserId != uid.(int64) {
		this.Data["json"] = models.ReurnError(403, "")
		this.ServeJSON()
		return
	}
	err = commentService.DelComment(id)
	if err == nil {
		this.Data["json"] = models.ReurnSuccess("")
	} else {
		this.Data["json"] = models.ReurnError(500, "保存失败")
	}
	this.ServeJSON()
	userService.CountComments(uid.(int64),id)
	return
}
