package controllers

import (
	"github.com/astaxie/beego/context"
	"strings"
)

var FilterUser = func(ctx *context.Context) {
	loginUser := ctx.Input.GetData("loginUser")
	// isLogin := true
	uri := ctx.Request.RequestURI
	//fmt.Printf("uri: %v\n", uri)
	isLoginUri := strings.Contains(uri, "login") || strings.Contains(uri, "register") || strings.Contains(uri, "verifyCode")
	if loginUser == nil || isLoginUri {
		ctx.Redirect(302, "/login?redirectUri="+uri)
	}
}
