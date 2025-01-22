package relationSql

import (
	"beeDemo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type O2O struct {
	beego.Controller
}

func (el *O2O) Get() {
	el.TplName = "sqlExe/o2o.html"
}

func (el *O2O) Post() {
	var user models.Person
	var userInDatabase models.Person
	var up models.PersonProfile
	//err := el.ParseForm(&user)
	//if err != nil {
	//	utils.LogToFile("ERROR", fmt.Sprintf("用户信息绑定失败，错误信息: %s", err.Error()))
	//	fmt.Println(err.Error())
	//}
	user.Name = el.GetString("userName")
	user.IdCardNumber = el.GetString("userIdCardNumber")
	age, err := el.GetInt64("age")
	up.Age = age
	up.Gender = el.GetString("gender")
	up.Email = el.GetString("email")
	up.Hobbies = el.GetString("hobbies")
	up.Skills = el.GetString("skills")
	up.EducationExperience = el.GetString("education_experience")
	up.Address = el.GetString("address")
	o := orm.NewOrm()
	qs := o.QueryTable("person")
	err = qs.Filter("id_card_number", user.IdCardNumber).One(&userInDatabase)
	if err != nil {
		fmt.Println(err)
	}
	if userInDatabase.Id > 0 {
		fmt.Printf("身份证号码为%s的用户存在", userInDatabase.IdCardNumber)
	} else {
		fmt.Printf("身份证号码为%s的用户不存在", userInDatabase.IdCardNumber)
	}
	_, err = o.InsertOrUpdate(&userInDatabase, "id_card_number")
	if err != nil {
		fmt.Printf("插入数据库出错，错误信息为: %s", err.Error())
	}
	up.Person = &models.Person{Id: userInDatabase.Id, Name: userInDatabase.Name, IdCardNumber: userInDatabase.IdCardNumber}
	_, err = o.InsertOrUpdate(&up, "person_id")
	el.TplName = "sqlExe/o2o.html"
}
