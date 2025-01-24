package utils

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Mail struct {
	Host     string `form:"host" json:"host" binding:"required" yaml:"host"`
	Password string `form:"password" json:"password" binding:"required"`
	Account  string `form:"account" json:"account" binding:"required"`
	Auth     string `form:"auth" json:"auth" binding:"required"`
	Addr     string `form:"addr" json:"addr" binding:"required" yaml:"addr"`
	Port     int    `form:"port" json:"port" binding:"required" yaml:"port"`
}

func LoadConfig() (mail Mail) {
	pwd, _ := os.Getwd()
	cfgFile := filepath.Join(pwd, "conf/", "mail.yaml")
	data, err := os.ReadFile(cfgFile)
	if err != nil {
		LogToFile(err.Error(), "LoadConfig")
	} else {
		err := yaml.Unmarshal(data, &mail)
		if err != nil {
			LogToFile(err.Error(), "LoadConfig")
			return
		}
	}
	return mail
}
