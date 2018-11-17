package service

import (
	"github.com/astaxie/beego/orm"
	"beeblog/models"
)

type BlogService struct {
}

func GetBlog(id int64) (*models.Blog, error) {
	o := orm.NewOrm()
	blog := &models.Blog{Id:id}
	err := o.Read(blog)
	if err != nil {
		return nil, err
	}
	return blog, nil
}


func FindBlogs() ([]*models.Blog, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(&models.Blog{})
	var blogs []*models.Blog
	_, err := qs.Filter("Delflag", 0).All(&blogs)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func SaveBlog(blog *models.Blog) error {
	o := orm.NewOrm()
	id, eror := o.Insert(blog)
	if eror != nil {
		return eror
	} else {
		blog.Id = id
		o.Commit()
	}
	return nil
}
