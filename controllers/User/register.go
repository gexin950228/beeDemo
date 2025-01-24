package User

import (
	"beeDemo/models"
	"beeDemo/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"regexp"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "user/register.html"
}

type RegisterInfo struct {
	Username       string `json:"username" form:"username" orm:"column(username)" valid:"Required"`
	Email          string `json:"email" form:"email" valid:"Required"`
	Password       string `json:"password" form:"password" orm:"column(password)" valid:"Required"`
	RepeatPassword string `json:"repeat_password" form:"repeat_password" orm:"column(repeat_password)" valid:"Required"`
	VerifyCode     string `json:"verify_code" form:"verify_code"`
}

type ResponseInfo struct {
	Code    int      `json:"code"`
	Msg     []string `json:"msg"`
	RegUser string   `json:"reg_user"`
}

func validateEmail(email string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(email)
}

func (c *RegisterController) Post() {
	var registerInfo RegisterInfo
	var repInfo ResponseInfo
	repInfo.Code = 0

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &registerInfo)
	// 数据校验
	if err != nil {
		utils.LogToFile("Error", fmt.Sprintf("注册信息绑定结构体失败： 错误信息为: %s", err))
		repInfo.Code = 1
		repInfo.Msg = append(repInfo.Msg, err.Error())
		c.Data["json"] = repInfo
		c.ServeJSON()
	}
	if registerInfo.Password != registerInfo.RepeatPassword {
		repInfo.Code = 1
		repInfo.Msg = append(repInfo.Msg, "两次输入的密码不一致")
		c.Data["json"] = repInfo
		c.ServeJSON()
	}
	if registerInfo.VerifyCode == "" {
		repInfo.Code = 1
		repInfo.Msg = append(repInfo.Msg, "验证码不能为空")
		c.Data["json"] = repInfo
		c.ServeJSON()
	}
	if registerInfo.Username == "" {
		repInfo.Code = 1
		repInfo.Msg = append(repInfo.Msg, "用户名不能为空")
		c.Data["json"] = repInfo
		c.ServeJSON()
	}
	if !validateEmail(registerInfo.Email) {
		repInfo.Code = 1
		repInfo.Msg = append(repInfo.Msg, "用户提交的邮箱格式不正确")
		c.Data["json"] = repInfo
		c.ServeJSON()
	}
	repInfo.RegUser = registerInfo.Username
	var searchResult utils.SearchRedisResult
	key := fmt.Sprintf("%s_register_code", registerInfo.Email)
	searchResult = utils.SearchRedis(key)
	if searchResult.Code == 0 {
		repInfo.Code = 1
		repInfo.Msg = append(repInfo.Msg, "验证码校验失败,请重新获取验证码后输入")
		repInfo.Msg = append(repInfo.Msg, "输入的验证码不正确")
		c.Data["json"] = repInfo
		c.ServeJSON()
	} else {
		if registerInfo.VerifyCode != searchResult.RedisResult {
			repInfo.Msg = append(repInfo.Msg, "输入的验证码不正确")
			c.Data["json"] = repInfo
			c.ServeJSON()
		}
	}
	o := orm.NewOrm()
	var saveRegisterUsers models.SaveRegisterUser
	saveRegisterUsers.Email = registerInfo.Email
	qs := o.QueryTable("save_register_users")
	exist := qs.Filter("email__exact", registerInfo.Email).Exist()
	if exist {
		repInfo.Code = 1
		repInfo.Msg = append(repInfo.Msg, fmt.Sprintf("邮箱%s已经被注册,请更换邮箱。", registerInfo.Email))
		c.Data["json"] = repInfo
		c.ServeJSON()
	} else {
		// 数据写入
		saveRegisterUsers.Username = registerInfo.Username
		saveRegisterUsers.Email = registerInfo.Email
		passwordByte := []byte(registerInfo.Password)
		hashedPassword := base64.StdEncoding.EncodeToString(passwordByte)
		saveRegisterUsers.Password = string(hashedPassword)
		_, err = o.Insert(&saveRegisterUsers)
		if err != nil {
			repInfo.Code = 1
			repInfo.Msg = append(repInfo.Msg, fmt.Sprintf("注册信息数据写入数据库失败，错误信息: %s", err))
		}
		c.Data["json"] = repInfo
		c.ServeJSON()
	}
}
