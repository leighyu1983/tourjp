package main

import (
	"controllers"
	"fmt"
	"utils"
	"strings"
	"github.com/gin-gonic/gin"
	"net/http"
 )

func main() {
	//services.Case_1();
	runListener();
	fmt.Print("finished");
}

func runListener() {
	defer util.PanicHandler()
	router := gin.Default()
	router.Use(Cors())

	setWebRouters(router)
	setH5Routers(router)

	router.Run(":805")
}

func setWebRouters(c *gin.Engine) {
	g := c.Group("/web") 
	{
		//添加店铺
		g.POST("/shop/create", controllers.CreateShop)
		//上传店铺图片
		g.POST("/shop/image/upload", controllers.UploadShopImage)
		//获取店铺信息
		g.GET("/shop/get", controllers.GetShopById)	


		//添加单品(日语信息)
		g.POST("/dish/add_jp", controllers.CreateDishJP)
		//更新单品(中文信息)
		g.POST("/dish/update_cn", controllers.UpdateDishCN)
		// 获取全部单品
		g.GET("/dish/all", controllers.GetAllDishes)
		//上传单品图片
		g.POST("/dish/image/upload", controllers.UploadDishImage)


		//添加套餐(日语信息)
		g.POST("/set/add_jp", controllers.CreateSetJP)
		//更新套餐(中文信息)
		g.POST("/set/update_cn", controllers.UpdateSetCN)
		//获取全部套餐
		g.GET("/set/all", controllers.GetAllSets)
		//上传套餐图片
		g.POST("/set/image/upload", controllers.UploadSetImage)
	}
}

func setH5Routers(c *gin.Engine) {
	g := c.Group("/h5") 
	{
		//顾客下单
		g.POST("/customer/order/submit", controllers.CreateOrder)
		//获取全部
		g.GET("/employee/order/all", controllers.GetAllOrders)
		//员工确认下单 (根据桌号获取订单)
		g.GET("/employee/order/get", controllers.GetOrderBySeatNo)
	}
}


func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        method := c.Request.Method      //请求方法
        origin := c.Request.Header.Get("Origin")        //请求头部
        var headerKeys []string                             // 声明请求头keys
        for k, _ := range c.Request.Header {
            headerKeys = append(headerKeys, k)
        }
        headerStr := strings.Join(headerKeys, ", ")
        if headerStr != "" {
            headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
        } else {
            headerStr = "access-control-allow-origin, access-control-allow-headers"
        }
        if origin != "" {
            c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
            c.Header("Access-Control-Allow-Origin", "*")        // 这是允许访问所有域
            c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")      //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
            //  header的类型
            c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
            //              允许跨域设置                                                                                                      可以返回其他子段
            c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")      // 跨域关键设置 让浏览器可以解析
            c.Header("Access-Control-Max-Age", "172800")        // 缓存请求信息 单位为秒
            c.Header("Access-Control-Allow-Credentials", "false")       //  跨域请求是否需要带cookie信息 默认设置为true
            c.Set("content-type", "application/json")       // 设置返回格式是json
        }

		//放行所有OPTIONS方法
		fmt.Println("~~~~~~~~~~~~~~~~~~" + method)

        if method == "OPTIONS" {
            c.JSON(http.StatusOK, "Options Request!")
        }
        // 处理请求
        c.Next()        //  处理请求
    }
}

