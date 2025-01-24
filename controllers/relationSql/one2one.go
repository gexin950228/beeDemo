package relationSql

import (
	"beeDemo/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type OneToOneController struct {
	beego.Controller
}

func (c *OneToOneController) Get() {
	orm.Debug = true
	o := orm.NewOrm()

	//person := models.Person{
	//	Name:         "葛禄辉",
	//	IdCardNumber: "43038119871013361X",
	//}
	//
	//_, err := o.Insert(&person)
	//if err != nil {
	//	return
	//}
	//
	//personProfile := models.PersonProfile{
	//	Age:                 57,
	//	Address:             "湖南省湘潭市",
	//	Hobbies:             "看剧",
	//	Email:               "geluhui@qq.com",
	//	EducationExperience: "初中",
	//	Person:              &person,
	//	Skills:              "没有爱好",
	//}
	//_, err = o.Insert(&personProfile)
	//if err != nil {
	//	return
	//}
	//
	//// 原生sql
	//ret := o.Raw("INSERT INTO person( name, id_card_number) VALUES(?, ?)", person.Name, person.IdCardNumber)
	//exec, err := ret.Exec()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//var insertId int64
	//insertId, err = exec.LastInsertId()
	//r2 := o.Raw("INSERT INTO person_profile(gender, age, address, hobbies, email, education_experience, person_id, skills)"+
	//	" VALUES(?, ?, ?, ?, ?, ?, ?, ?)", personProfile.Gender, personProfile.Age, personProfile.Address, personProfile.Hobbies, personProfile.Email, personProfile.EducationExperience, insertId, personProfile.Skills)
	//exec, err = r2.Exec()
	//if err != nil {
	//	return
	//}
	//id, _ := exec.LastInsertId()
	//fmt.Println(id)

	// 一对一修改
	//
	//update, err := qs.Filter("person_id__exact", 3).Update(orm.Params{"age": 57, "hobbies": "看小说"})
	//if err != nil {
	//	return
	//} else {
	//	fmt.Println(update)
	//}

	// 一对一删除orm
	//i, err := qs.Filter("id__exact", 3).Delete()
	//if err != nil {
	//	return
	//} else {
	//	utils.LogToFile("INFO", fmt.Sprintf("id为%d的用户已删除", i))
	//}
	//  一对一删除原生sql
	//r := o.Raw("delete from person_profile where id=?", 8)
	//execId, err := r.Exec()
	//if err != nil {
	//	return
	//}
	//r1 := o.Raw("DELETE FROM person where id=?", 8)
	//_, err := r1.Exec()
	//if err != nil {
	//	return
	//}

	var personProfile models.PersonProfile
	pid := 1
	qs1 := o.QueryTable("person_profile")
	err := qs1.Filter("id__exact", pid).RelatedSel().One(&personProfile)
	if err != nil {
		return
	}
	c.TplName = "relationSql/oneToOne.html"
}
