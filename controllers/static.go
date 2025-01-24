package controllers

import (
	"github.com/astaxie/beego"
)

type StaticsController struct {
	beego.Controller
}

func (s *StaticsController) Get() {
	name := s.GetString("name")
	s.Data["name"] = name
	s.TplName = "static.html"
}
