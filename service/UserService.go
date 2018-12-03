package service

import (
	"github.com/astaxie/beego/orm"
	"beeblog/models"
)

type UserService struct {
}

func GetUser(id int64) (*models.User, error) {
	o := orm.NewOrm()
	user := &models.User{Id: id}
	err := o.Read(user)
	return user, err
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
	browses := 0
	o.Raw("UPDATE `auth_user` SET `blog_count` = (SELECT count(id) num FROM blog where delflag = 0 and user_id =? ) WHERE `id` = ? ", uid, uid).QueryRow(&browses)
	return
}
func CountBrows(uid int64) {
	o := orm.NewOrm()
	browses := 0
	o.Raw("UPDATE `auth_user` SET `blog_browes` = (select  SUM(browses) browses from blog where user_id = ?) WHERE `id` = ? ", uid, uid).QueryRow(&browses) //获取总条数
	return
}
func CountComments(uid int64,bid int64) {
	o := orm.NewOrm()
	browses := 0
	o.Raw("UPDATE `auth_user` SET `blog_comment` = (select  count(id) num  from comment where cuser_id = ?) WHERE `id` = ? ", uid, uid).QueryRow(&browses) //获取总条数
	o.Raw("UPDATE `blog` SET `comments` = (select  count(id) num  from comment where blog_id = ?) WHERE `id` = ? ", bid, bid).QueryRow(&browses) //获取总条数
	return
}
func CountLike(uid int64,bid int64) {
	o := orm.NewOrm()
	browses := 0
	o.Raw("UPDATE `auth_user` SET `blog_like` = (select count(id) num from tb_like where user_id = ?) WHERE `id` = ?", uid, uid).QueryRow(&browses)
	o.Raw("UPDATE `blog` SET `likes` = (select count(id) num from tb_like where blog_id = ?) WHERE `id` = ?", bid, bid).QueryRow(&browses)
	return
}

func EditUser(user *models.User) (int64, error) {
	return orm.NewOrm().Update(user)
}

func EditHeadImg(user *models.User) (int64, error) {
	return orm.NewOrm().Update(user,"Headimg")
}
