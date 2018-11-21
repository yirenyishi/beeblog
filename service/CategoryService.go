package service

import (
	"github.com/astaxie/beego/orm"
	"beeblog/models"
)

func GetCats() ([]*models.Category, error) {
	var notes []*models.Category
	o := orm.NewOrm()
	qs := o.QueryTable(models.Category{})
	_, err := qs.All(&notes)
	return notes, err
}
