package main

import (
	_ "beeDemo/routers"
	"fmt"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetViewsPath("front/")
	beego.SetStaticPath("/aaa", "static")
	username := beego.AppConfig.String("username")
	fmt.Println(username)
	beego.Run()
}
