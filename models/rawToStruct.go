package models

import "github.com/astaxie/beego/orm"

type RawToStruct struct {
	Id    int64
	Total string
	Page  string
}

func init() {
	orm.RegisterModel(new(RawToStruct))
	orm.Debug = true
}
