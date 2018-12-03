package models

import "time"

/**
 笔记
 */
type Note struct {
	Id       int64
	UserId   int64
	Title    string
	NoteHtml string    `orm:"type(text)"`
	Pid      int64
	Utime    time.Time `orm:"auto_now_add;type(datetime)"`
}
