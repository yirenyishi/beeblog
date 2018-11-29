package models

import "time"

type Comment struct {
	Id      int64
	CuserId int64     `orm:"default(0)"`
	BuserId int64     `orm:"default(0)"`
	BlogId  int64     `orm:"default(0)"`
	Ctime   time.Time `orm:"auto_now_add;type(datetime)"`
	Pid     int64     `orm:"default(0)"`
	ComVal  string    `orm:"type(text)"`

	Childs []*Comment `orm:"-"`
	CUser  *User      `orm:"-"`
	BUser  *User      `orm:"-"`
}
