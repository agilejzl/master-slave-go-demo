package controllers

import (
	"master-slave-go-demo/helper"
	"master-slave-go-demo/models"
	"strconv"
)

// ProductsController operations for Products
type ProductsController struct {
	BaseController
}

// URLMapping ...
func (c *ProductsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
}

// Post ...
// @Title Post
// @Description create Products
// @Param	body		body 	models.Products	true		"body for Products content"
// @Success 201 {int} models.Products
// @Failure 403 body is empty
// @router / [post]
func (c *ProductsController) Post() {
	// var product models.Products
	// json.Unmarshal(c.Ctx.Input.RequestBody, &product)
	data, err := helper.FakeData{}.FakeNewProduct(c.currUserId())
	if err == nil {
		c.SuccessJson(data)
	} else {
		c.ErrorJson(400, err.Error(), data)
	}
}

// GetOne ...
// @Title Get One
// @Description get Products by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Products
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ProductsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 0)
	data, err := models.GetProductsById(id)
	if err != nil {
		c.ErrorJson(400, err.Error(), data)
	} else {
		c.SuccessJson(data)
	}
}

// GetAll ...
// @Title Get All
// @Description get Products
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Products
// @Failure 403
// @router / [get]
func (c *ProductsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	query["OwnerId"] = c.currUserIdStr()
	data, err := models.GetAllProducts(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ErrorJson(400, err.Error(), data)
	} else {
		c.SuccessJson(data)
	}
}
