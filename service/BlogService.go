package service

import (
	"github.com/astaxie/beego/orm"
	"beeblog/models"
	"beeblog/utils"
)

type BlogService struct {
}

func count(num int, size int, cat int64) (*utils.Page, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(&models.Blog{})
	qs.Filter("Delflag", 0)
	if cat != -1 {
		qs = qs.Filter("CategoryId", cat)
	}
	totalCount, err := qs.Count()
	if err != nil {
		return nil, err
	}
	return utils.PageUtil(totalCount, num, size), nil
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

func FindBlogs(num int, size int, cat int64, flag int) (*utils.Page, error) {
	page, err := count(num, size, cat)
	if err != nil {
		return nil, err
	}
	var blogs []*models.Blog
	o := orm.NewOrm()
	qs := o.QueryTable(&models.Blog{})
	qs = qs.Filter("Delflag", 0)
	if cat != -1 {
		qs = qs.Filter("CategoryId", cat)
	}
	if flag == 0 {
		qs = qs.OrderBy("-Ctime")
	} else {
		qs = qs.OrderBy("-Browses")
	}

	qs = qs.Limit(size,page.PageNo)
	_, err = qs.All(&blogs)
	if err != nil {
		return nil, err
	}
	page.List = blogs
	return page, nil
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
