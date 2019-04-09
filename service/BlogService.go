package service

import (
	"github.com/astaxie/beego/orm"
	"beeblog/models"
	"beeblog/utils"
)

type BlogService struct {
}

func (this *BlogService) count(num int, size int, cat int64) (*utils.Page, error) {
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

func (this *BlogService) countByUser(num int, size int, uid int64) (*utils.Page, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(&models.Blog{})
	totalCount, err := qs.Filter("Delflag", 0).Filter("UserId", uid).Count()
	if err != nil {
		return nil, err
	}
	return utils.PageUtil(totalCount, num, size), nil
}

func (this *BlogService) EditBlogBrows(id int64) {
	o := orm.NewOrm()
	blog := &models.Blog{Id: id}
	err := o.Read(blog)
	if err == nil {
		blog.Browses += 1
		o.Update(blog, "Browses")
	}
}

func (this *BlogService) TopBlogByUser(uid int64) ([]*models.Blog, error) {
	o := orm.NewOrm()
	var blogs []*models.Blog
	o.QueryTable(models.Blog{}).Filter("Delflag", 0).Filter("UserId", uid).Limit(12, 0).OrderBy("-Browses").All(&blogs)
	return blogs, nil
}

func (this *BlogService) ReadBlog(id int64) (*models.Blog, error) {
	o := orm.NewOrm()
	blog := &models.Blog{Id: id}
	if err := o.Read(blog); err != nil {
		return nil, err
	}
	return blog, nil
}

func (this *BlogService) GetBlog(id int64) (*models.Blog, error) {
	commentService := CommentService{}
	o := orm.NewOrm()
	blog := &models.Blog{Id: id}
	err := o.Read(blog)
	if err != nil {
		return nil, err
	}
	user := &models.User{Id: blog.UserId}
	err = o.Read(user)
	if err == nil {
		blog.User = user
	}
	var labels []*models.NLabel
	qs := o.QueryTable(&models.NLabel{})
	_, err = qs.Filter("BlogId", id).All(&labels)
	if err == nil {
		blog.Lables = labels
	}
	comms, berr := commentService.FindCommentByBlog(id)
	if berr == nil {
		blog.Comms = comms
	}
	return blog, nil
}

func (this *BlogService) DelBlog(blog *models.Blog) error {
	o := orm.NewOrm()
	_, err := o.Update(blog, "Delflag")
	return err
}

func (this *BlogService) FindBlogs(num int, size int, cat int64, flag int) (*utils.Page, error) {
	page, err := this.count(num, size, cat)
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

	qs = qs.Limit(size, (page.PageNo-1)*size)
	_, err = qs.All(&blogs)
	if err != nil {
		return nil, err
	}
	if len(blogs) > 0 {
		for i := 0; i < len(blogs); i++ {
			user := &models.User{Id: blogs[i].UserId}
			err = o.Read(user)
			if err == nil {
				blogs[i].UserName = user.UserName
				blogs[i].HeadImg = user.Headimg
			}
			category := &models.Category{Id: blogs[i].CategoryId}
			err = o.Read(category)
			if err == nil {
				blogs[i].CateName = category.Title
			}
		}
	}
	page.List = blogs
	return page, nil
}

func (this *BlogService) SaveBlog(blog *models.Blog, strs []string) error {
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

func (this *BlogService) EditBlog(blog *models.Blog, strs []string) error {
	o := orm.NewOrm()
	o.Begin()
	_, eror := o.Update(blog)
	if eror != nil {
		o.Rollback()
		return eror
	} else {
		o.QueryTable(models.NLabel{}).Filter("BlogId", blog.Id).Delete()
		if strs != nil && len(strs) > 0 {
			nlabels := make([]*models.NLabel, len(strs))
			for i := 0; i < len(strs); i++ {
				nlabels[i] = &models.NLabel{Title: strs[i], BlogId: blog.Id, UserId: blog.UserId}
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

func (this *BlogService) MeBlogs(num int, size int, flag int, uid int64) (*utils.Page, error) {
	page, err := this.countByUser(num, size, uid)
	if err != nil {
		return nil, err
	}
	var blogs []*models.Blog
	o := orm.NewOrm()
	qs := o.QueryTable(&models.Blog{})
	qs = qs.Filter("Delflag", 0)
	qs = qs.Filter("UserId", uid)
	if flag == 0 {
		qs = qs.OrderBy("-Ctime")
	} else {
		qs = qs.OrderBy("-Browses")
	}
	qs = qs.Limit(size, (page.PageNo-1)*size)
	_, err = qs.All(&blogs)
	if err != nil {
		return nil, err
	}
	page.List = blogs
	return page, nil
}

func (this *BlogService) IndexBlogs(size int, flag int) ([]*models.Blog, error) {
	var blogs []*models.Blog
	o := orm.NewOrm()
	qs := o.QueryTable(&models.Blog{})
	qs = qs.Filter("Delflag", 0)
	switch flag {
	case 0:
		qs = qs.OrderBy("-Ctime") //最新
	case 1:
		qs = qs.OrderBy("-Browses") //浏览量
	case 2:
		qs = qs.OrderBy("-Likes") // 收藏量
	case 3:
		qs = qs.OrderBy("-Comments") // 评论量
	}
	qs = qs.Limit(size, 0)
	_, err := qs.All(&blogs)
	return blogs, err
}

func (this *BlogService) GetNLabel(id int64) ([]*models.NLabel, error) {
	var labels []*models.NLabel
	o := orm.NewOrm()
	_, err := o.QueryTable(&models.NLabel{}).Filter("BlogId", id).All(&labels)
	if err != nil {
		return labels, err
	}
	return labels, nil
}
