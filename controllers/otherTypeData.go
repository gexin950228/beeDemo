package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type XmlController struct {
	beego.Controller
}

func (x *XmlController) Get() {
	user := User{Id: 1, Name: "明哥", Age: 36, Address: "北京市"}
	// x.Data["xml"] = &user
	// x.ServeXML()
	// x.Data["json"] = &user
	// x.ServeJSON()
	// x.Data["jsonp"] = &user
	// x.ServeJSONP()
	x.Data["yaml"] = &user
	x.ServeYAML()
}

type FlashController struct {
	beego.Controller
}

func (f *FlashController) Get() {
	flash := beego.ReadFromRequest(&f.Controller)
	err := flash.Data["error"]
	not := flash.Data["notice"]
	fmt.Println(err)
	if len(err) != 0 {
		f.TplName = "error.html"
		fmt.Println("前端获取到没有用户名")
	} else if len(not) != 0 {
		fmt.Println("前端获取到密码不对")
		f.TplName = "notice.html"
	} else {
		fmt.Println("起始位置")
		f.TplName = "flash.html"
	}
}

func (f *FlashController) Post() {
	//  初始化flash
	flash := beego.NewFlash()
	fmt.Println("=================================================")
	name := f.Input().Get("name")
	age := f.Input().Get("age")
	fmt.Printf("%v, %v\n", name, age)
	if name == "" {
		flash.Error("用户名不能为空")
		flash.Store(&f.Controller)
		f.Redirect("/flash", 302)
	} else if age != "123" {
		flash.Notice("密码错误")
		flash.Store(&f.Controller)
		f.Redirect("/flash", 302)
	} else {
		flash.Success("成功")
		flash.Store(&f.Controller)
		f.Redirect("/flash", 302)
	}
}
