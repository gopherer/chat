/**
* @Auth:ShenZ
* @Description:
* @CreateDate:2022/06/15 12:57:55
 */
package service

import (
	"fmt"
	"ginchat/utils"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Upload(c *gin.Context) {
	UploadLocal(c)
	//UploadOOS(c)
}

//上传文件到本地
func UploadLocal(c *gin.Context) {
	w := c.Writer
	req := c.Request
	srcFile, head, err := req.FormFile("file")
	if err != nil {
		utils.RespFail(w, err.Error())
	}
	suffix := ".png"
	ofilName := head.Filename
	tem := strings.Split(ofilName, ".")
	if len(tem) > 1 {
		suffix = "." + tem[len(tem)-1]
	}
	fileName := fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	dstFile, err := os.Create("./asset/upload/" + fileName)
	if err != nil {
		utils.RespFail(w, err.Error())
	}
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		utils.RespFail(w, err.Error())
	}
	url := "./asset/upload/" + fileName
	utils.RespOK(w, url, "发送图片成功")
}

//上传文件到阿里云服务
func UploadOOS(c *gin.Context) {
	w := c.Writer
	req := c.Request
	srcFile, head, err := req.FormFile("file")
	if err != nil {
		utils.RespFail(w, err.Error())
	}
	suffix := ".png"
	ofilName := head.Filename
	tem := strings.Split(ofilName, ".")
	if len(tem) > 1 {
		suffix = "." + tem[len(tem)-1]
	}
	fileName := fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New(viper.GetString("oss.Endpoint"), viper.GetString("oss.AccessKeyId"), viper.GetString("oss.AccessKeySecret"))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket(viper.GetString("oss.Bucket"))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。
	err = bucket.PutObject(fileName, srcFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	url := "http://" + viper.GetString("oos.Bucket") + "." + viper.GetString("oos.EndPoint") + "/" + fileName
	utils.RespOK(w, url, "发送图片成功")
}
