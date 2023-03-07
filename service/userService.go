package service

import (
	"chat/models"
	"chat/utils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//GetUserList
//@Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} data
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"code":    0, //	0 成功   -1 失败
		"message": "用户列表",
		"data":    data,
	})

}

//CreateUser
//@Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param rePassword query string false "确认密码"
// @Success 200 {string} data
// @Router /user/CreateUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	rePassword := c.Request.FormValue("rePassword")

	salt := fmt.Sprintf("%06d", rand.Int31())

	if user.Name == "" || user.PassWord == "" {

		c.JSON(http.StatusOK, gin.H{
			"code":    -1, //	0 成功   -1 失败
			"message": "用户名或密码不能为空",
			"data":    user,
		})
		return
	}

	data := models.FindUserByName(user.Name)
	if data.Name != "" {

		c.JSON(http.StatusOK, gin.H{
			"code":    -1, //	0 成功   -1 失败
			"message": "用户已被注册",
			"data":    user,
		})
		return
	}
	if password != rePassword {

		c.JSON(http.StatusOK, gin.H{
			"code":    -1, //	0 成功   -1 失败
			"message": "两次密码不一致",
			"data":    user,
		})
		return
	}
	//user.PassWord = password
	user.PassWord = utils.MakePassword(password, salt)
	user.Salt = salt
	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code":    0, //	0 成功   -1 失败
		"message": "新增用户成功",
		"data":    user,
	})
}

//DeleteUser
//@Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} data
// @Router /user/DeleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)

	c.JSON(http.StatusOK, gin.H{
		"code":    0, //	0 成功   -1 失败
		"message": "用户删除成功",
		"data":    user,
	})
}

//UpdateUser
//@Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} data
// @Router /user/UpdateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1, //	0 成功   -1 失败
			"message": "修改参数不匹配",
			"data":    user,
		})
		return
	} else {
		models.UpdateUser(user)

		c.JSON(http.StatusOK, gin.H{
			"code":    0, //	0 成功   -1 失败
			"message": "用户修改成功",
			"data":    user,
		})
	}
}

//FindUserByNameAndPwd
//@Summary 用户登录
// @Tags 用户模块
// @param name formData string false "name"
// @param password formData string false "password"
// @Success 200 {string} data
// @Router /user/FindUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	data := models.UserBasic{}
	name := c.PostForm("name")
	passWord := c.PostForm("password")
	fmt.Println(name, passWord)
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1, //	0 成功   -1 失败
			"message": "该用户不存在",
			"data":    data,
		})
		return
	}
	flag := utils.ValidPassword(passWord, user.Salt, user.PassWord)
	if !flag {
		c.JSON(http.StatusOK, gin.H{
			"message": "密码不正确",
		})
		return
	}
	pwd := utils.MakePassword(passWord, user.Salt)
	data = models.FindUserByNameAndPwd(name, pwd)

	c.JSON(http.StatusOK, gin.H{
		"code":    0, //	0 成功   -1 失败
		"message": "登陆成功",
		"data":    data,
	})
}

//防止跨域站点伪造请求
var upGrad = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upGrad.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	MsgHandle(ws, c)
}

func MsgHandle(ws *websocket.Conn, c *gin.Context) {
	msg, err := utils.Subscribe(c, utils.PublishKey)
	if err != nil {
		fmt.Println(err)
	}
	tm := time.Now().Format("2006-06-02 15-01-05")
	m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
	err = ws.WriteMessage(1, []byte(m))
	if err != nil {
		fmt.Println(err)
	}
}

func SendUserMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}
