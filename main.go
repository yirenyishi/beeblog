package main

import (
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"beeblog/models"
	"github.com/astaxie/beego/orm"
	"beeblog/filter"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	models.RegistDB()
	beego.InsertFilter("/*", beego.BeforeRouter, filter.LogFilter)
	beego.InsertFilter("/api/*", beego.BeforeRouter, filter.FilterAdmin)
}
func main() {
	orm.Debug = false
	orm.RunSyncdb("default", false, true)
	beego.AddFuncMap("NAdd",NAdd)
	beego.SetLevel(beego.LevelInformational)
	beego.SetLogger("file", `{"filename":"/opt/logs/aiprose.log"}`)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.Run()
}

func NAdd(n1 int, n2 int) int{
	return n1 + n2
}
