package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	web "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"master-slave-go-demo/helper"
	"master-slave-go-demo/models"
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

func (this *BaseController) cookieUser(clazzName string) interface{} {
	cookie, _ := this.GetSecureCookie("", "userInfo")
	currentUser := models.Users{}
	currentUserResp := models.UsersResp{}

	if clazzName == "Users" {
		json.Unmarshal([]byte(cookie), &currentUser)
		return currentUser
	} else {
		json.Unmarshal([]byte(cookie), &currentUserResp)
		return currentUserResp
	}
	return nil
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

func init() {
	fmt.Println("BaseController init")
	var FilterAuthUser = func(ctx *context.Context) {
		errorCode := 0
		authUserStr := ctx.Request.Header["Authorization"]
		userId := helper.FakeData{}.RandUserId(authUserStr)

		if userId <= 0 {
			errorCode = 401
		} else {
			currentUser, err := helper.FakeData{}.FakeNewUser(userId)
			if currentUser == nil {
				if err == nil {
					errorCode = 401
					logs.Warn("Failed Auth user:", authUserStr)
				} else {
					errorCode = 400
					logs.Error("Error Auth user:", authUserStr)
				}
			} else {
				userInfo, _ := json.Marshal(currentUser)
				ctx.SetSecureCookie("", "userInfo", string(userInfo))
			}
		}

		if errorCode > 0 {
			ctx.Redirect(errorCode, "/login") // Todo render json
		}
	}
	web.InsertFilter("/api/products/*", web.BeforeExec, FilterAuthUser)
	web.InsertFilter("/api/orders/*", web.BeforeExec, FilterAuthUser)
}
