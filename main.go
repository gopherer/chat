package main

import (
	"chat/router"
	"chat/utils"
)

func main() {
	r := router.Router()
	utils.InitConfig()
	utils.InitMySQL()
	r.Run(":80")
}
