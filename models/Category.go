package models

import "time"

type Category struct {
	Id    int64
	Title string
	Ctime time.Time `orm:"auto_now_add;type(datetime)"`
}
