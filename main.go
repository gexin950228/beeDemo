package main

import (
	_ "beeDemo/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetViewsPath("front/")
	beego.SetStaticPath("/aaa", "static")
	beego.Run()
}
