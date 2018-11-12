package models

import "time"

type Topic struct {
	Id         int64
	Uid        int64
	Title      string
	Content    string    `orm:"size(5000)"`
	Attachment string
	Created    time.Time `orm:"index"`
	ViewCount  int64     `orm:"index"`
	Author     string
	ReplayCount int64
}