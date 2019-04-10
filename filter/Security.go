package filter

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"beeblog/models"
	"encoding/json"
)

var LogFilter = func(ctx *context.Context) {
	logs.Info(ctx.Input.URI())
}
var FilterAdmin = func(ctx *context.Context) {
	if ctx.Input.Session("userid") != nil {
		ctx.Output.SetStatus(200)
		ctx.Output.Header("Access-Control-Allow-Origin","*")
		result := models.ReurnError(401, "")
		if b,err := json.Marshal(result); err != nil {
			ctx.Output.Body(b)
		}
		return
	}
}
