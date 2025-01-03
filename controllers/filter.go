package controllers

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"net/http"
	"strings"
)

var FilterUser = func(ctx *context.Context) {
	loginUser := ctx.Input.Session("loginUser")
	//fmt.Printf("loginUser: %T  %v\n", loginUser, loginUser)
	isLogin := loginUser != nil
	uri := ctx.Request.RequestURI
	//fmt.Printf("uri: %v\n", uri)
	isLoginUri := strings.Contains(uri, "login")
	//fmt.Printf("isLogin: %v isLoginUri: %v\n", isLogin, isLoginUri)
	cookieUser := ctx.GetCookie("loginUser")
	fmt.Printf("cookieUser: %v\n", cookieUser)
	if !isLogin {
		if !isLoginUri {
			if !isLogin {
				ctx.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("/login/?redirectUri=%s", uri))
			}
		}
	}
}
