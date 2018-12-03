package models

import (
	"time"
)

type Blog struct {
	Id         int64
	UserId     int64
	Title      string
	BlogValue  string    `orm:"type(text)"`
	BlogHtml   string    `orm:"type(text)"`
	Ctime      time.Time `orm:"auto_now_add;type(datetime)"`
	Utime      time.Time `orm:"auto_now_add;type(datetime)"`
	Browses    int64     `orm:"default(0)"`
	Likes      int64     `orm:"default(0)"`
	Comments   int64     `orm:"default(0)"`
	Top        int       `orm:"default(0)"`
	Hot        int       `orm:"default(0)"`
	Ttime      time.Time `orm:"null;type(date)"`
	Htime      time.Time `orm:"null;type(date)"`
	Delflag    int       `orm:"default(0)"`
	CategoryId int64

	User     *User      `orm:"-"`
	UserName string     `orm:"-"`
	HeadImg  string     `orm:"-"`
	CateName string     `orm:"-"`
	Lables   []*NLabel  `orm:"-"`
	Comms    []*Comment `orm:"-"`
}
