package controllers

import (
	"net/http"
	"services"
	"entities"
	"github.com/gin-gonic/gin"
	"fmt"
	"utils"
)


// 创建套餐(日语信息)
func CreateSetJP(c *gin.Context) {
	var setJson entities.SetJsonIn
	err := c.BindJSON(&setJson)

	if err != nil {
		fmt.Printf("[controllers.CreateSet] err %s\n", err)
		c.JSON(200, gin.H{"code": 0, "data": "", "message": err})
		return
   }

   services.CreateSetJP(&setJson)
   fmt.Printf("[controllers.CreateSet] ----- %s\n", setJson)
   c.JSON(http.StatusOK, gin.H{"code": 1, "data": setJson, "message": nil})
}


// 获取全部套餐
func GetAllSets(c *gin.Context) {
	fmt.Printf("[controllers.GetAllSets] \n")
	sets := services.GetAllSets()
	fmt.Println(sets)
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": sets, "message": nil})
}


// 上传套餐图片
func UploadSetImage(c *gin.Context) {
	defer util.PanicHttpHandler(c)

	setId := c.PostForm("set_id")

	//得到上传的文件
	file, header, err := c.Request.FormFile("importfile") 
	if(err != nil) {
		panic(err)
	}
	
	fmt.Printf("[controllers.UploadSetImage] '%s'--->'%s'\n", setId, header.Filename)
	services.UploadSetImage(file, header.Filename, setId)

	fmt.Println(file, err, header.Filename) // useless
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": nil, "message": nil})
}

// 更新套餐信息(中文)
func UpdateSetCN(c *gin.Context) {
	defer util.PanicHttpHandler(c)

	var setJson entities.SetJsonIn
	err := c.BindJSON(&setJson)

	if(err != nil) {
		panic(err)
	}

	fmt.Printf("[controllers.UpdateSetCN] ----- %s\n", setJson)
    services.UpdateSetCN(&setJson)
    c.JSON(http.StatusOK, gin.H{"code": 1, "data": setJson, "message": nil})
}