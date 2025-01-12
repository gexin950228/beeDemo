package OrmTest

import (
	"github.com/astaxie/beego"
)

type OrmInterfaceController struct {
	beego.Controller
}

func (c *OrmInterfaceController) Get() {
	//o := orm.NewOrm()
	//orm.Debug = true

	// QueryRaws SetArgs接口
	//article := models.Article{}
	//r := o.Raw("SELECT * FROM article WHERE title=?", "红楼梦")
	//err := r.SetArgs("水浒传").QueryRows(&article)
	//if err != nil {
	//	utils.LogToFile("Error", fmt.Sprintf("查询数据错误，错误信息为: %s", err.Error()))
	//}
	//fmt.Println(article)
	//var articles []models.Article
	//articles = queryUtils.QueryArticlesByTitle("title", "ubernetes")
	//fmt.Println(articles)

	//r := o.Raw("SELECT * FROM article WHERE title=?", "红楼梦")
	//var params []orm.Params
	//
	//r.Values(&params)
	//var valueList []orm.ParamsList
	//r.ValuesList(&valueList)
	//fmt.Println(valueList)

	//var valueFloat orm.ParamsList
	//r.ValuesFlat(&valueFloat)
	//fmt.Println("\r")
	//fmt.Println("\r")
	//fmt.Println("\r")
	//fmt.Println("\r")
	//fmt.Printf("%v", valueFloat)

	//var rawToStruct models.RawToStruct
	//
	//r := o.Raw("SELECT id, name, value FROM raw_to_struct")
	//num, err := r.RowsToStruct(&rawToStruct, "name", "value")
	//if err != nil {
	//	fmt.Println("查询出错")
	//	fmt.Println(err)
	//}
	//fmt.Println(num)
	//fmt.Println(rawToStruct)

	// RawPrepare 支持CRU不支持D
	//var registerUsers []models.SaveRegisterUser
	//r := o.Raw("INSERT INTO save_register_users(username, email, password) VALUES (?, ?, ?)")
	//rp, err := r.Prepare()
	//if err != nil {
	//	utils.LogToFile("ERROR", fmt.Sprintf("执行批量操作失败，错误信息为: %v", err.Error()))
	//}
	//rp.Exec("王科", "leo65000@163.com", "Leo@960422")
	//rp.Exec("小太阳", "SmaillSun950228@163.com", "Sm@I1SvN9%0@28")
	//rp.Exec("宋威", "S0nyW1i@163.com", "(JYhcyt8662hv")
	//err = rp.Close()
	//if err != nil {
	//	fmt.Printf("执行报错操作失败，错误信息为： %s\n", err.Error())
	//	utils.LogToFile("Error", err.Error())
	//}

	c.TplName = "orm5.html"
}
