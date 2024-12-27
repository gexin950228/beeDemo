package OrmTest

import (
	"beeDemo/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type TestOrmController struct {
	beego.Controller
}

type User struct {
	Id            int       `orm:"pk;auto"`
	Username      string    `orm:"size(500)"`
	Address       string    `orm:"size(500)"`
	LastLoginTime time.Time `orm:"auto_now_add;type(datetime)"`
	RedirectUri   string    `orm:"size(500)"`
	Password      string    `orm:"size(500)"`
}

func (u *User) TableName() string {
	return "sys_user"
}

func init() {
	mysqlConn := utils.LoadMysqlConfig()
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dst := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", mysqlConn.User, mysqlConn.Password, mysqlConn.Host, mysqlConn.Port, mysqlConn.Database)
	orm.RegisterDataBase("default", "mysql", dst)
	orm.RegisterModel(new(User))
}

func (t *TestOrmController) Get() {
	o := orm.NewOrm()

	//插入单条数据
	//u := User{Username: "葛新", Address: "北京市海淀区", LastLoginTime: time.Now(), RedirectUri: "/parameter", Password: "Ch1nZh1nhg$a"}
	//id, err := o.Insert(&u)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	utils.LogToFile("Error", fmt.Sprintf("插入数据错误， error: %s", err.Error()))
	//} else {
	//	fmt.Println(id)
	// }
	//mysqlConn := utils.LoadMysqlConfig()
	//fmt.Println(mysqlConn)

	// 插入多条数据
	//users := []User{
	//	{Username: "高林飞", Address: "北京市", LastLoginTime: time.Now(), RedirectUri: "/", Password: "123456"},
	//	{Username: "刘继雄", Address: "北京市", LastLoginTime: time.Now(), RedirectUri: "/", Password: "123456"},
	//	{Username: "张旭涛", Address: "北京市", LastLoginTime: time.Now(), RedirectUri: "/", Password: "123456"},
	//}
	//multi, err := o.InsertMulti(3, &users)
	//if err != nil {
	//	return
	//}
	//fmt.Printf("Multi: %v\n", multi)

	//读取数据
	//user := User{Id: 3}
	//o.Read(&user)
	//fmt.Println(user)
	//
	//user1 := User{Username: "刘继雄", Address: "北京市昌平区未来科学神华研究院"}
	//o.Read(&user1, "Username", "Address")
	//fmt.Printf("user1: %+v \n", user1)

	// ReadOrCreate
	user := User{Username: "乔建康", Address: "北京市东城区安定门国家能源集团C座", Password: "Q!@oJianK24", RedirectUri: "/health", LastLoginTime: time.Now()}
	create, i, err := o.ReadOrCreate(&user, "Username", "Address", "Password")
	if err != nil {
		return
	}
	fmt.Println(create, i)
	t.TplName = "testOrm/testorm1.html"
}