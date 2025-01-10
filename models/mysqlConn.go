package models

import (
	"beeDemo/utils"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id      int    `orm:"pk;auto;unique"`
	Name    string `orm:"size(255)"`
	Address string `orm:"size(255)"`
}

func (u User) TableName() string {
	return "sys_user"
}

type SaveRegisterUser struct {
	Id       int64  `orm:"pk;auto"`
	Username string `json:"username" form:"username" orm:"column(username)" valid:"Required"`
	Email    string `json:"email" form:"email" valid:"Required"`
	Password string `json:"password" form:"password" orm:"column(password)" valid:"Required"`
}

func (s *SaveRegisterUser) Prepare() {
	orm.RegisterModel(new(SaveRegisterUser))
	orm.Debug = true
}

func (s *SaveRegisterUser) TableName() string {
	return "save_register_users"
}

func init() {
	mysqlConn := utils.LoadMysqlConfig()
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		utils.LogToFile("初始化数据库驱动失败，错误信息位：%s", err.Error())
	}
	dst := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", mysqlConn.User, mysqlConn.Password, mysqlConn.Host, mysqlConn.Port, mysqlConn.Database)
	err = orm.RegisterDataBase("default", "mysql", dst)
	if err != nil {
		utils.LogToFile("Error", fmt.Sprintf("数打开数据库链接失败，错误: %s", err.Error()))
		return
	}
	//orm.RegisterModel(new(LoginUser))
	orm.RegisterModel(new(SaveRegisterUser))

}
