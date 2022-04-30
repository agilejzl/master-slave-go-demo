package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id     int
	Name   string
	Credit string
}

func (u *User) TableName() string {
	return "users"
}

func init() {
	// need to register models in init
	orm.RegisterModel(new(User))
	// need to register default database
	//orm.RegisterDataBase("default", "mysql", "root:123@tcp(127.0.0.1:3306)/master_db?charset=utf8")
}
