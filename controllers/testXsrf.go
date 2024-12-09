package controllers

import (
	"fmt"
	"html/template"

	"github.com/astaxie/beego"
)

type TestXsrfController struct {
	beego.Controller
}

func (t *TestXsrfController) Get() {
	t.Data["xsrfData"] = template.HTML(t.XSRFFormHTML())
	username := t.GetSession("beegoSession")
	fmt.Println(username)
	t.TplName = "xsrf.html"
}

func (t *TestXsrfController) Prepare() {
	//  controller级别关闭|开启xsrf
	t.EnableXSRF = false
}

func (t *TestXsrfController) Post() {
	fmt.Println("haaaaaa")
	t.TplName = "ok.html"
}
