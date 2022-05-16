package controllers

import (
	"master-slave-go-demo/helpers"
	"master-slave-go-demo/models"
	"strconv"
)

// OrdersController operations for Orders
type OrdersController struct {
	BaseController
}

// URLMapping ...
func (c *OrdersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
}

// Post ...
// @Title Post
// @Description create Orders
// @Param	body		body 	models.Orders	true		"body for Orders content"
// @Success 201 {int} models.Orders
// @Failure 403 body is empty
// @router / [post]
func (c *OrdersController) Post() {
	data, err := helpers.FakeData{}.FakeNewOrder(c.currUserId())
	if err == nil {
		c.SuccessJson(data)
	} else {
		c.ErrorJson(400, err.Error(), data)
	}
}

// Put ...
// @router /:id [put]
func (c *OrdersController) Put() {
	//idStr := c.Ctx.Input.Param(":id")
	//id, _ := strconv.ParseInt(idStr, 10, 0)
	order := helpers.FakeData{}.FakeUpdateOrderStatus()
	data := c.asJson(order, "OrdersResp", map[string]string{})
	c.SuccessJson(data)
}

// GetOne ...
// @Title Get One
// @Description get Orders by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Orders
// @Failure 403 :id is empty
// @router /:id [get]
func (c *OrdersController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 0)
	order, err := models.GetOrdersById(id)
	data := c.asJson(order, "OrdersResp", map[string]string{})

	if err != nil {
		c.ErrorJson(400, err.Error(), data)
	} else {
		c.SuccessJson(data)
	}
}

// GetAll ...
// @Title Get All
// @Description get Orders
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Orders
// @Failure 403
// @router / [get]
func (c *OrdersController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	query["UserId"] = c.currUserIdStr()
	orders, err := models.GetAllOrders(query, fields, sortby, order, offset, limit)
	query["NoOwner"] = "true"
	data := c.asJsonArray(orders, "OrdersResp", query)

	if err != nil {
		c.ErrorJson(400, err.Error(), data)
	} else {
		c.SuccessJson(data)
	}
}
