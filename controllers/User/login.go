package User

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

func SendVerifyCode(username, verifyCode, email, verifyCodeType string) SendVerifyCodeResult {
	var result SendVerifyCodeResult
	mailContext := utils.MailContext{
		From:        "861439031@qq.com",
		To:          email,
		Subject:     "[beedemo登录验证码]",
		Body:        fmt.Sprintf("<h1>%s的%s验证码是： %s，过期时间120秒。</h1>", username, verifyCodeType, verifyCode),
		ContextType: "text/html",
	}
	utils.SendMail(mailContext)
	return result
}

type SendVerifyCodeController struct {
	beego.Controller
}

type VerifyData struct {
	Username       string `form:"username" json:"username"`
	Mail           string `json:"email" form:"email"`
	VerifyCodeType string `json:"verify_code_type" form:"verify_code_type"`
}

func (s *SendVerifyCodeController) Prepare() {
}

func (s *SendVerifyCodeController) Post() {
	var data VerifyData
	var sendVerifyCodeResult SendVerifyCodeResult
	errUrmarshal := json.Unmarshal(s.Ctx.Input.RequestBody, &data)
	if errUrmarshal != nil {
		utils.LogToFile("Error", errUrmarshal.Error())
		fmt.Println("Error:", errUrmarshal.Error())
		sendVerifyCodeResult = SendVerifyCodeResult{Code: 0, Msg: "发送邮箱验证码出错"}
	} else {
		fmt.Println(data)
	}
	rand.Seed(time.Now().UnixNano())
	verifyCode := strconv.Itoa(rand.Intn(999999))
	fmt.Println(verifyCode)
	sendVerifyCodeResult = SendVerifyCode(data.Username, verifyCode, data.Mail, data.VerifyCodeType)
	redisInfo := utils.RedisInfo{
		Key:        fmt.Sprintf("%s_%s_code", data.Mail, data.VerifyCodeType),
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
	l.TplName = "user/login.html"
}

type LoginUser struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	RedirectUri    string `json:"redirectUri"`
	VerifyCode     string `json:"verifyCode"`
	VerifyCodeType string `json:"verify_code_type"`
}

func (l *LoginController) Post() {
	var loginData LoginUser
	err := json.Unmarshal(l.Ctx.Input.RequestBody, &loginData)
	if err != nil {
		utils.LogToFile("Error", "解析用户数据出错")
		utils.LogToFile("Error", err.Error())
	}
	redisKey := fmt.Sprintf("%s%s", loginData.Username, loginData.VerifyCodeType)
	searchRedisResult := utils.SearchRedis(redisKey)
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
		//fmt.Println("登录成功")
		l.Data["json"] = data
		l.ServeJSON()
	} else {
		//fmt.Println("登录失败")
		data := map[string]string{"code": "400", "msg": "fail", "redirectUri": loginData.RedirectUri}
		l.Data["json"] = data
		utils.LogToFile("Info", "用户登录出错")
		l.ServeJSON()
	}
}
