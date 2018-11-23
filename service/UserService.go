package service

import (
	"github.com/astaxie/beego/orm"
	"beeblog/models"
)

type UserService struct {
}

func GetUser(id int64) (*models.User,error) {
	o := orm.NewOrm()
	user := &models.User{Id:id}
	err := o.Read(user)
	return user,err
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

func CountBlog(uid int64) {
	o := orm.NewOrm()
	totalItem := 0
	err := o.Raw("SELECT count(*) FROM blog where delflag = 0 and user_id =? ", uid).QueryRow(&totalItem) //获取总条数
	if err != nil {
		return
	}
	user := &models.User{Id: uid}
	err = o.Read(user)
	if err != nil {
		return
	}
	user.BlogCount = totalItem
	o.Update(user, "BlogCount")
	return
}
func CountBrows(uid int64){
	o := orm.NewOrm()
	browses := 0
	o.Raw("UPDATE `user` SET `blog_browes` = (select  SUM(browses) browses from blog where user_id = ?1) WHERE `id` = ?2 ", uid,uid).QueryRow(&browses) //获取总条数
	return
}
func CountComments(uid int64) {

}
func CountLike(uid int64) {

}
