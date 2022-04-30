package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"master-slave-go-demo/models"
)

func init() {
	logs.SetLogger(logs.AdapterConsole)
	logs.SetLevel(logs.LevelDebug)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123@tcp(127.0.0.1:3306)/master_db?charset=utf8")
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	Orm := orm.NewOrm()
	var user models.User
	Orm.QueryTable(models.User{}).Filter("id", 1).One(&user)
	c.Data["user"] = user

	var product models.Product
	err := Orm.QueryTable(models.Product{}).Filter("id", 1).One(&product)
	if err == orm.ErrNoRows {
		fmt.Printf("查询不到数据")
	} else {
		logs.Debug("product", product)
	}
	c.Data["product"] = product

	c.Data["Website"] = "https://github.com/agilejzl"
	c.Data["Email"] = "teddy.jiang@qq.com"
	c.TplName = "index.tpl"
}
