package service

import (
	"beeblog/models"
	"github.com/astaxie/beego/orm"
)

type NoteService struct {
}

func SaveNote(note *models.Note) error {
	o := orm.NewOrm()
	id, err := o.Insert(note)
	if err == nil {
		note.Id = id
	}
	return err
}

func GetNote(note *models.Note) error {
	o := orm.NewOrm()
	return o.Read(note)
}

func GetNoteByPid(pid int64) ([]*models.Note, error) {
	var notes []*models.Note
	o := orm.NewOrm()
	qs := o.QueryTable(models.Note{})
	_, err := qs.Filter("Pid", pid).All(&notes)
	return notes, err
}

func SaveNoteColl(note *models.NoteColl) error {
	o := orm.NewOrm()
	id, err := o.Insert(note)
	if err == nil {
		note.Id = id
	}
	return err
}

func GetNoteColl(uid int64) ([]*models.NoteColl, error) {
	var notes []*models.NoteColl
	o := orm.NewOrm()
	qs := o.QueryTable(models.NoteColl{})
	_, err := qs.Filter("UserId", uid).All(&notes)
	return notes, err
}
