package models

import "github.com/astaxie/beego/orm"

type Labels struct {
	Id         int           `orm:"pk;auto"`
	Name       string        `orm:"column(name)"`
	ReportMany []*ReportMany `orm:"rel(m2m)"` // 正向关系，rel(many)
}

func (l *Labels) TableName() string {
	return "labels"
}

func init() {
	orm.RegisterModel(new(Labels))
}
