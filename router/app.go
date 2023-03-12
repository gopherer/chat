/**
* @Auth:ShenZ
* @Description:
* @CreateDate:2022/06/14 11:57:55
 */
package router

import (
	"ginchat/docs"
	"ginchat/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {

	r := gin.Default()
	//swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//静态资源
	r.Static("/asset", "asset/")
	r.StaticFile("/favicon.ico", "asset/images/favicon.ico")
	//	r.StaticFS()
	r.LoadHTMLGlob("views/**/*")

	//首页
	r.GET("/", service.GetIndex)
	r.GET("/index", service.GetIndex)
	r.GET("/toRegister", service.ToRegister)
	r.GET("/toChat", service.ToChat)
	r.GET("/chat", service.Chat)
	r.POST("/searchFriends", service.SearchFriends)

	//用户模块
	r.POST("/user/getUserList", service.GetUserList)
	r.POST("/user/createUser", service.CreateUser)
	r.POST("/user/deleteUser", service.DeleteUser)
	r.POST("/user/updateUser", service.UpdateUser)
	r.POST("/user/findUserByNameAndPwd", service.FindUserByNameAndPwd)
	r.POST("/user/find", service.FindByID)
	//发送消息
	r.GET("/user/sendMsg", service.SendMsg)
	//发送消息
	r.GET("/user/sendUserMsg", service.SendUserMsg)
	//添加好友
	r.POST("/contact/addfriend", service.AddFriend)
	//上传文件
	r.POST("/attach/upload", service.Upload)
	//创建群
	r.POST("/contact/createCommunity", service.CreateCommunity)
	//群列表
	r.POST("/contact/loadcommunity", service.LoadCommunity)
	r.POST("/contact/joinGroup", service.JoinGroups)
	//心跳续命 不合适  因为Node  所以前端发过来的消息再receProc里面处理
	// r.POST("/user/heartbeat", service.Heartbeat)
	r.POST("/user/redisMsg", service.RedisMsg)
	return r
}
