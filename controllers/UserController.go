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
	u.TplName = "login.html"
}
func (u *UserController) RegistPage() {
	u.TplName = "regist.html"
}

func (this *UserController) PersonBlog() {
	uid := this.GetSession("userid")
	if uid == nil {
		this.Redirect("/login", 302)
		return
	}
	size := 15
	num, _ := this.GetInt("num")
	if num <= 0 {
		num = 1
	}
	flag, _ := this.GetInt("flag")
	page, err := service.MeBlogs(num, size, flag, uid.(int64))
	if err != nil {
		if uid == nil {
			this.Redirect("/500", 302)
			return
		}
	}
	user, uerr := service.GetUser(uid.(int64))
	if uerr != nil {
		if uid == nil {
			this.Redirect("/500", 302)
			return
		}
	}
	this.Data["Page"] = page
	this.Data["IsMeBlog"] = true
	this.Data["Flag"] = 0
	this.Data["User"] = user
	this.TplName = "ublogs.html"
}

func (this *UserController) PersonNote() {
	uid := this.GetSession("userid")
	if uid == nil {
		this.Redirect("/login", 302)
		return
	}
	notColl, err := service.GetNoteColl(uid.(int64))
	if err == nil {
		if len(notColl) > 0 {
			for i := 0; i < len(notColl); i++ {
				count, _ := service.CountNote(notColl[i].Id)
				notColl[i].Count = count
			}
		}
	} else {
		notColl = make([]*models.NoteColl, 0)
	}
	if err != nil {
		if uid == nil {
			this.Redirect("/500", 302)
			return
		}
	}
	user, uerr := service.GetUser(uid.(int64))
	if uerr != nil {
		if uid == nil {
			this.Redirect("/500", 302)
			return
		}
	}
	this.Data["Note"] = notColl
	this.Data["IsMeNote"] = true
	this.Data["User"] = user
	this.TplName = "unote.html"
}

func (this *UserController) PersonLike() {
	uid := this.GetSession("userid")
	if uid == nil {
		this.Redirect("/login", 302)
		return
	}
	size := 15
	num, _ := this.GetInt("num")
	if num <= 0 {
		num = 1
	}
	flag, _ := this.GetInt("flag")
	page, err := service.MeBlogs(num, size, flag, uid.(int64))
	if err != nil {
		if uid == nil {
			this.Redirect("/500", 302)
			return
		}
	}
	user, uerr := service.GetUser(uid.(int64))
	if uerr != nil {
		if uid == nil {
			this.Redirect("/500", 302)
			return
		}
	}
	this.Data["Page"] = page
	this.Data["IsMeBlog"] = true
	this.Data["Flag"] = 0
	this.Data["User"] = user
	this.TplName = "ublogs.html"
}

func (this *UserController) Login() {
	username := this.GetString("username")
	userpwd := this.GetString("userpwd")
	if username == "" {
		this.Data["json"] = models.ReurnError(1, "用户名为空")
		this.ServeJSON()
		return
	}
	if len(username) < 4 {
		this.Data["json"] = models.ReurnError(1, "用户名最低4位")
		this.ServeJSON()
		return
	}
	if userpwd == "" {
		this.Data["json"] = models.ReurnError(1, "密码为空")
		this.ServeJSON()
		return
	}
	if len(userpwd) < 6 {
		this.Data["json"] = models.ReurnError(1, "密码最低6位")
		this.ServeJSON()
		return
	}
	user, error := service.FindByUserName(username)
	if error == nil && user != nil {
		h := md5.New()
		h.Write([]byte(userpwd + user.Salt))
		userpwd = hex.EncodeToString(h.Sum(nil))
		if userpwd == user.UserPwd {
			this.Data["json"] = models.ReurnSuccess("")
			this.SetSession("userid", user.Id)
			this.SetSession("nickname", user.NickName)
		} else {
			this.Data["json"] = models.ReurnError(1, "用户名或密码错误")
		}
	} else {
		this.Data["json"] = models.ReurnError(1, "用户名不存在")
	}
	this.ServeJSON()
	return
}

func (this *UserController) Regist() {
	username := this.GetString("username")
	userpwd := this.GetString("userpwd")
	username = strings.Replace(username, " ", "", -1)
	userpwd = strings.Replace(userpwd, " ", "", -1)
	if username == "" {
		this.Data["json"] = models.ReurnError(1, "用户名为空")
		this.ServeJSON()
		return
	}
	if len(username) < 4 {
		this.Data["json"] = models.ReurnError(1, "用户名最低4位")
		this.ServeJSON()
		return
	}
	if userpwd == "" {
		this.Data["json"] = models.ReurnError(1, "密码为空")
		this.ServeJSON()
		return
	}
	if len(userpwd) < 6 {
		this.Data["json"] = models.ReurnError(1, "密码最低6位")
		this.ServeJSON()
		return
	}
	user, _ := service.FindByUserName(username)
	if user != nil {
		this.Data["json"] = models.ReurnError(1, "用户已经存在")
		this.ServeJSON()
		return
	}
	h := md5.New()
	h.Write([]byte(strconv.FormatInt(time.Now().Unix(), 10) + beego.AppConfig.String("host")))
	salt := hex.EncodeToString(h.Sum(nil))
	h = md5.New()
	h.Write([]byte(userpwd + salt))
	userpwd = hex.EncodeToString(h.Sum(nil))
	user = &models.User{UserName: username, NickName: username, UserPwd: userpwd, Salt: salt}
	err := service.SaveUser(user)
	if err == nil {
		this.Data["json"] = models.ReurnSuccess("")
	} else {
		this.Data["json"] = models.ReurnError(1, "注册失败")
	}
	this.ServeJSON()
	return
}

func (this *UserController) Logout() {
	this.DelSession("userid")
	this.DelSession("nickname")
	this.Redirect("/", 302)
	return
}
