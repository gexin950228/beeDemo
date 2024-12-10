package main

import (
	"beeDemo/controllers"
	_ "beeDemo/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetViewsPath("front/")
	beego.SetStaticPath("/aaa", "static")
	username := beego.AppConfig.String("username")
	fmt.Println(username)
	beego.BConfig.WebConfig.EnableXSRF = false
	beego.BConfig.WebConfig.XSRFKey = "610ETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
	beego.BConfig.WebConfig.XSRFExpire = 3600
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.InsertFilter("/*", beego.BeforeRouter, controllers.FilterUser)
	beego.Run()
}
