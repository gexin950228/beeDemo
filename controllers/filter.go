package controllers

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"strings"
)

var FilterUser = func(ctx *context.Context) {
	loginUser := ctx.Input.GetData("loginUser")
	fmt.Printf("loginUser: %T  %v\n", loginUser, loginUser)
	// isLogin := true
	uri := ctx.Request.RequestURI
	fmt.Printf("uri: %T  %v\n", uri, uri)
	//fmt.Printf("uri: %v\n", uri)
	isLoginUri := strings.Contains(uri, "login") || strings.Contains(uri, "register") || strings.Contains(uri, "verifyCode")
	fmt.Printf("isLogin: %T  %v\n", true, isLoginUri)
	ctx.Redirect(302, "http://localhost:8080/article")
}
