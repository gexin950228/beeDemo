package controllers

import (
	"beeDemo/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type ValicatorController struct {
	beego.Controller
}

func (c *ValicatorController) Get() {
	c.TplName = "validator.html"
}

type Validate struct {
	Name  string `form:"name" json:"name" binding:"name" valid:"Required"`
	Phone string `form:"phone" json:"phone" valid:"Required;Mobile" binding:"phone"`
	Age   int    `form:"age" json:"age" valid:"Numeric;Required;Min(0);Max(200)" binding:"age,min=0,max=200"`
	Email string `form:"email" json:"email" valid:"Email;Required" binding:"email"`
}

type ValidateResult struct {
	Code     int      `json:"code"`
	ErrorMsg []string `json:"errorMsg"`
}

func (c *ValicatorController) Post() {
	//name := c.GetString("name")
	//phone := c.GetString("phone")
	//age, _ := c.GetInt64("age")
	//fmt.Printf("获取的年龄类型: %T, 数字: %d\n", age, age)
	//email := c.GetString("email")

	// 解析到结构体
	var vData Validate
	c.ParseForm(&vData)
	//
	var MessagTpls = map[string]string{
		"Required": "不能为空",
		"MinSize":  "最短长度为 %d",
		"Length":   "长度必须为 %d",
		"Numeric":  "必须是有效的数字",
		"Email":    "无效的电子邮件地址",
		"Mobile":   "无效的手机号码",
	}
	validation.SetDefaultMessage(MessagTpls)
	valid := validation.Validation{}
	//valid.Required(vData.Name, "姓名").Message("不能为空")
	//valid.Email(vData.Email, "邮箱").Message("必须为正确的邮箱地址")
	//valid.Length(vData.Phone, 11, "电话号码").Message("必须为11位数字的电话号码")
	//valid.Numeric(age, "年龄").Message("必须为整数")
	//valid.Max(vData.Age, 200, "年龄").Message("不能大于100")
	//valid.Min(vData.Age, 0, "年龄").Message("年龄必须大于0")
	valid.Valid(vData)

	//// 判断有没有错误信息
	var validateError ValidateResult
	if valid.HasErrors() {
		validateError.Code = 0
		validateError.ErrorMsg = []string{}
		for _, err := range valid.Errors {
			fmt.Println(fmt.Sprintf("%s校验出错，validation error: %s", err.Key, err.Error()))
			validateError.ErrorMsg = append(validateError.ErrorMsg, err.Key+err.Message)
			fmt.Println(validateError.ErrorMsg)
			utils.LogToFile("Error", fmt.Sprintf("%s校验错误，错误信息: %v", err.Name, err.Error()))
			c.Data["json"] = validateError
			//c.ServeJSON()
		}
	} else {
		validateError.Code = 1
		validateError.ErrorMsg = append(validateError.ErrorMsg, "没有错误")
		c.Data["json"] = validateError
		//c.ServeJSON()
	}
	c.ServeJSON()
}
