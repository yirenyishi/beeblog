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
	beego.InsertFilter("/*", beego.BeforeRouter, filter.FilterAdmin1)
}
func main() {
	orm.Debug = true
	orm.RunSyncdb("default",false,true)


	beego.Run()
}

