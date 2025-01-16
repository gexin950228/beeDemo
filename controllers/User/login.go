package User

import (
	"beeDemo/models"
	"beeDemo/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
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
	Email          string `json:"email" form:"email"`
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
		sendVerifyCodeResult = SendVerifyCodeResult{Code: 0, Msg: "发送邮箱验证码出错"}
	}
	rand.Seed(time.Now().UnixNano())
	verifyCode := strconv.Itoa(rand.Intn(999999))
	sendVerifyCodeResult = SendVerifyCode(data.Username, verifyCode, data.Email, data.VerifyCodeType)
	redisInfo := utils.RedisInfo{
		Key:        fmt.Sprintf("%s_%s_code", data.Email, data.VerifyCodeType),
		Value:      verifyCode,
		ExpireTime: 120,
	}
	redisConfig := utils.LoadRedisConfig()
	saveResult := utils.SaveToRedis(redisConfig, redisInfo)
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
	fmt.Println("redirectUri:", redirectUri)
	l.Data["redirectUri"] = redirectUri
	l.TplName = "user/login.html"
}

type LoginUser struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	RedirectUri    string `json:"redirectUri"`
	VerifyCode     string `json:"verifyCode"`
	Email          string `json:"email" orm:"column(email)"`
	VerifyCodeType string `json:"verify_code_type"`
}

type RepData struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	RedirectUri string `json:"redirectUri"`
}

type CustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (l *LoginController) Post() {
	var loginData LoginUser
	var userInfo models.SaveRegisterUser
	err := json.Unmarshal(l.Ctx.Input.RequestBody, &loginData)
	if err != nil {
		utils.LogToFile("Error", "解析用户数据出错")
		utils.LogToFile("Error", err.Error())
	}
	fmt.Printf("redirectUri: %s\n", loginData.RedirectUri)
	redisKey := fmt.Sprintf("%s_%s_code", loginData.Email, loginData.VerifyCodeType)
	searchRedisResult := utils.SearchRedis(redisKey)
	if searchRedisResult.Code != 1 {
		utils.LogToFile("Error", fmt.Sprintf("%s查询验证码失败", loginData.Username))
	}
	o := orm.NewOrm()
	var repData RepData
	err = o.QueryTable("save_register_users").Filter("email", loginData.Email).One(&userInfo)
	if err != nil {
		utils.LogToFile("Error", fmt.Sprintf("用户%s登陆失败，错误信息：%s", loginData.Username, err.Error()))
		repData.Code = "1"
		repData.Msg = err.Error()
	} else {
		hashedPassword := base64.StdEncoding.EncodeToString([]byte(loginData.Password))
		if hashedPassword != userInfo.Password {
			repData.Code = "1"
			repData.Msg = "密码校验出错"
		} else {
			repData.Code = "0"
			repData.Msg = "用户名密码校验成功"
			repData.RedirectUri = loginData.RedirectUri
			l.SetSecureCookie("session_token", loginData.Email, userInfo.Password)
			redisInfo := utils.RedisInfo{
				Key:        fmt.Sprintf("%s_login", loginData.Email),
				Value:      "true",
				ExpireTime: 86400,
			}
			redisConn := utils.RedisConn{}
			utils.SaveToRedis(redisConn, redisInfo)
			l.Data["json"] = repData
			l.ServeJSON()
		}
	}
}
