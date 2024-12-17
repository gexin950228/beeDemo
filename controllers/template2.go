package controllers

import (
	"github.com/astaxie/beego"
)

type Template2Controller struct {
	beego.Controller
}

func (t2 *Template2Controller) Get() {
	t2.Data["arrs"] = []int{1, 2, 3}
	t2.Data["str"] = "知了课堂"
	t2.TplName = "tpls/tpl2.html"
}
