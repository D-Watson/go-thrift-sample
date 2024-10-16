package main

import (
	_ "my-web-app/routers"
	"my-web-app/services"
)

func main() {
	//beego.Run()
	services.InitUserService()
}
