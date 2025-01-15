package models

import "github.com/astaxie/beego/orm"

type Person struct {
	Id            int64          `orm:"pk;auto"`
	Name          string         `orm:"column(name);size(255)"`
	IdCardNumber  string         `orm:"column(id_card_number);size(20)"`
	PersonProfile *PersonProfile `orm:"reverse(one)"`
	ReportMany    []*ReportMany  `orm:"reverse(many)"`
}

func (p *Person) TableName() string {
	return "person"
}

type PersonProfile struct {
	Id                  int64   `orm:"pk;auto"`
	Gender              string  `orm:"size(20);description(性别)"`
	Age                 int64   `orm:"default(1);description(年龄)"`
	Email               string  `orm:"size(50);column(email);description(邮箱)"`
	Address             string  `orm:"size(50);column(address);description(地址)"`
	Person              *Person `orm:"rel(one)"`
	Hobbies             string  `orm:"size(50);column(hobbies);description(兴趣爱好)"`
	EducationExperience string  `orm:"column(education_experience);description(教育经历)"`
	Skills              string  `orm:"size(50);column(skills);description(技能)"`
}

func (pp *PersonProfile) TableName() string {
	return "person_profile"
}

func init() {
	orm.RegisterModel(new(Person), new(PersonProfile))
}
