package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"beeblog/models"
	"beeblog/service"
	"fmt"
)

type NoteController struct {
	beego.Controller
}

func (this *NoteController) Save() {
	pid,_ := this.GetInt64("pid")
	fmt.Println("pid",pid)
	title := this.GetString("title")
	uid := this.GetSession("userid").(int64)

	note := &models.Note{Title: title, Pid: pid, UserId: uid}
	err := service.SaveNote(note)
	if err == nil {
		this.Data["json"] =  note
	} else {
		this.Data["json"] = models.ReurnError("保存失败")
	}
	this.ServeJSON()
}

func (this *NoteController) SaveNoteColl() {
	title := this.GetString("title")
	uid := this.GetSession("userid").(int64)

	note := &models.NoteColl{Title: title, UserId: uid}
	err := service.SaveNoteColl(note)
	if err == nil {
		this.Data["json"] =  note
	} else {
		this.Data["json"] = models.ReurnError("保存失败")
	}
	this.ServeJSON()
}

func (this *NoteController) Get() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	note := &models.Note{Id: id}
	err := service.GetNote(note)
	if err == nil {
		this.Data["json"] = note
	}
	this.ServeJSON()
}
