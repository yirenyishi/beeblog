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
	pid, _ := this.GetInt64("pid")
	fmt.Println("pid", pid)
	title := this.GetString("title")
	uid := this.GetSession("userid").(int64)

	note := &models.Note{Title: title, Pid: pid, UserId: uid}
	err := service.SaveNote(note)
	if err == nil {
		this.Data["json"] = note
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
		this.Data["json"] = note
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

func (this *NoteController) Note() {
	uid := this.GetSession("userid")
	fmt.Println("userid", uid)
	if uid == nil {
		this.Redirect("/login", 302)
	}
	noteColls, err := service.GetNoteColl(uid.(int64))
	if err != nil {
		if len(noteColls) > 0 {
			for i := 0; i < len(noteColls); i++ {
				notes, err1 := service.GetNoteByPid(noteColls[i].Id)
				if err1 != nil {
					noteColls[i].Notes = notes
				}
			}
		}
	} else {
		noteColls = make([]*models.NoteColl, 0)
	}
	fmt.Println(noteColls)
	this.Data["NoteColls"] = noteColls
	this.TplName = "note.html"
}
