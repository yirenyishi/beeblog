package models

import "time"

type Like struct {
	Id     int64
	UserId int64
	BlogId int64
	Ltime  time.Time `orm:"auto_now_add;type(datetime)"`

	Note Note `orm:"-"`
}
