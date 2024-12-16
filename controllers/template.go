package controllers

import "github.com/astaxie/beego"

type TemplateController struct {
	beego.Controller
}

func (t *TemplateController) Get() {
	var user User
	user.Age = 28
	user.Name = "葛新"
	user.Address = "北京市"
	ares := []int{1, 2, 3, 4, 5}
	t.Data["ares"] = ares
	t.Data["user"] = user
	t.Data["str"] = "hahaha"
	t.Data["age"] = 19
	t.Data["isVIP"] = false
	t.TplName = "tpls/tpl1.html"
}
