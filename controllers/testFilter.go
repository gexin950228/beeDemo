package controllers

import (
	"fmt"
	"github.com/astaxie/beego/context"
)

// FilterUser 访问的url如果不是login，就跳转到login
var FilterUser = func(ctx *context.Context) {
	redirectUri := ctx.Input.Query(":redirectUri")
	fmt.Printf("redirectUri: %s, %T", redirectUri, redirectUri)
	ok := ctx.Input.Session("beegoDemo")
	if ok == "" {
		ctx.Redirect(302, "/login?redirectUri="+redirectUri)
	} else {
		fmt.Println("已登陆")
	}

}
