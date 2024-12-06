package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type GreetController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "gexin17812126257@163.com"
	c.TplName = "index.tpl"
}

func (g *GreetController) Get() {
	g.Data["Website"] = "gexin.me"
	g.Data["Email"] = "861439031@qq.com"
	g.TplName = "greet.tpl"
}

type User struct {
	Id      int64
	Name    string
	Age     int64
	Address string
}

func (u *UserController) Get() {
	var user User
	user.Address = "北京市"
	user.Age = 29
	user.Name = "葛新"
	arrs := []int{1, 2, 3, 4, 5, 6}
	u.Data["user"] = user
	u.Data["arrs"] = arrs
	arrStruct := [4]User{
		{Id: 1, Name: "葛新", Age: 29, Address: "北京市"},
		{Id: 2, Name: "周航", Age: 18, Address: "北京市"},
		{Id: 3, Name: "刘继雄", Age: 18, Address: "北京市"},
		{Id: 4, Name: "高林飞", Age: 18, Address: "北京市"}}
	u.Data["arrStruct"] = arrStruct
	mapc := map[string]interface{}{"name": "葛新", "age": 29, "gender": "male", "hobbies": []string{"羽毛球", "跑步", "骑行"}}
	u.Data["mapc"] = mapc
	// slice
	slice := []int{1, 3, 5, 6, 8, 17}
	u.Data["slice"] = slice
	u.TplName = "user.html"
}
