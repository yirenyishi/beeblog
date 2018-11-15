package models

import "time"

type Attachment struct {
	Id         int64
	Name       string
	Alias       string
	Path       string    `orm:"index"`
	UploadTime time.Time `orm:"datetime"`
}
