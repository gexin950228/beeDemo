package controllers

import "github.com/astaxie/beego"

type AbortController struct {
	beego.Controller
}

func (a *AbortController) Get() {
	a.Abort("700")
	a.TplName = "abort.html"
}

type ErrorController struct {
	beego.Controller
}

func (e *ErrorController) Error700() {
	e.Data["message"] = "数据库连接错误"
	e.TplName = "error700.html"
}
