package controllers

import (
	"net/http"
	"services"
	"entities"
	"github.com/gin-gonic/gin"
	"fmt"
	"utils"
)


// 创建店铺信息
func CreateShop(c *gin.Context) {
	var shopJson entities.Shop
	err := c.BindJSON(&shopJson)

	if err != nil {
		fmt.Printf("[controllers.GetShopById] err %s\n", err)
		c.JSON(200, gin.H{"code": 0, "data": "", "message": err})
		return
   }
   
   config, err := util.GetConfig();
   if err != nil {
		fmt.Printf("[controllers.GetShopById] err %s\n", err)
		c.JSON(200, gin.H{"code": 0, "data": "", "message": err})
		return
	}

   fmt.Printf("[controllers.CreateShop] --referer--- %s\n", c.Request.Header.Get("Referer"))
   headUrl := util.GetIpPort(c.Request.Header.Get("Referer"), config.UrlH5Boss)
   
   fmt.Printf("[controllers.CreateShop] --shop url inside qrcode--- %s\n", headUrl + config.UrlH5Boss + config.QrImgNameWeb)
   util.GenerateQrImg(headUrl + config.UrlH5Boss + config.QrImgNameWeb, config.ImageFolder, config.QrImgNameWeb)

   services.CreateShop(&shopJson)


   fmt.Printf("[controllers.CreateShop] ----- %s\n", shopJson)
   c.JSON(http.StatusOK, gin.H{"code": 1, "data": shopJson, "message": nil})
}


// 获取店铺信息
func GetShopById(c *gin.Context) {
	fmt.Printf("[controllers.GetShopById] param is:'%s'\n", c.Query("id"))
	shop := services.GetShop(c.Query("id"))
	fmt.Println(shop)
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": shop, "message": nil})
}


// 上传店铺图片
func UploadShopImage(c *gin.Context) {
	shopId := c.PostForm("shop_id")
	fmt.Printf("[controllers.UploadShopByImage] param is:'%s'\n", shopId)

	//得到上传的文件
	file, header, err := c.Request.FormFile("importfile") 
	if err != nil {
		fmt.Println("-----------------------------")
		fmt.Println(err)
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	services.UploadShopImage(file, shopId)

	fmt.Println(file, err, header.Filename) // useless
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": nil, "message": nil})
}

