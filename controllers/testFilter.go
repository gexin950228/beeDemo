package controllers

import (
	context1 "context"
	"fmt"
	context2 "github.com/astaxie/beego/context"
	"net/http"
	"strings"
)

func withRequestUri(r *http.Request) context1.Context {
	return context1.WithValue(r.Context(), "RequestURI", r.RequestURI)
}

// FilterUser 访问的url如果不是login，就跳转到login
var FilterUser = func(ctx *context2.Context) {
	uri := ctx.Request.RequestURI
	fmt.Printf("uri: %s\n", uri)
	if !strings.Contains(uri, "/login") {
		ok := ctx.Input.Session("beegoDemo")
		if ok == nil {
			ctx.Redirect(302, "/login?redirectUri="+uri)
		}
	}
}
