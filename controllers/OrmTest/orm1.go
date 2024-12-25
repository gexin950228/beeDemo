package OrmTest

import (
	"beeDemo/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type TestOrmController struct {
	beego.Controller
}

type User struct {
	Id            int       `orm:"pk;auto"`
	Username      string    `orm:"size(500)"`
	Address       string    `orm:"size(500)"`
	LastLoginTime time.Time `orm:"auto_now_add;type(datetime)"`
	RedirectUri   string    `orm:"size(500)"`
	Password      string    `orm:"size(500)"`
}

func (u *User) TableName() string {
	return "sys_user"
}

func init() {
	mysqlConn := utils.LoadMysqlConfig()
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dst := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", mysqlConn.User, mysqlConn.Password, mysqlConn.Host, mysqlConn.Port, mysqlConn.Database)
	orm.RegisterDataBase("default", "mysql", dst)
	orm.RegisterModel(new(User))
}

func (t *TestOrmController) Get() {
	o := orm.NewOrm()
	u := User{Username: "葛新", Address: "北京市海淀区", LastLoginTime: time.Now(), RedirectUri: "/parameter", Password: "Ch1nZh1nhg$a"}
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
