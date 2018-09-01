package controllers

import (
	"utils"
	"net/http"
	"services"
	"entities"
	"github.com/gin-gonic/gin"
	"fmt"
)


// 下单
func CreateOrder(c *gin.Context) {
	defer util.PanicHttpHandler(c)

	var orderJ entities.OrderJson
	err := c.BindJSON(&orderJ)
	if(err != nil) {
		panic(err)
	}
	
	//fmt.Printf("[controllers.CreateOrder] ----- %s\n", dishJson)
	services.CreateOrder(orderJ)
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": "", "message": nil})
}


func GetAllOrders(c *gin.Context) {
	fmt.Printf("[controllers.GetAllOrders] \n")
	orders := services.GetAllOrders()
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": orders, "message": nil})
}


func GetOrderBySeatNo(c *gin.Context) {
	fmt.Printf("[controllers.GetOrderBySeatNo] \n")
	orders := services.GetOrderBySeatNo(c.Query("seat_no"))
	fmt.Println("------>>>>>>>>>>>------------")
	fmt.Println(orders)
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": orders, "message": nil})
}