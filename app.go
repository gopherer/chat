package router

import (
	"chat/docs"
	"chat/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	//Swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//首页
	r.GET("/index", service.GetIndex)
	//用户模块
	r.GET("/user/GetUserList", service.GetUserList)
	r.GET("/user/CreateUser", service.CreateUser)
	r.GET("/user/DeleteUser", service.DeleteUser)
	r.POST("/user/UpdateUser", service.UpdateUser)
	r.POST("/user/FindUserByNameAndPwd", service.FindUserByNameAndPwd)

	//发送消息
	r.GET("/user/SendMsg", service.SendMsg)

	r.GET("/user/SendUserMsg", service.SendUserMsg)
	return r
}
