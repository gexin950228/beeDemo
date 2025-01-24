package controllers

import (
	"html/template"

	"github.com/astaxie/beego"
)

type TestXsrfController struct {
	beego.Controller
}

func (t *TestXsrfController) Get() {
	t.Data["xsrfData"] = template.HTML(t.XSRFFormHTML())
	username := t.GetSession("beegoSession")
	t.Data["username"] = username
	t.TplName = "xsrf.html"
}

func (t *TestXsrfController) Prepare() {
	//  controller级别关闭|开启xsrf
	t.EnableXSRF = false
}

func (t *TestXsrfController) Post() {
	t.TplName = "ok.html"
}
