package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Id          int
	OwnerId     int
	Name        string
	StockAmount int
	PdPrice     float32
}

func (u *Product) TableName() string {
	return "products"
}

func init() {
	// need to register models in init
	orm.RegisterModel(new(Product))
	// need to register default database
	//orm.RegisterDataBase("default", "mysql", "root:123@tcp(127.0.0.1:3306)/master_db?charset=utf8")
}
