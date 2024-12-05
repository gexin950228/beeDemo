package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type GreetController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "gexin17812126257@163.com"
	c.TplName = "index.tpl"
}

func (g *GreetController) Get() {
	g.Data["Website"] = "gexin.me"
	g.Data["Email"] = "861439031@qq.com"
	g.TplName = "greet.tpl"
}
