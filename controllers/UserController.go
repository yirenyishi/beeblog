package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/service"
	"beeblog/models"
	"strings"
	"crypto/md5"
	"time"
	"strconv"
	"encoding/hex"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) LoginPage() {
	//o := orm.NewOrm()
	//
	//category := &models.Category{Title: "slene", Views: 5, TopicCount: 0}
	//
	//id,err := o.Insert(category)
	//fmt.Printf("ID: %d, ERR: %v\n", id, err)
	//
	//category.Title = "nelson"
	//num, err := o.Update(&category)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	//
	//c := &models.Category{Id:1}
	//err := o.Read(c)
	//fmt.Printf("ERR: %v\n", err)
	//fmt.Println(c.Title)
	////c.Title = "nelson"
	////num, err := o.Update(c)
	////fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	//qs := o.QueryTable(&models.Category{})
	//var tests []*models.Category
	//qs.Filter("Title","slene")
	//qs.All(&tests)
	//
	//for i:=0; i<len(tests) ; i++ {
	//	fmt.Println(tests[i].Title,tests[i].Id)
	//}
	//fmt.Println(len(tests))
	//num, err = o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	u.Ctx.WriteString("login page")
}

func (this *UserController) Login() {
	username := this.GetString("username")
	userpwd := this.GetString("userpwd")
	if username == "" {
		this.Data["json"] = models.ReurnError("用户名为空")
		this.ServeJSON()
	}
	if len(username) < 4 {
		this.Data["json"] = models.ReurnError("用户名最低4位")
		this.ServeJSON()
	}
	if userpwd == "" {
		this.Data["json"] = models.ReurnError("密码为空")
		this.ServeJSON()
	}
	if len(userpwd) < 6 {
		this.Data["json"] = models.ReurnError("密码最低6位")
		this.ServeJSON()
	}
	user, error := service.FindByUserName(username)
	if error == nil && user != nil {
		h := md5.New()
		h.Write([]byte(userpwd + user.Salt))
		userpwd = hex.EncodeToString(h.Sum(nil))
		if userpwd == user.UserPwd {
			this.Data["json"] = models.ReurnSuccess("")
		} else {
			this.Data["json"] = models.ReurnError("用户名或密码错误")
		}
	} else {
		this.Data["json"] = models.ReurnError("用户名不存在")
	}
	this.ServeJSON()
}

func (this *UserController) Regist() {
	username := this.GetString("username")
	userpwd := this.GetString("userpwd")
	username = strings.Replace(username, " ", "", -1)
	userpwd = strings.Replace(userpwd, " ", "", -1)
	if username == "" {
		this.Data["json"] = models.ReurnError("用户名为空")
		this.ServeJSON()
	}
	if len(username) < 4 {
		this.Data["json"] = models.ReurnError("用户名最低4位")
		this.ServeJSON()
	}
	if userpwd == "" {
		this.Data["json"] = models.ReurnError("密码为空")
		this.ServeJSON()
	}
	if len(userpwd) < 6 {
		this.Data["json"] = models.ReurnError("密码最低6位")
		this.ServeJSON()
	}
	user, _ := service.FindByUserName(username)
	if user != nil {
		this.Data["json"] = models.ReurnError("用户已经存在")
		this.ServeJSON()
	}
	h := md5.New()
	h.Write([]byte(strconv.FormatInt(time.Now().Unix(), 10) + beego.AppConfig.String("host")))
	salt := hex.EncodeToString(h.Sum(nil))
	h = md5.New()
	h.Write([]byte(userpwd + salt))
	userpwd = hex.EncodeToString(h.Sum(nil))
	user = &models.User{UserName: username, UserPwd: userpwd, Salt: salt}
	err := service.SaveUser(user)
	if err == nil {
		this.Data["json"] = user
	} else {
		this.Data["json"] = models.ReurnError("注册失败")
	}
	this.ServeJSON()
}
