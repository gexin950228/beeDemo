package relationSql

import (
	"github.com/astaxie/beego"
)

type ManyToManyController struct {
	beego.Controller
}

func (m *ManyToManyController) Get() {
	//o := orm.NewOrm()
	//reportMany := models.ReportMany{Id: 1}
	//labels := models.Labels{Name: "beego技术详解"}
	//o.Insert(&labels)
	//labels := models.Labels{Id: 2}
	//labels := models.Labels{Id: 1}
	//m2mMgr := o.QueryM2M(&labels, "ReportMany")
	//m2mMgr.Add(&reportMany)

	// remove
	//m2mMgr.Remove(&labels)
	//labels := models.Labels{Id: 1}
	//m2mMgr.Remove(&labels)
	//m2mMgr.Clear()

	// count
	//reportsMany := models.ReportMany{Id: 2}
	//
	//m2m := o.QueryM2M(&reportsMany, "Labels")
	//count, _ := m2m.Count()
	//fmt.Println(count)

	// Update
	//lid := 3
	//rid := 5
	//qs := o.QueryTable("labels_report_manys")
	//_, err := qs.Filter("labels_id", lid).Filter("report_many_id", rid).Update(orm.Params{
	//	"labels_id":      4,
	//	"report_many_id": 3,
	//})
	//if err != nil {
	//	return
	//}

	// 查询： 已知report_id查询所有绑定的label_id, 已知labels_id查询reports_id

	//rid := 2
	//var reports models.ReportMany
	//qs := o.QueryTable("report_many")
	//qs.Filter("id", rid).One(&reports)
	//o.LoadRelated(&reports, "Labels")
	//for _, label := range reports.Labels {
	//	fmt.Println(label.Name)
	//}

	//lid := 1
	//qs := o.QueryTable("report_many")
	//var reportMany []models.ReportMany
	//qs.Filter("Labels__labels__id", lid).All(&reportMany)
	//fmt.Println(reportMany)

	// 级联删除

	m.TplName = "many_to_many.html"
}
