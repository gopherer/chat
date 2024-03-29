package main

import (
	"chat/models"
	"chat/router" //  router "ginchat/router"
	"chat/utils"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/spf13/viper"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	InitTimer()
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	r := router.Router()
	err := r.Run(viper.GetString("port.server"))
	if err != nil {
		fmt.Println("r.Run err", err)
		return
	} // listen and serve on 0.0.0.0:80
}

// InitTimer 初始化定时器
func InitTimer() {
	utils.Timer(time.Duration(viper.GetInt("timeout.DelayHeartbeat"))*time.Second, time.Duration(viper.GetInt("timeout.HeartbeatHz"))*time.Second, models.CleanConnection, "")
}
