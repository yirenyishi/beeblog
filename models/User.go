package models

import "time"

type User struct {
	Id       int64
	UserName string    `orm:"unique"`
	NickName string
	UserPwd  string
	Salt     string
	Headimg  string
	Birthday time.Time `orm:"null;type(date)"`
	Email    string
	Mobile   string
	QQ       string
	HomeUrl  string
	Sex      int
	DescInfo string
	Ctime    time.Time `orm:"auto_now_add;type(datetime)"`
}
