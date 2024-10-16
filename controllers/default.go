package controllers

import (
	"context"

	beego "github.com/beego/beego/v2/server/web"
	"my-web-app/client"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) GetRpc() {
	c.Data["json"] = client.HelloClient(context.Background())
	c.ServeJSON()
}
