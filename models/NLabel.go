package models

import "time"

type NLabel struct {
	Id     int64
	Title  string
	UserId int64
	BlogId int64
	Ctime  time.Time `orm:"auto_now_add;type(datetime)"`
}
