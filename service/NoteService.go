package service

import (
	"beeblog/models"
	"github.com/astaxie/beego/orm"
)

type NoteService struct {
}

func EditNote(note *models.Note) error {
	o := orm.NewOrm()
	id, err := o.Update(note)
	if err == nil {
		note.Id = id
	}
	return err
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

func DelNote(note *models.Note) error {
	o := orm.NewOrm()
	_, err := o.Delete(note)
	return err
}

func GetNoteByPid(pid int64) ([]*models.Note, error) {
	var notes []*models.Note
	o := orm.NewOrm()
	qs := o.QueryTable(models.Note{})
	_, err := qs.Filter("Pid", pid).All(&notes)
	return notes, err
}

func CountNote(pid int64) (int64, error) {
	o := orm.NewOrm()
	totalCount, err := o.QueryTable(&models.Note{}).Filter("Pid", pid).Count()
	if err != nil {
		return 0,err
	}
	return totalCount, nil
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
