package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/adapter/logs"
	web "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"master-slave-go-demo/helpers"
	"master-slave-go-demo/models"
	"reflect"
	"strconv"
)

type BaseController struct {
	web.Controller
}

type ReturnMsg struct {
	Code        int
	Msg         string
	CurrentUser interface{}
	Data        interface{}
}

func (this *BaseController) currentUser() models.UsersResp {
	currentUser, _ := (this.cookieUser("UsersResp")).(models.UsersResp)
	return currentUser
}

func (this *BaseController) currUserId() int64 {
	return this.currentUser().Id
}

func (this *BaseController) currUserIdStr() string {
	return strconv.FormatInt(this.currUserId(), 10)
}

// 将数组里的每个对象转化为指定结构体
func (this *BaseController) asJsonArray(data []interface{}, clazzName string, options map[string]string) []interface{} {
	var jsonResp []interface{}
	for _, model := range data {
		jsonResp = append(jsonResp, this.asJson(model, clazzName, options))
	}
	return jsonResp
}

// 将单个对象转化为指定结构体，data 为字符串或者 Model 对象
func (this *BaseController) asJson(data interface{}, clazzName string, options map[string]string) interface{} {
	var resByte string
	//fmt.Println(clazzName, " reflect.TypeOf(data).Name(): ", reflect.TypeOf(data).Kind())
	if "string" == reflect.TypeOf(data).Name() {
		resByte = data.(string)
	} else {
		marshalInfo, _ := json.Marshal(data)
		resByte = string(marshalInfo)
	}

	var productsRes models.ProductsResp
	var users models.Users
	var usersResp models.UsersResp
	var orderResp models.OrdersResp

	if clazzName == "ProductsResp" {
		json.Unmarshal([]byte(resByte), &productsRes)
		if options["NoOwner"] == "true" {
			usersResp.Id = productsRes.OwnerId
			productsRes.Owner = usersResp
		} else {
			owner, _ := models.GetUsersById(productsRes.OwnerId)
			productsRes.Owner = (this.asJson(owner, "UsersResp", options)).(models.UsersResp)
		}
		return productsRes
	} else if clazzName == "OrdersResp" {
		json.Unmarshal([]byte(resByte), &orderResp)
		product, _ := models.GetProductsById(orderResp.ProductId)
		orderResp.Product = (this.asJson(product, "ProductsResp", options)).(models.ProductsResp)
		return orderResp
	} else if clazzName == "UsersResp" {
		json.Unmarshal([]byte(resByte), &usersResp)
		return usersResp
	} else if clazzName == "Users" {
		json.Unmarshal([]byte(resByte), &users)
		return users
	}
	return nil
}

// 从 Cookie 获取当前用户信息，转化成 Users 或 UsersResp
func (this *BaseController) cookieUser(clazzName string) interface{} {
	cookie, _ := this.GetSecureCookie("", "userInfo")
	return this.asJson(cookie, clazzName, map[string]string{})
}

func (this *BaseController) SuccessJson(data interface{}) {
	res := ReturnMsg{
		200, "success", this.cookieUser("UsersResp"), data,
	}
	this.Data["json"] = res
	this.ServeJSON() // 对json进行序列化输出
	this.StopRun()
}

func (this *BaseController) ErrorJson(code int, msg string, data interface{}) {
	res := ReturnMsg{
		code, msg, models.UsersResp{}, data,
	}
	this.Data["json"] = res
	this.ServeJSON() // 对json进行序列化输出
	this.StopRun()
}

func setLogger(filename string) {
	if len(filename) == 0 {
		filename = "debug.log"
	}
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/`+filename+`","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
}

func init() {
	setLogger("")
	var FilterAuthUser = func(ctx *context.Context) {
		errorCode := 0
		authUserStr := ctx.Request.Header["Authorization"]
		userId := helpers.FakeData{}.RandUserId(authUserStr)

		if userId <= 0 {
			errorCode = 401
		} else {
			currentUser, err := helpers.FakeData{}.FakeNewUser(userId)
			if currentUser == nil {
				if err == nil {
					errorCode = 401
					logs.Warn("Failed Auth user:", authUserStr)
				} else {
					errorCode = 400
					logs.Error("Error Auth user:", authUserStr)
				}
			} else {
				// 将当前用户信息转为字符串存储到 Cookie
				logs.Debug("currentUser:", currentUser.Name)
				userInfo, _ := json.Marshal(currentUser)
				ctx.SetSecureCookie("", "userInfo", string(userInfo))
			}
		}

		if errorCode > 0 {
			ctx.Redirect(errorCode, "/login") // Todo render json
		}
	}
	// 验证接口的当前用户权限
	web.InsertFilter("/api/products/*", web.BeforeExec, FilterAuthUser)
	web.InsertFilter("/api/orders/*", web.BeforeExec, FilterAuthUser)
}
