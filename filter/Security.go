package filter

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

var FilterAdmin = func(ctx *context.Context) {
	url := ctx.Input.URI()
	refer := ctx.Input.Refer()
	logs.Info(url)
	logs.Info(refer)
	logs.Info("https://www.aiprose.com"+url)
	if "https://www.aiprose.com"+url == refer {
		ctx.Input.SetData("refresh", true)
	}
	//beego.Informational(url)
	//if  url != "/login"{
	//	ctx.Redirect(302, "/login")
	//}
}

var FilterLoginInfo = func(ctx *context.Context) {
	if ctx.Input.Session("userid") != nil {
		//ctx.
	}
}
