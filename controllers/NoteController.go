package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"beeblog/models"
	"beeblog/service"
)

type NoteController struct {
	beego.Controller
}

func (this *NoteController) Save() {
	pid, _ := this.GetInt64("pid")
	title := this.GetString("title")
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["json"] = models.ReurnError(401, "保存失败")
		this.ServeJSON()
		return
	}
	note := &models.Note{Title: title, Pid: pid, UserId: uid.(int64)}
	err := service.SaveNote(note)
	if err == nil {
		this.Data["json"] = note
	} else {
		this.Data["json"] = models.ReurnError(500, "保存失败")
	}
	this.ServeJSON()
	return
}
func (this *NoteController) Edit() {
	idStr := this.Ctx.Input.Param(":id")
	noteHtml := this.GetString("noteHtml")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["json"] = models.ReurnError(401, "")
		this.ServeJSON()
		return
	}
	note := &models.Note{Id: id}
	err1 := service.GetNote(note)
	if err1 != nil {
		this.Data["json"] = models.ReurnError(500, "保存失败")
		this.ServeJSON()
		return
	}
	if uid != note.UserId {
		this.Data["json"] = models.ReurnError(403, "")
		this.ServeJSON()
		return
	}
	note.NoteHtml = noteHtml
	err := service.EditNote(note)
	if err == nil {
		this.Data["json"] = models.ReurnSuccess("")
	} else {
		this.Data["json"] = models.ReurnError(500, "保存失败")
	}
	this.ServeJSON()
	return
}

func (this *NoteController) SaveNoteColl() {
	title := this.GetString("title")
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["json"] = models.ReurnError(401, "")
		this.ServeJSON()
		return
	}
	note := &models.NoteColl{Title: title, UserId: uid.(int64)}
	err := service.SaveNoteColl(note)
	if err == nil {
		this.Data["json"] = models.ReurnSuccess("")
	} else {
		this.Data["json"] = models.ReurnError(500, "保存失败")
	}
	this.ServeJSON()
	return
}

func (this *NoteController) EditNoteColl() {
	title := this.GetString("title")
	id, _ := this.GetInt64("id")
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["json"] = models.ReurnError(401, "")
		this.ServeJSON()
		return
	}
	err := service.EditNoteColl(title, id, uid.(int64))
	if err == nil {
		this.Data["json"] = models.ReurnSuccess("")
	} else {
		this.Data["json"] = models.ReurnError(500, "保存失败")
	}
	this.ServeJSON()
	return
}

func (this *NoteController) Get() {
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["json"] = models.ReurnError(401, "")
		this.ServeJSON()
		return
	}
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	note := &models.Note{Id: id}
	err := service.GetNote(note)
	if err == nil {
		this.Data["json"] = note
	}
	if note.UserId != uid.(int64) {
		this.Data["json"] = models.ReurnError(403, "")
		this.ServeJSON()
		return
	}
	this.ServeJSON()
	return
}
func (this *NoteController) DelNoteColl() {
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["json"] = models.ReurnError(401, "")
		this.ServeJSON()
		return
	}
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	err := service.DelNoteColl(id, uid.(int64))
	if err != nil {
		this.Data["json"] = models.ReurnError(500, "")
	} else {
		this.Data["json"] = models.ReurnSuccess("")
	}
	this.ServeJSON()
	return
}

func (this *NoteController) Delete() {
	uid := this.GetSession("userid")
	if uid == nil {
		this.Data["json"] = models.ReurnError(401, "")
		this.ServeJSON()
		return
	}
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	note := &models.Note{Id: id}
	err := service.GetNote(note)
	if err != nil {
		this.Data["json"] = models.ReurnError(500, "")
		this.ServeJSON()
		return
	}
	if note.UserId != uid.(int64) {
		this.Data["json"] = models.ReurnError(403, "")
		this.ServeJSON()
		return
	}
	err = service.DelNote(note)
	if err != nil {
		this.Data["json"] = models.ReurnError(500, "")
		this.ServeJSON()
		return
	}
	this.Data["json"] = models.ReurnSuccess("")
	this.ServeJSON()
	return
}

func (this *NoteController) Note() {
	uid := this.GetSession("userid")
	if uid == nil {
		this.Redirect("/login", 302)
	}
	noteColls, err := service.GetNoteColl(uid.(int64))
	if err == nil {
		if len(noteColls) > 0 {
			for i := 0; i < len(noteColls); i++ {
				notes, err1 := service.GetNoteByPid(noteColls[i].Id)
				if err1 == nil {
					noteColls[i].Notes = notes
				}
			}
		}
	} else {
		noteColls = make([]*models.NoteColl, 0)
	}
	this.Data["NoteColls"] = noteColls
	this.TplName = "note.html"
}
