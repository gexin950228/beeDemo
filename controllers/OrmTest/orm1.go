package OrmTest

import (
	"beeDemo/models"
	"beeDemo/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type TestOrmController struct {
	beego.Controller
}

func (t *TestOrmController) Get() {
	o := orm.NewOrm()
	u := models.LoginUser{Username: "诚征", Address: "北京市海淀区", LastLoginTime: time.Now(), RedirectUri: "/parameter", Password: "Ch1nZh1nhg$a"}
	id, err := o.Insert(&u)
	if err != nil {
		fmt.Println(err.Error())
		utils.LogToFile("Error", fmt.Sprintf("插入数据错误， error: %s", err.Error()))
	} else {
		fmt.Println(id)
	}
	mysqlConn := utils.LoadMysqlConfig()
	fmt.Println(mysqlConn)
	t.TplName = "testOrm/testorm1.html"
}
