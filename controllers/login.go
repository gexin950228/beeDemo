package controllers

import (
	"beeDemo/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"math/rand"
	"strconv"
	"time"
)

type LoginController struct {
	beego.Controller
}

type SendVerifyCodeResult struct {
	Code int    `json:"code" gorm:"code"`
	Msg  string `json:"msg" gorm:"msg"`
}

func SendVerifyCode(verifyCode string) SendVerifyCodeResult {
	var result SendVerifyCodeResult
	mailContext := utils.MailContext{
		From:        "861439031@qq.com",
		To:          "gexin17812126257@163.com",
		Subject:     "[beedemo登录验证码]",
		Body:        fmt.Sprintf("<h1>验证码是： %s，过期时间120秒。</h1>", verifyCode),
		ContextType: "text/html",
	}
	utils.SendMail(mailContext)
	return result
}

type SendVerifyCodeController struct {
	beego.Controller
}

type VerifyData struct {
	Username string `json:"username"`
	Mail     string `json:"mail"`
}

func (s *SendVerifyCodeController) Prepare() {
}

func (s *SendVerifyCodeController) Post() {
	var data VerifyData
	var sendVerifyCodeResult SendVerifyCodeResult
	errUrmarshal := json.Unmarshal(s.Ctx.Input.RequestBody, &data)
	if errUrmarshal != nil {
		utils.LogToFile("Error", errUrmarshal.Error())
		sendVerifyCodeResult = SendVerifyCodeResult{Code: 0, Msg: "发送邮箱验证码出错"}
	} else {
		sendVerifyCodeResult = SendVerifyCodeResult{Code: 0, Msg: "发送邮箱验证码出错"}
	}
	rand.Seed(time.Now().UnixNano())
	verifyCode := strconv.Itoa(rand.Intn(999999))
	fmt.Println(verifyCode)
	sendVerifyCodeResult = SendVerifyCode(verifyCode)
	redisInfo := utils.RedisInfo{
		Key:        data.Username,
		Value:      verifyCode,
		ExpireTime: 120,
	}
	redisConfig := utils.LoadRedisConfig()
	saveResult := utils.SaveToRedis(redisConfig, redisInfo)
	fmt.Println(saveResult)
	if saveResult.Code != 1 {
		sendVerifyCodeResult.Code = saveResult.Code
		sendVerifyCodeResult.Msg = saveResult.Msg
	} else {
		sendVerifyCodeResult.Code = saveResult.Code
		sendVerifyCodeResult.Msg = "验证码发送成功"
	}
	s.Data["json"] = sendVerifyCodeResult
	s.ServeJSON()
}

func (l *LoginController) Get() {
	redirectUri := l.GetString("redirectUri")
	l.Data["redirectUri"] = redirectUri
	l.TplName = "login.html"
}

type LoginUser struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	RedirectUri string `json:"redirectUri"`
	VerifyCode  string `json:"verifyCode"`
}

func (l *LoginController) Post() {
	var loginData LoginUser
	err := json.Unmarshal(l.Ctx.Input.RequestBody, &loginData)
	if err != nil {
		utils.LogToFile("Error", "解析用户数据出错")
		utils.LogToFile("Error", err.Error())
	}
	searchRedisResult := utils.SearchRedis(loginData.Username)
	if searchRedisResult.Code != 1 {
		utils.LogToFile("Error", fmt.Sprintf("%s查询验证码失败", loginData.Username))
	}

	if loginData.Username == "gexin" && loginData.Password == "123456" && loginData.VerifyCode == searchRedisResult.RedisResult {
		redisConn := utils.LoadRedisConfig()
		redisInfo := utils.RedisInfo{Key: loginData.Username + "loginStatus", Value: "Logined", ExpireTime: 86400}
		utils.SaveToRedis(redisConn, redisInfo)
		data := map[string]string{"code": "200", "msg": "success", "redirectUri": loginData.RedirectUri}
		//l.Data["json"] = data
		utils.LogToFile("Info", "用户登录成功")
		l.SetSession("loginUser", loginData.Username)
		l.SetSecureCookie("loginUser", loginData.Username, "", 30, redisConn)
		fmt.Println("登录成功")
		l.Data["json"] = data
		l.ServeJSON()
	} else {
		fmt.Println("登录失败")
		data := map[string]string{"code": "400", "msg": "fail", "redirectUri": loginData.RedirectUri}
		l.Data["json"] = data
		utils.LogToFile("Info", "用户登录出错")
		l.ServeJSON()
	}
}
