package main

import (
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"beeblog/models"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegistDB()
}
func main() {
	orm.Debug = true
	orm.RunSyncdb("default",false,true)
	beego.Run()
}

