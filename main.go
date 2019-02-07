package main

import (
	"animation/controller"
	"animation/entity/link"
	"animation/luff/middleware/jwt"
	"github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//读取配置文件
	cfg, _ := goconfig.LoadConfigFile("config/config.ini")
	cfd, err := cfg.GetSection(goconfig.DEFAULT_SECTION)
	if err != nil {
		log.Fatalf("无法加载配置文件: %s", err)
	}

	//配置xorm
	dns := string(cfd["xorm_dns"])
	err = link.Init(dns)
	if err != nil {
		panic(err)
	}

	//配置路由
	router := gin.Default()

	router.POST("/login", controller.Login)
	router.POST("/register", controller.Register)

	admin := router.Group("/admin")
	admin.Use(jwt.JWTAuth())
	{
		/**
		lens: 镜头操作
		temp: 模版操作
		ad: 广告操作
		cost: 制作价格费用操作
		vip: 会员价格费用操作
		login: 登录操作
		 */
		admin.POST("/lens", controller.AdminAddLens)
		admin.PUT("/lens", controller.AdminUpdateLens)
		admin.DELETE("lens", controller.AdminDeleteLens)
		admin.GET("/lens", controller.AdminListLens)

		admin.POST("/temp", controller.AdminAddTemp)
		admin.PUT("/temp", controller.AdminUpdateTemp)
		admin.DELETE("temp", controller.AdminDeleteTemp)
		admin.GET("/temp", controller.AdminListTemp)

		admin.POST("/ad", controller.AdminAddAd)
		admin.PUT("/ad", controller.AdminUpdateAd)
		admin.DELETE("ad", controller.AdminDeleteAd)
		admin.GET("/ad", controller.AdminListAd)

		admin.POST("/cost", controller.AdminAddCost)
		admin.PUT("/cost", controller.AdminUpdateCost)
		admin.DELETE("/cost", controller.AdminDeleteCost)
		admin.GET("/cost", controller.AdminListCost)

		admin.POST("/vip", controller.AdminAddVip)
		admin.PUT("/vip", controller.AdminUpdateVip)
		admin.DELETE("/vip", controller.AdminDeleteVip)
		admin.GET("/vip", controller.AdminListVip)

		//admin.POST("/login", controller.Login)
	}


	router.Run()
}
