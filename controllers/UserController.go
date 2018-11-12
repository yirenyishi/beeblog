package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (u *UserController) LoginPage() {
	u.Ctx.WriteString("login page")
}

func (u *UserController) Login() {
	u.Ctx.WriteString("login method")
}