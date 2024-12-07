package controllers

import "github.com/astaxie/beego"

type StaticsController struct {
	beego.Controller
}

func (s *StaticsController) Get() {
	s.TplName = "static.html"
}
