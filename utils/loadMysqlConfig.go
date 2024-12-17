package utils

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type MysqlConn struct {
	Host     string `json:"host" form:"host"`
	Port     string `json:"port" form:"port"`
	User     string `json:"user" form:"user"`
	Password string `json:"password" form:"password"`
	Database string `json:"database" form:"database"`
}

func LoadMysqlConfig() MysqlConn {
	pwd, _ := os.Getwd()
	var mysqlConn MysqlConn
	cfgFile := filepath.Join(pwd, "conf/", "mysql.yaml")
	data, err := os.ReadFile(cfgFile)
	if err != nil {
		LogToFile("Error", fmt.Sprintf("数据库配置解析失败，错误为： %s", err.Error()))
	} else {
		err := yaml.Unmarshal(data, &mysqlConn)
		if err != nil {
			LogToFile("Error", fmt.Sprintf("数据库信息解析失败，错误: %s", err.Error()))
		}
	}
	return mysqlConn
}
