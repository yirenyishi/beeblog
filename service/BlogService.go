package service

import (
	"github.com/astaxie/beego/orm"
	"beeblog/models"
)

type BlogService struct {
}

func GetBlog(id int64) (*models.Blog, error) {
	o := orm.NewOrm()
	blog := &models.Blog{Id: id}
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

func SaveBlog(blog *models.Blog, strs []string) error {
	o := orm.NewOrm()
	o.Begin()
	id, eror := o.Insert(blog)
	if eror != nil {
		o.Rollback()
		return eror
	} else {
		blog.Id = id
		if strs != nil && len(strs) > 0 {
			nlabels := make([]*models.NLabel, len(strs))
			for i := 0; i < len(strs); i++ {
				nlabels[i] = &models.NLabel{Title: strs[i], BlogId: id, UserId: blog.UserId}
			}
			if _, err := o.InsertMulti(len(nlabels), nlabels); err != nil {
				o.Rollback()
				return err
			}
		}
		o.Commit()
	}
	return nil
}
