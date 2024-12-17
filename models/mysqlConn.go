package models

import (
	"beeDemo/utils"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	mysqlConn := utils.LoadMysqlConfig()
	fmt.Println(mysqlConn)
	fmt.Println(mysqlConn)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dst := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", mysqlConn.User, mysqlConn.Password, mysqlConn.Host, mysqlConn.Port, mysqlConn.Database)
	err := orm.RegisterDataBase("default", "mysql", dst)
	if err != nil {
		utils.LogToFile("Error", fmt.Sprintf("数据库初始化失败，错误: %s", err.Error()))
		return
	}
}
