package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type UserInfo struct {
	//修改源码： cmd_utils.go的getColumnAddQuery()最后添加fi.description,;
	Id                         int64     `orm:"primary_key;AUTO_INCREMENT"`
	UserIcon                   string    `orm:"type:varchar(255);column(user_icon);null;description(头像)"`
	Gender                     int       `orm:"column(gender);description(性别,男:1 女:2);default(1)"`
	Vocation                   string    `orm:"column(vocation);size(50);description(职业)"`
	EducationBackground        string    `orm:"size(100);column(education_background);null;description(教育背景)"`
	SelfComment                string    `orm:"column(self_comment);size(500);null;description(自我评价)"`
	Salary                     float64   `orm:"column(salary);digits(10);decimals(2);null;description(薪资)"`
	Skills                     string    `orm:"column(skills);size(500);null;description(技能)"`
	RegisterTime               time.Time `orm:"auto_now;column(register_time);type(datetime)"`
	EmergencyContactor         string    `orm:"column(emergency_contactor);description(紧急联系人)"`
	EmergencyContactorRelation string    `orm:"column(emergency_contactor_relation);description(紧急联系人关系)"`
	EmergencyContactorPhone    string    `orm:"column(emergency_contactor_phone);description(紧急联系人电话)"`
}

func init() {
	orm.RegisterModel(new(UserInfo))
}

func (userinfo *UserInfo) TableName() string {
	return "user_info"
}
