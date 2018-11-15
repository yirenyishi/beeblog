package models

import "time"

type Comment struct {
	Id      int64
	CuserId int64
	BuserId int64
	BlogId  int64
	Ctime   time.Time `orm:"datetime"`
	Pid     int64
	ComVal  string    `orm:"size(2000)"`
}
