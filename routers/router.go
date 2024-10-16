package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"my-web-app/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get,post:Get")
	beego.Router("/rr", &controllers.MainController{}, "get,post:GetRpc")
}
