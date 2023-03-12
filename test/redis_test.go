/**
* @Auth:ShenZ
* @Description:
* @CreateDate:2022/06/15 17:27:15
 */
package mq

import (
	"context"
	"fmt"
	"ginchat/utils"
	"testing"
	"time"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

// TestPublish 测试发布消息到redis
func TestPublish(t *testing.T) {
	msg := "当前时间: " + time.Now().Format("15:04:05")
	err := utils.Publish(ctx, utils.PublishKey, msg)
	if err != nil {

		fmt.Println(err)
	}
}
