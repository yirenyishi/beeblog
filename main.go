package main

import (
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"beeblog/models"
	"github.com/astaxie/beego/orm"
	"beeblog/filter"
)

func init() {
	models.RegistDB()
	beego.InsertFilter("/*", beego.BeforeRouter, filter.FilterAdmin)
	//beego.InsertFilter("/*", beego.FinishRouter, filter.FilterLoginInfo)
}
func main() {
	orm.Debug = false
	orm.RunSyncdb("default", false, true)
	beego.AddFuncMap("NAdd",NAdd)
	beego.SetLevel(beego.LevelInformational)
	beego.SetLogger("file", `{"filename":"/opt/logs/aiprose.log"}`)
	beego.Run()
}

func NAdd(n1 int, n2 int) int{
	return n1 + n2
}
