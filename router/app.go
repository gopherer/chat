package router

import (
	"chat/docs"
	"chat/service"
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
	user := r.Group("/user")
	{
		user.POST("/getUserList", service.GetUserList)
		user.POST("/createUser", service.CreateUser)
		user.POST("/deleteUser", service.DeleteUser)
		user.POST("/updateUser", service.UpdateUser)
		user.POST("/findUserByNameAndPwd", service.FindUserByNameAndPwd)
		user.POST("/find", service.FindByID)
		//发送消息
		user.GET("/sendMsg", service.SendMsg)
		//发送消息
		user.GET("/sendUserMsg", service.SendUserMsg)
		//心跳续命 不合适  因为Node  所以前端发过来的消息再receProc里面处理
		// r.POST("/user/heartbeat", service.Heartbeat)
		user.POST("/redisMsg", service.RedisMsg)
	}

	contact := r.Group("contact")
	{
		//添加好友
		contact.POST("/addfriend", service.AddFriend)
		//创建群
		contact.POST("/createCommunity", service.CreateCommunity)
		//群列表
		contact.POST("/loadcommunity", service.LoadCommunity)
		contact.POST("/joinGroup", service.JoinGroups)
	}

	//上传文件
	r.POST("/attach/upload", service.Upload)

	return r
}
