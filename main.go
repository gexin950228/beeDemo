package main

import (
	"beeDemo/controllers"
	_ "beeDemo/routers"
	"beeDemo/utils"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
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

	// 创建通道
	//ch1 := make(chan string)
	//fmt.Println(ch1)
	//
	//ch14 := make(chan string, 10)
	//ch14 <- "haha"
	//st, ok := <-ch14
	//if ok {
	//	fmt.Printf("从管道读取到的内容： %T, %s\n", st, st)
	//}
	// 注册自定义错误信息
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
