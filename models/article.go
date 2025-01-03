package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id        int    `form:"id" json:"id" orm:"pk;auto;column(id)"`
	Title     string `form:"title" json:"title" orm:"column(title)" valid:"Required"`
	Author    string `form:"author" json:"author" orm:";column(author)" valid:"Required"`
	Desc      string `json:"desc" form:"desc" orm:"column(desc)" valid:"Required"`
	IsDeleted int    `form:"is_deleted" json:"is_deleted" orm:"column(is_deleted)"`
}

func init() {
	orm.RegisterModel(new(Article))
}

func (a *Article) TableName() string {
	return "article"
}
