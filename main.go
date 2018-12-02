package main

import (
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"beeblog/models"
	"beeblog/filter"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegistDB()
	beego.InsertFilter("/*", beego.BeforeRouter, filter.FilterAdmin)
	beego.InsertFilter("/*", beego.FinishRouter, filter.FilterLoginInfo)
}
func main() {
	orm.Debug = false
	orm.RunSyncdb("default", false, true)
	beego.AddFuncMap("NAdd",NAdd)
	beego.Run()
}

func NAdd(n1 int, n2 int) int{
	return n1 + n2
}
