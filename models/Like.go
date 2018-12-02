package models

import "time"

type Like struct {
	Id     int64
	UserId int64
	BlogId int64
	Ltime  time.Time `orm:"auto_now_add;type(datetime)"`

	Blog *Blog `orm:"-"`
}

func (u *Like) TableName() string {
	return "tb_like"
}