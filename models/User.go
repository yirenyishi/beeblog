package models

import "time"

type User struct {
	Id       int64
	UserName string    `orm:"unique"`
	NickName string
	UserPwd  string
	Salt     string
	Headimg  string
	Birthday time.Time `orm:"auto_now_add;type(datetime)"`
	Email    string
	Mobile   string
	QQ       string
	HomeUrl  string
	Sex      int       `orm:"default(1)"`
	Status      int       `orm:"default(0)"`
	DescInfo string
	Ctime    time.Time `orm:"auto_now_add;type(datetime)"`

	BlogCount   int `orm:"default(0)"`
	BlogBrowes  int `orm:"default(0)"`
	BlogComment int `orm:"default(0)"`
	BlogLike    int `orm:"default(0)"`
}

func (u *User) TableName() string {
	return "auth_user"
}