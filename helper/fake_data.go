package helper

import (
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/bxcodec/faker/v3"
	"master-slave-go-demo/models"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type FakeData struct{}

func (fd FakeData) FakeUser(userId int) (interface{}, error) {
	existUser, err := models.GetUsersById(userId)
	if existUser != nil {
		// logs.Debug("existUser:", existUser)
		return existUser, err
	} else {
		userIdStr := "No." + strconv.Itoa(userId) + " " + faker.FirstName()
		userModel := models.Users{Id: userId, Name: userIdStr, CreatedAt: time.Now(), UpdatedAt: time.Now()}
		_, err := models.AddUsers(&userModel)
		if err == nil {
			return userModel, err
		} else {
			logs.Error("Error AddUser:", err)
			return nil, err
		}
	}
}

func (fd FakeData) RandUserId(authUserStr []string) int {
	minId := 0
	if authUserStr == nil {
		return minId
	} else {
		rangeIds := strings.Split(authUserStr[0], "-")
		minId, _ = strconv.Atoi(rangeIds[0])
		if len(rangeIds) >= 2 {
			maxId, _ := strconv.Atoi(rangeIds[1])
			return minId + rand.Intn(maxId-minId)
		} else {
			return minId
		}
	}
}
