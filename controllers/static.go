package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type StaticsController struct {
	beego.Controller
}

func (s *StaticsController) Get() {
	name := s.GetString("name")
	fmt.Printf("name: %v\n", name)
	s.TplName = "static.html"
}
