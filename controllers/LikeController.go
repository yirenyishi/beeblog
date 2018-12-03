package controllers

import (
	"strconv"
	"beeblog/service"
	"github.com/astaxie/beego"
	"beeblog/models"
)

type LikeController struct {
	beego.Controller
}

func (this *LikeController) Save() {
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["json"] = models.ReurnError(401, "")
		this.ServeJSON()
		return
	}
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	like := &models.Like{BlogId: id, UserId: uid.(int64)}
	if _, err := service.SaveLike(like); err != nil {
		this.Data["json"] = models.ReurnError(500, "保存失败")
	}else{
		this.Data["json"] = models.ReurnSuccess("")
	}
	this.ServeJSON()
	service.CountLike(uid.(int64),id)
	return
}

func (this *LikeController) Delete() {
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["json"] = models.ReurnError(401, "")
		this.ServeJSON()
		return
	}
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	like := &models.Like{BlogId: id, UserId: uid.(int64)}
	if _, err := service.DelLike(like); err != nil {
		this.Data["json"] = models.ReurnError(500, "保存失败")
	}else{
		this.Data["json"] = models.ReurnSuccess("")
	}
	this.ServeJSON()
	service.CountLike(uid.(int64),id)
	return
}
