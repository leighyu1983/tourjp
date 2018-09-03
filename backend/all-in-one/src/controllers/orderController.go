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
	
	config, err := util.GetConfig();
	if err != nil {
		 fmt.Printf("[controllers.GetShopById] err %s\n", err)
		 c.JSON(200, gin.H{"code": 0, "data": "", "message": err})
		 return
	 }

	//fmt.Printf("[controllers.CreateOrder] ----- %s\n", dishJson)
	services.CreateOrder(orderJ)

	fmt.Printf("[controllers.CreateOrder] --referer--- %s\n", c.Request.Header.Get("Referer"))
	headUrl := util.GetIpPort(c.Request.Header.Get("Referer"), config.UrlH5Customer)
	
	fmt.Printf("[controllers.CreateOrder] --customer url inside qrcode--- %s\n", headUrl + config.UrlH5Customer)
	util.GenerateQrImg(headUrl + config.UrlH5Employee + "?seat_no=" + orderJ.SeatNo, config.ImageFolder, config.QrImgNameEmployee)

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