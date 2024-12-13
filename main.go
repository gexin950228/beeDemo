package main

import (
	// "beeDemo/controllers"
	_ "beeDemo/routers"
	"beeDemo/utils"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	utils.LoadConfig()
	beego.SetViewsPath("front/")
	beego.SetStaticPath("/aaa", "static")
	beego.BConfig.WebConfig.EnableXSRF = false
	beego.BConfig.WebConfig.XSRFKey = "610ETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
	beego.BConfig.WebConfig.XSRFExpire = 3600
	beego.BConfig.WebConfig.Session.SessionOn = true
	// 全局过滤器
	//beego.InsertFilter("/*", beego.BeforeRouter, controllers.FilterUser)
	redisConn := utils.LoadRedisConfig()
	fmt.Println(redisConn)
	beego.Run()
}
