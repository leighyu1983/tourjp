package controllers

import (
	"utils"
	"net/http"
	"services"
	"entities"
	"github.com/gin-gonic/gin"
	"fmt"
)


// 创建单品信息
func CreateDishJP(c *gin.Context) {
	defer util.PanicHttpHandler(c)

	var dishJson entities.DishJson
	err := c.BindJSON(&dishJson)
	if(err != nil) {
		panic(err)
	}
	
	//fmt.Printf("[controllers.CreateDishJP] ----- %s\n", dishJson)
	id := services.CreateDishJP(&dishJson)
	idJson := fmt.Sprintf("{\"id\":\"%s\"}", id)
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": idJson, "message": nil})
}


// 更新单品信息(中文)
func UpdateDishCN(c *gin.Context) {
	defer util.PanicHttpHandler(c)

	var dishJson entities.DishJson
	err := c.BindJSON(&dishJson)

	if(err != nil) {
		panic(err)
	}

	fmt.Printf("[controllers.UpdateDishCN] ----- %s\n", dishJson)
    services.UpdateDishCN(&dishJson)
    c.JSON(http.StatusOK, gin.H{"code": 1, "data": dishJson, "message": nil})
}

// 获取单品信息
func GetAllDishes(c *gin.Context) {
	defer util.PanicHttpHandler(c)
	fmt.Printf("[controllers.GetAllDishes] ...\n")
	dishes := services.GetAllDishes()
	fmt.Println(dishes)
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": dishes, "message": nil})
}


// 上传单品图片
func UploadDishImage(c *gin.Context) {
	defer util.PanicHttpHandler(c)

	dishId := c.PostForm("dish_id")

	//得到上传的文件
	file, header, err := c.Request.FormFile("importfile") 
	if(err != nil) {
		panic(err)
	}
	
	fmt.Printf("[controllers.UploadDishImage] '%s'--->'%s'\n", dishId, header.Filename)
	services.UploadDishImage(file, header.Filename, dishId)

	fmt.Println(file, err, header.Filename) // useless
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": nil, "message": nil})
}