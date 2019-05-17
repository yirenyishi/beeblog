package main

import (
	"beeblog/filter"
	"beeblog/models"
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
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

	logs.SetLogger(logs.AdapterFile, `{"filename":"/opt/logs/aiprose.log","level":3}`)
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
