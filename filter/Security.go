package filter

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/context"
)

var FilterAdmin = func(ctx *context.Context) {
	url := ctx.Input.URL()
	logs.Info("##### filter url : %s", url)
	//if  url != "/login"{
	//	ctx.Redirect(302, "/login")
	//}
}

var FilterLoginInfo = func(ctx *context.Context) {
	if ctx.Input.Session("userid") != nil {
		//ctx.
	}
}
