package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "my-web-app/routers"
	"my-web-app/services"
)

func main() {
	beego.Run()
	services.InitServer()
}
