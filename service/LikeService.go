package service

import (
	"beeblog/models"
	"github.com/astaxie/beego/orm"
	"fmt"
)

func SaveLike(like *models.Like) (int64, error) {
	return orm.NewOrm().Insert(like)
}
func DelLike(like *models.Like) (int64, error) {
	return orm.NewOrm().QueryTable(models.Like{}).Filter("BlogId",like.BlogId).Filter("UserId", like.UserId).Delete()
}

func IsLike(bid int64, uid int64) (bool, error) {
	totalCount, err := orm.NewOrm().QueryTable(&models.Like{}).Filter("BlogId", bid).Filter("UserId", uid).Count()
	if err == nil {
		fmt.Println(totalCount,"like count")
		if totalCount > 0 {
			return true, nil
		} else {
			return false, nil
		}
	}else{
		fmt.Println(err)
		return false,err
	}
}
