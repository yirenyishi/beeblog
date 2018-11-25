package service

import (
	"beeblog/models"
	"github.com/astaxie/beego/orm"
	"fmt"
	"beeblog/utils"
)

func SaveLike(like *models.Like) (int64, error) {
	return orm.NewOrm().Insert(like)
}
func DelLike(like *models.Like) (int64, error) {
	return orm.NewOrm().QueryTable(models.Like{}).Filter("BlogId", like.BlogId).Filter("UserId", like.UserId).Delete()
}

func IsLike(bid int64, uid int64) (bool, error) {
	totalCount, err := orm.NewOrm().QueryTable(&models.Like{}).Filter("BlogId", bid).Filter("UserId", uid).Count()
	if err == nil {
		fmt.Println(totalCount, "like count")
		if totalCount > 0 {
			return true, nil
		} else {
			return false, nil
		}
	} else {
		fmt.Println(err)
		return false, err
	}
}

func MeLikes(num int, size int, uid int64) (*utils.Page, error) {
	page, err := countLike(num, size, uid)
	if err != nil {
		return nil, err
	}
	var likes []*models.Like
	o := orm.NewOrm()
	qs := o.QueryTable(&models.Like{})
	qs = qs.Filter("UserId", uid)
	qs = qs.Limit(size, (page.PageNo-1)*size)
	if _, err = qs.All(&likes); err != nil {
		return nil, err
	}
	if len(likes) > 0 {
		for i := 0; i < len(likes); i++ {
			blog := &models.Blog{Id: likes[i].BlogId}
			if err := o.Read(blog); err == nil {
				likes[i].Blog = blog
			}
		}
	}
	page.List = likes
	return page, nil
}

func countLike(num int, size int, uid int64) (*utils.Page, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(&models.Like{})
	totalCount, err := qs.Filter("UserId", uid).Count()
	if err != nil {
		return nil, err
	}
	return utils.PageUtil(totalCount, num, size), nil
}
