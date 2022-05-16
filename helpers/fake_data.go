package helpers

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/bxcodec/faker/v3"
	"master-slave-go-demo/models"
	"math/rand"
	"strconv"
	"strings"
)

type FakeData struct{}

func (fd FakeData) genRandInt(minNum int, maxNum int) int {
	return minNum + rand.Intn(maxNum-minNum+1)
}

func (fd FakeData) genRandInt64(minNum int64, maxNum int64) int64 {
	return minNum + rand.Int63n(maxNum-minNum+1)
}

func (fd FakeData) genRandFloat64(min, max float64) float64 {
	res := make([]float64, 1)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res[0]
}

func (fd FakeData) FakeUpdateOrderStatus() *models.Orders {
	randOrder := models.GetRandomOrder()
	status := fd.genRandInt(1, 2)
	randOrder = models.UpdatePayStatus(randOrder.Id, status)
	return randOrder
}

func (fd FakeData) FakeNewOrder(userId int64) (*models.Orders, error) {
	randProduct := models.GetRandomProduct()
	PdAmount := fd.genRandInt(1, 10)
	order := models.Orders{UserId: userId, ProductId: randProduct.Id,
		Status: 0, PdAmount: PdAmount, TotalPrice: float64(PdAmount) * randProduct.PdPrice}
	id, err := models.AddOrders(&order)
	if err == nil {
		fmt.Println("NewOrder Id", id, ":", order)
		return &order, err
	} else {
		logs.Error("Error NewOrder:", err)
		return nil, err
	}
}

func (fd FakeData) FakeNewProduct(userId int64) (*models.Products, error) {
	product := models.Products{OwnerId: userId}
	product.Name = "No." + " " + faker.Phonenumber()
	product.StockAmount = fd.genRandInt(5400, 5600)
	product.PdPrice = fd.genRandFloat64(0.0, 9.99)

	id, err := models.AddProducts(&product)
	if err == nil {
		fmt.Println("NewProduct Id", id, ":", product)
		return &product, err
	} else {
		logs.Error("Error NewProduct:", err)
		return nil, err
	}
}

// FakeNewUser 根据用户ID，查找或创建用户
func (fd FakeData) FakeNewUser(userId int64) (*models.Users, error) {
	existUser, err := models.GetUsersById(userId)
	if existUser != nil {
		// logs.Debug("existUser:", existUser)
		return existUser, err
	} else {
		userIdStr := "No." + strconv.FormatInt(userId, 10) + " " + faker.FirstName()
		userModel := models.Users{Id: userId, Name: userIdStr}
		_, err := models.AddUsers(&userModel)
		if err == nil {
			return &userModel, err
		} else {
			logs.Error("Error NewUser:", err)
			return nil, err
		}
	}
}

func (fd FakeData) RandUserId(authUserStr []string) int64 {
	var minId int64
	if authUserStr == nil {
		return minId
	} else {
		rangeIds := strings.Split(authUserStr[0], "-")
		minId, _ = strconv.ParseInt(rangeIds[0], 10, 0)
		if len(rangeIds) >= 2 {
			maxId, _ := strconv.ParseInt(rangeIds[1], 10, 0)
			return fd.genRandInt64(minId, maxId)
		} else {
			return minId
		}
	}
}
