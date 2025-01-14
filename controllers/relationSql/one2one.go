package relationSql

import (
	"beeDemo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type OneToOneController struct {
	beego.Controller
}

type rep struct {
	Code string
	Msg  []string
}

func (c *OneToOneController) Get() {
	orm.Debug = true
	person := models.Person{
		Id:           2,
		Name:         "葛香",
		IdCardNumber: "43038119871013361X",
	}
	o := orm.NewOrm()
	//o.Insert(&person)
	personProfile := models.PersonProfile{
		Id:                  2,
		Age:                 37,
		Address:             "湖南省湘潭市",
		Hobbies:             "看剧",
		Email:               "email@qq.com",
		EducationExperience: "高中",
		Person:              &person,
		Skills:              "美甲",
	}
	//o.Insert(&personProfile)
	// 原生sql
	ret := o.Raw("INSERT INTO person(id, name, id_card_number) VALUES(?, ?, ?)", person.Id, person.Name, person.IdCardNumber)
	_, err := ret.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}
	r2 := o.Raw("INSERT INTO person_profile(id, gender, age, address, hobbies, email, education_experience, person_id, skills) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)",
		person.Id, personProfile.Gender, personProfile.Age, personProfile.Address, personProfile.Hobbies, personProfile.Email, personProfile.EducationExperience, &person.Id, personProfile.Skills)
	_, err = r2.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}
	c.TplName = "relationSql/oneToOne.html"
}
