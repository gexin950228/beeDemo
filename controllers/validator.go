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
	Name  string `form:"name" json:"name" binding:"required"`
	Phone string `form:"phone" json:"phone" binding:"required"`
	Age   int64  `form:"age" json:"age" binding:"required"`
	Email string `form:"email" json:"email" binding:"required"`
}

type ValidateResult struct {
	Code     int      `json:"code"`
	ErrorMsg []string `json:"errorMsg"`
}

func (c *ValicatorController) Post() {
	name := c.GetString("name")
	phone := c.GetString("phone")
	age, _ := c.GetInt("age")
	fmt.Printf("获取的年龄类型: %T, 数字: %d\n", age, age)
	email := c.GetString("email")

	// 初始化validator并对name进行校验
	valid := validation.Validation{}
	var validateError ValidateResult
	valid.Required(name, "姓名").Message("不能为空")
	valid.Email(email, "邮箱").Message("必须为正确的邮箱地址")
	valid.Length(phone, 11, "电话号码").Message("必须为11位数字的电话号码")
	valid.Numeric(age, "年龄").Message("必须为整数")
	valid.Max(age, 100, "年龄").Message("不能大于100")

	// 判断有没有错误信息
	if valid.HasErrors() {
		validateError.Code = 0
		validateError.ErrorMsg = []string{}
		for _, err := range valid.Errors {
			fmt.Println(fmt.Sprintf("%s校验出错，validation error: %s", err.Key, err.Error()))
			validateError.ErrorMsg = append(validateError.ErrorMsg, err.Key+err.Message)
			fmt.Println(validateError.ErrorMsg)
			utils.LogToFile("Error", fmt.Sprintf("%s校验错误，错误信息: %v", err.Name, err.Error()))
		}
	} else {
		validateError.Code = 1
		validateError.ErrorMsg = append(validateError.ErrorMsg, "没有错误")
	}
	fmt.Println(name, phone, age, email)

	c.Data["json"] = validateError
	c.ServeJSON()
}
