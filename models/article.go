package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id     int    `json:"id" orm:"id:pk";column(id)`
	Title  string `json:"title" orm:"column(title)""`
	Author string `json:"author" orm:";column(author)""`
	Desc   string `json:"desc" orm:"column(desc)""`
}

func init() {
	orm.RegisterModel(new(Article))
}

func (a *Article) TableName() string {
	return "article"
}
