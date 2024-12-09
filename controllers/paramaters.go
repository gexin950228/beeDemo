package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type User struct {
	Id      int64   `form:"id" json:"id" gorm:"id"`
	Name    string  `form:"username" json:"username" grom:"username"`
	Age     int64   `form:"age" json:"age" gorm:"age"`
	Address string  `form:"address" json:"address" gorm:"address"`
	Price   float64 `form:"price" json:"price" gorm:"price"`
}

type ParameterController struct {
	beego.Controller
}

func (p *ParameterController) Get() {
	// name := p.GetString("name")
	// fmt.Printf("前端传递的uri参数name: %s\n", name)
	// name1 := p.Input().Get("name")
	// fmt.Printf("Input.GET()方法获取到的name: %s\n", name1)
	// name3 := p.Ctx.Input.Param(":name")
	// fmt.Printf("Ctx.Input.Param()方法获取到的name: %s\n", name3)
	// name4 := p.Ctx.Input.Param(":name")
	// fmt.Printf("Ctx.Input.Param()方法获取到的name: %s\n", name4)
	p.TplName = "paramaters.html"
}

func (p *ParameterController) Post() {
	// name := p.GetString("username")
	// age, _ := p.GetInt("age")
	// isTrue, _ := p.GetBool("isTrue")
	// price, _ := p.GetFloat("price")
	// fmt.Println(name, age, isTrue, price)
	// paramaters := p.GetStrings("username")
	// fmt.Println(paramaters)
	// user := User{}
	// err := p.ParseForm(&user)
	// if err != nil {
	// 	fmt.Printf("user数据绑定出错： %v\n", err.Error())
	// } else {
	// 	fmt.Println("========================================================")
	// 	fmt.Println(user)
	// }
	//  ajax获取
	body := p.Ctx.Input.RequestBody
	user := User{}
	error := json.Unmarshal(body, &user)
	if error != nil {
		fmt.Printf("解析json数据失败, 错误： %v\n", error.Error())
	}
	// fmt.Printf("username: %v\n", user.Name)
	// fmt.Printf("user: %v\n", user)
	result := map[string]string{"code": "200", "msg": "success"}
	p.Data["json"] = result
	p.ServeJSON()
}
