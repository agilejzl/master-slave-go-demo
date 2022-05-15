package models

import (
	"github.com/beego/beego/v2/client/orm"
)

func UpdateRawBySQL(updateSQL string, args ...interface{}) int64 {
	res, err := orm.NewOrm().Raw(updateSQL, args).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		if num > 0 {
			// fmt.Println("Updated", num, "rows,", updateSQL, args)
			return num
		}
	}
	return 0
}
