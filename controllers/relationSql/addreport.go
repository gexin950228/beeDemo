package relationSql

import (
	"beeDemo/models"
	"beeDemo/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type AddReport struct {
	beego.Controller
}

func (c *AddReport) Get() {
	c.SetSession("id", "1")
	uid := c.GetSession("id")
	c.Data["uid"] = uid
	c.TplName = "sqlExe/addreport.html"
}

func (c *AddReport) Post() {
	uid := 1
	o := orm.NewOrm()
	var person models.Person
	err := o.QueryTable(`person`).Filter("id", uid).One(&person)
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("查询person出错，错误信息:%v", err))
	}
	var report models.ReportMany
	title := c.GetString("title")
	report.Title = title
	author := c.GetString("author")
	report.Author = author
	description := c.GetString("description")
	report.Description = description
	classify := c.GetString("classify")
	report.Classify = classify
	readCount := c.GetString("read_count")
	parseInt, err := strconv.ParseInt(readCount, 10, 64)
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("获取的id转化为int出错，错误信息:%v", err))
	}
	report.ReadCount = parseInt
	report.Person = &person
	_, err = o.Insert(&report)
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("获取的id转化为int出错，错误信息:%v", err))
	}
	c.TplName = "sqlExe/addreport.html"
}

type UpdateReport struct {
	beego.Controller
}

func (c *UpdateReport) Get() {
	o := orm.NewOrm()
	id := c.GetString("id")
	var report models.ReportMany
	reportId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.LogToFile("Error", fmt.Sprintf("更新文章时获取文章信息失败，错误信息为： %s", err.Error()))
	}
	err = o.QueryTable("report_many").Filter("id", int(reportId)).One(&report)
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("更新文章查询文章数据出错 %s", err.Error()))
	}
	c.Data["report_many"] = report
	c.TplName = "sqlExe/updatereport.html"
}

func (c *UpdateReport) Post() {
	id := c.GetString("id")
	var reportMany models.ReportMany
	var person models.Person
	reportId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("获取report id出错，错误信息:%v", err))
	}
	o := orm.NewOrm()
	o.QueryTable("report_many").Filter("id", reportId).One(&reportMany)
	personId := reportMany.Person.Id
	person.Id = personId
	o.QueryTable("person").Filter("id", personId).One(&person)
	title := c.GetString("title")
	description := c.GetString("description")
	classify := c.GetString("classify")
	read_count := c.GetString("read_count")
	readCount, err := strconv.Atoi(read_count)
	reportMany.Title = title
	reportMany.Description = description
	reportMany.Classify = classify
	reportMany.ReadCount = int64(readCount)
	reportMany.Person = &person
	_, err = o.Update(&reportMany)
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("更新person信息出错，错误信息:%v", err))
	}
	c.Redirect("/e12m", 302)
}

type DeleteReport struct {
	beego.Controller
}

func (c *DeleteReport) Get() {
	o := orm.NewOrm()
	var deleteReport models.ReportMany
	id := c.GetString("id")
	reportId, _ := strconv.ParseInt(id, 10, 64)
	deleteReport.Id = int(reportId)
	qs := o.QueryTable("report_many").Filter("id", reportId)
	err := qs.One(&deleteReport)
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("查询数据库出错，错误信息问： %s\n", err.Error()))
	}
	c.Data["report_many"] = &deleteReport
	c.TplName = "sqlExe/deletereport.html"
}

func (c *DeleteReport) Post() {
	var report models.ReportMany
	id := c.GetString("id")
	reportId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("获取的删除report id出错，错误信息:%v", err))
	}
	report.Id = int(reportId)
	o := orm.NewOrm()
	_, err = o.QueryTable("ReportMany").Filter("id", reportId).Update(orm.Params{"is_deleted": true})
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("获取的id转化为int出错，错误信息:%v", err))
	}
	c.Redirect("/e12m", 302)
}
