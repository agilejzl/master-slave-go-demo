package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	web "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"master-slave-go-demo/models"
	"strconv"
)

type BaseController struct {
	web.Controller
}

type ReturnMsg struct {
	Code int
	Msg  string
	Data interface{}
}

func (this *BaseController) SuccessJson(data interface{}) {
	res := ReturnMsg{
		200, "success", data,
	}
	this.Data["json"] = res
	this.ServeJSON() // 对json进行序列化输出
	this.StopRun()
}

func (this *BaseController) ErrorJson(code int, msg string, data interface{}) {
	res := ReturnMsg{
		code, msg, data,
	}
	this.Data["json"] = res
	this.ServeJSON() // 对json进行序列化输出
	this.StopRun()
}

func init() {
	fmt.Println("BaseController init")
	var FilterUser = func(ctx *context.Context) {
		authUserStr := ctx.Request.Header["Authorization"]
		if authUserStr != nil {
			userId, _ := strconv.Atoi(authUserStr[0])
			currentUser, _ := models.GetUsersById(userId)
			if currentUser == nil {
				logs.Error("Invalid Authorization:", authUserStr)
			}
		}
	}
	web.InsertFilter("/api/*", web.BeforeExec, FilterUser)
}
