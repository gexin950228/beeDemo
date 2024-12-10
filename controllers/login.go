package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	sourceUri := l.GetString("redirectUri")
	fmt.Println("====================")
	fmt.Printf("sourceUri: %s\n", sourceUri)
	l.Data["sourceUri"] = sourceUri
	l.TplName = "login.html"
}

type LoginUser struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	RedirectUri string `json:"redirectUri"`
}

func (l *LoginController) Post() {
	var data LoginUser
	err := json.Unmarshal(l.Ctx.Input.RequestBody, &data)
	if err != nil {
		return
	}
	if data.Username == "gexin" && data.Password == "123456" {
		l.SetSession("beegoDemo", data.Username)
		data := map[string]string{"code": "200", "msg": "success", "redirectUri": data.RedirectUri}
		l.Data["json"] = data
		l.ServeJSON()
	} else {
		data := map[string]string{"code": "400", "msg": "fail", "redirectUri": data.RedirectUri}
		l.Data["json"] = data
		l.ServeJSON()
	}
}
