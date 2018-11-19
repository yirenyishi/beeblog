package models
/**
	文章
 */
type NoteColl struct {
	Id     int64
	UserId int64
	Title string

	Notes []*Note `orm:"-"`
}