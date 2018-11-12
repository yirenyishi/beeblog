package models

type Category struct {
	Id         int64
	Title      string
	Views      int64 `orm:"index"`
	TopicCount int64
}
