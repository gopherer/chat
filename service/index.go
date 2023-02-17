package service

import "github.com/gin-gonic/gin"

//GetIndex
// @Tags GetIndex
// @Success 200 {string} pong
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
