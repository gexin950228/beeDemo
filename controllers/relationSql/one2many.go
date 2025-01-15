package relationSql

import (
	"beeDemo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type OneToMany struct {
	beego.Controller
}

func (m *OneToMany) Get() {
	o := orm.NewOrm()
	//qs := o.QueryTable("report_many")
	// 插入
	//person := models.Person{
	//	Name:         "葛禄辉",
	//	IdCardNumber: "4004111967092998762",
	//}
	//_, err := o.Insert(&person)
	//if err != nil {
	//	return
	//}
	//report := models.ReportMany{Title: "title1", Description: "description1", Author: "葛吕辉", ReadCount: 10, IsDeleted: 0, Classify: "demo1", Person: &person}
	//_, err = o.Insert(&report)
	//if err != nil {
	//	return
	//}

	// 一对一更新
	//qs := o.QueryTable("report_many")
	//update, err := qs.Filter("id__exact", 1).Update(orm.Params{"person_id": 1})
	//if err != nil {
	//	utils.LogToFile("ERROR", fmt.Sprintf("一对多更新数据出错，错误信息: %v", err))
	//} else {
	//	utils.LogToFile("INFO", fmt.Sprintf("一对多更新数据成功，修改的数据id为: %d", update))
	//}

	// 一对多删除
	//_, err := qs.Filter("title__exact", "title2").Delete()
	//if err != nil {
	//	return
	//}

	// 一对多查询
	//var reports []models.ReportMany
	//count, err := qs.Filter("id", 1).RelatedSel().All(&reports)
	//if err != nil {
	//	return
	//} else {
	//	fmt.Println("================================================")
	//	fmt.Println(count)
	//	fmt.Println(reports.Person)
	//}
	//uid := 1
	//_, err := o.QueryTable("report_many").Filter("person_id", uid).All(&reports)
	//if err != nil {
	//	return
	//}
	//fmt.Println(reports)

	// 一直report差查询作者
	rid := 1
	var person models.Person
	qs := o.QueryTable("person")
	err := qs.Filter("ReportMany__Id", rid).One(&person)
	if err != nil {
		fmt.Println("======================")
		fmt.Println(err)
	}
	fmt.Println(person)

	m.TplName = "one_to_many.html"
}
