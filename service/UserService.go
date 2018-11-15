package service

import (
	"github.com/astaxie/beego/orm"
	"beeblog/models"
)

type UserService struct {
}

func FindByUserName(username string) (*models.User, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(&models.User{})
	var users []*models.User
	_, err := qs.Filter("UserName", username).All(&users)
	if err != nil {
		return nil, err
	}
	if len(users) != 0 {
		return users[0], nil
	}
	return nil, nil
}

func SaveUser(user *models.User) error {
	o := orm.NewOrm()
	id, eror := o.Insert(user)
	if eror != nil {
		return eror
	} else {
		user.Id = id
		o.Commit()
	}
	return nil
}
