package controllers

import "github.com/astaxie/beego"

type Template2Controller struct {
	beego.Controller
}

func (t2 *Template2Controller) Get() {
	t2.TplName = "tpls/tpl2.html"
}
