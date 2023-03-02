package main

import (
	"chat/router"
	"chat/utils"
)

func main() {
	r := router.Router()
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	r.Run(":80")
}
