package models

import "time"

type Blog struct {
	Id int64
	UserId int64
	Title string
	BlogValue string  `orm:"size(5000)"`
	BlogHtml string  `orm:"size(5000)"`
	Ctime time.Time `orm:"datetime"`
	Utime time.Time `orm:"datetime"`
	Browses int64 `orm:"datetime"`
	Top int
	Hot int
	Ttime time.Time `orm:"datetime"`
	Htime time.Time `orm:"datetime"`
	Delflag int
	CategoryId int64
}
