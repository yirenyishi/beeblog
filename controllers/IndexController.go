package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"beeblog/service"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	blogs,_ := service.FindBlogs()
	fmt.Println(blogs)
	c.Data["Blogs"] = blogs
	c.Data["IsHome"] = true
	c.TplName = "index.html"
}
