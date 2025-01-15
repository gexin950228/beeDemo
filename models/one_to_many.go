package models

import "github.com/astaxie/beego/orm"

type ReportMany struct {
	Id          int       `form:"id" json:"id" orm:"pk;auto;column(id)"`
	Title       string    `form:"title" json:"title" orm:"column(title)" valid:"Required"`
	Author      string    `form:"author" json:"author" orm:";column(author)" valid:"Required"`
	Description string    `json:"description" form:"description" orm:"column(description)" valid:"Required"`
	IsDeleted   int       `form:"is_deleted" json:"is_deleted" orm:"column(is_deleted)"`
	ReadCount   int64     `form:"read_count" json:"read_count" orm:"column(read_count)"`
	Classify    string    `form:"classify" json:"classify" orm:"column(classify)"`
	Person      *Person   `orm:"rel(fk);on_delete(do_nothing)"` // 级联删除操作，可选三个值，set_null, set_default, do_nothing
	Labels      []*Labels `orm:"reverse(many)"`                 // 反向关系，reverse(many)
}

func init() {
	orm.RegisterModel(new(ReportMany), new(User))
}

func (m *ReportMany) TableName() string {
	return "report_many"
}
