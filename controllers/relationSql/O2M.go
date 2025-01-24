package relationSql

import (
	"beeDemo/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type O2M struct {
	beego.Controller
}

func (l *O2M) Get() {
	l.SetSession("person_id", "1")

	personId := l.GetSession("person_id")

	o := orm.NewOrm()
	var reports []models.ReportMany
	_, err := o.QueryTable("report_many").Filter("person_id", personId).Filter("is_deleted", 0).All(&reports)
	if err != nil {
		return
	}
	l.Data["reports"] = reports
	l.TplName = "sqlExe/o2m.html"
}
