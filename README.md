# chat


## 项目导入

~~~
go env -w GO111MODULE=on

go env -w GOPROXY=https://goproxy.cn,direct

git clone https://github.com/gopherer/chat.git

go mod tidy
~~~

## 目录结构

- *asset*:存放项目图片，图标，插件，页面样式等 
- *asset/upload*：存储用户上传的图片
- *config*:存放项目配置信息
- *docs*:存放swagger配置信息
- *models*:存放用户、群组、消息结构体信息
- *router*:存放项目路由
- *service*:存放对各个api接口的具体实现
- *test*:存放测试文件
- *tmp*:存放临时文件
- *utils*:存放项目特定功能所需的go文件的实现
- *views*:存放项目页面信息

## Air 自动重载(.air.toml)
使用以下命令来安装 air ：
~~~
GO111MODULE=on  go install github.com/cosmtrek/air@latest
~~~

安装成功后使用以下命令检查下：

~~~
air -v

  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ , built with Go
~~~

在我们的 chat 项目根目录运行以下命令：
~~~
air
~~~

![](https://github.com/gopherer/chat/blob/main/MDPhoto/air.png)

## 项目功能

- 用户注册
- 用户登录
- 添加好友
- 创建群组
- 添加群组
- 单聊（文本、表情包、图片、语言）
- 群聊（文本、表情包、图片、语言）
- 心跳检测
- 基于ChatGPT实现的AI助理

部分页面截图

主要功能

![](https://github.com/gopherer/chat/blob/main/MDPhoto/main.png)

![](https://github.com/gopherer/chat/blob/main/MDPhoto/chat.png)

![](https://github.com/gopherer/chat/blob/main/MDPhoto/me.png)

拓展功能

![](https://github.com/gopherer/chat/blob/main/MDPhoto/chatgpt.png)

![](https://github.com/gopherer/chat/blob/main/MDPhoto/chatgpt2.png)
