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

type UserController struct {
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

type User struct {
	Id      int64
	Name    string
	Age     int64
	Address string
}

func (u *UserController) Get() {
	var user User
	user.Address = "北京市"
	user.Age = 29
	user.Name = "葛新"
	arrs := []int{1, 2, 3, 4, 5, 6}
	u.Data["user"] = user
	u.Data["arrs"] = arrs
	u.TplName = "user.html"
}
