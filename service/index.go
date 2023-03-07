package service

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

//GetIndex
// @Tags GetIndex
// @Success 200 {string} pong
// @Router /index [get]
func GetIndex(c *gin.Context) {
	inx, err := template.ParseFiles("views/user/index.html")
	if err != nil {
		panic(err)
	}
	inx.Execute(c.Writer, nil)
	//c.JSON(200, gin.H{
	//	"message": "pong",
	//})
}

func ToRegister(c *gin.Context) {
	inx, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	inx.Execute(c.Writer, "index")
	//c.JSON(200, gin.H{
	//	"message": "pong",
	//})
}
