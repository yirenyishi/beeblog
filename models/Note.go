package models

type Note struct {
	Id     int64
	UserId int64
	Title string
	NoteVal string `orm:"size(3500)"`
	NoteHtml string `orm:"size(5000)"`
	Pid int64
}