package main

import (
	"beeDemo/controllers"
<<<<<<< HEAD
	// "beeDemo/controllers"
=======
>>>>>>> 8b8e7ec9f9f25040bdeebafc323e59cc42cb1e31
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
<<<<<<< HEAD
	// beego.InsertFilter("/*", beego.BeforeRouter, controllers.FilterUser)
=======
	beego.InsertFilter("/*", beego.BeforeRouter, controllers.FilterUser)
>>>>>>> 8b8e7ec9f9f25040bdeebafc323e59cc42cb1e31
	redisConn := utils.LoadRedisConfig()
	fmt.Println(redisConn)

	// 注册自定义错误信息
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
