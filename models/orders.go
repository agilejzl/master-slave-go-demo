package models

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Orders struct {
	Id         int64     `orm:"column(id);auto"`
	UserId     int64     `orm:"column(user_id)"`
	ProductId  int64     `orm:"column(product_id)"`
	PdAmount   int       `orm:"column(pd_amount)"`
	TotalPrice float64   `orm:"column(total_price);digits(10);decimals(2)"`
	Status     int       `orm:"column(status);null"`
	CreatedAt  time.Time `orm:"column(created_at);type(datetime);auto_now_add"`
	UpdatedAt  time.Time `orm:"column(updated_at);type(datetime);auto_now"`
}

type OrdersResp struct {
	Id         int64
	UserId     int64
	ProductId  int64
	PdAmount   int
	PdPrice    float64
	TotalPrice float64
	Status     int8
	UpdatedAt  time.Time `orm:"column(updated_at);type(datetime)"`
	Product    ProductsResp
}

func (t *Orders) TableName() string {
	return "orders"
}

func init() {
	orm.RegisterModel(new(Orders))
}

func (t *Orders) Statuses() map[string]int {
	statuses := make(map[string]int)
	statuses["unpaid"] = 0
	statuses["closed"] = 1
	statuses["successful"] = 2
	return statuses
}

// AddOrders insert a new Orders into database and returns
// last inserted Id on success.
func AddOrders(m *Orders) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// retrieves Orders by Id. Returns error if Id doesn't exist
func GetOrdersById(id int64) (v *Orders, err error) {
	o := orm.NewOrm()
	v = &Orders{Id: id}
	if err = o.Read(v); err == nil {
		//user := &Users{}
		//o.QueryTable("users").Filter("Id", id).RelatedSel().One(user)
		//fmt.Println("OrdersById", id, user)
		return v, nil
	}
	return nil, err
}

func GetRandomOrder() (v *Orders) {
	o := orm.NewOrm()
	o.Raw("SELECT * FROM orders ORDER BY RAND() LIMIT 1;").QueryRow(&v)
	return v
}

func UpdatePayStatus(id int64, status int) (order *Orders) {
	order, _ = GetOrdersById(id)
	order.Status = status
	if order.Status == order.Statuses()["closed"] {
		UpdateOrdersById(order)
	} else if order.Status == order.Statuses()["successful"] {
		o := orm.NewOrm()
		o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
			product, _ := GetProductsById(order.ProductId)
			UpdateOrdersById(order)
			UpdateRawBySQL("UPDATE users SET credit=credit+? WHERE id=?", order.TotalPrice, product.OwnerId)
			UpdateRawBySQL("UPDATE users SET credit=credit-? WHERE id=?", order.TotalPrice, order.UserId)
			// Todo ensure credit >= 0
			return nil
		})
	}
	return
}

// GetAllOrders retrieves all Orders matches certain condition. Returns empty list if
// no records exist
func GetAllOrders(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Orders))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Orders
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateOrders updates Orders by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrdersById(m *Orders) (err error) {
	o := orm.NewOrm()
	v := Orders{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOrders deletes Orders by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrders(id int64) (err error) {
	o := orm.NewOrm()
	v := Orders{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Orders{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
