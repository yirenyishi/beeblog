package service

import (
	"beeblog/models"
	"github.com/astaxie/beego/orm"
	"fmt"
)

func FindCommentByBlog(bid int64) ([]*models.Comment, error) {
	var comms []*models.Comment
	o := orm.NewOrm()
	_, err := o.QueryTable(&models.Comment{}).Filter("Pid", 0).OrderBy("-Ctime").All(&comms)
	if err != nil {
		return nil, err
	}
	if len(comms) > 0 {
		for i := 0; i < len(comms); i++ {
			var childs []*models.Comment
			_, childerrr := o.QueryTable(&models.Comment{}).Filter("Pid", comms[i].Id).OrderBy("-Ctime").All(&childs)
			if childerrr == nil {
				if len(childs) > 0{
					comms[i].Childs = childs
					for j:=0 ; j<len(childs);j++{
						cuser := &models.User{Id: childs[j].CuserId}
						o.Read(cuser)
						childs[j].CUser = cuser
						buser := &models.User{Id: childs[j].BuserId}
						o.Read(buser)
						childs[j].BUser = buser
					}
				}
			}
			cuser := &models.User{Id: comms[i].CuserId}
			o.Read(cuser)
			comms[i].CUser = cuser
		}
	}
	return comms, nil
}

func SaveComment(comment *models.Comment) error {
	o := orm.NewOrm()
	id, err := o.Insert(comment)
	if err == nil {
		comment.Id = id
		cuser := &models.User{Id: comment.CuserId}
		o.Read(cuser)
		comment.CUser = cuser
		if comment.BuserId !=0 {
			buser := &models.User{Id: comment.BuserId}
			o.Read(buser)
			comment.BUser = buser
		}
		return nil
	}
	return err
}

func ReadComment(comment *models.Comment) error {
	return orm.NewOrm().Read(comment)
}

func DelComment(id int64) error {
	comm := &models.Comment{Id: id}
	o := orm.NewOrm()
	err := o.Read(comm)
	if err != nil {
		return err
	}
	if comm.Pid != 0 {
		if _, err := o.QueryTable(models.Comment{}).Filter("Pid", id).Delete(); err != nil {
			fmt.Println(err)
		}

	}
	if _, err := o.Delete(comm); err != nil {
		return err
	}
	return nil
}
