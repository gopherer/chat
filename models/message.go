package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net"
	"net/http"
	"strconv"
	"sync"
)

//Message 消息
type Message struct {
	gorm.Model
	FormId   int64  //发送者
	TargetId int64  //接收者
	Type     int    //发送类型 1群聊 2私聊 3广播
	Media    int    //消息类型 1文字 2表情包 3图片 4音频
	Content  string //消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int //其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

//映射关系
var clientMap map[int64]*Node = make(map[int64]*Node)

//读写锁
var rwLocker sync.RWMutex

//Chat 需要：发送者ID，接收者ID，消息类型，发送的内容，发送类型
func Chat(write http.ResponseWriter, request *http.Request) {
	//1、获取参数并校验token等合法性
	//token := query.Get("token")
	query := request.URL.Query()
	Id := query.Get("userId")
	userId, _ := strconv.ParseInt(Id, 10, 64)
	//msgType := query.Get("type")
	//targetId := query.Get("targetId")
	//context := query.Get("context")
	isvalida := true //checkToken() 待。。。。。。
	conn, err := (&websocket.Upgrader{
		//token 校验
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(write, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//2、获取conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	//3、用户关系
	//4、userId 跟node 绑定并加锁
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()
	//5、完成发送逻辑
	go sendProc(node)
	//6、完成接受逻辑
	go recvProc(node)
	sendMsg(userId, []byte("欢迎进入聊天室系统"))
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		broadMsg(data)
		fmt.Println("[ws] <<<< ", data)
	}
}

var udpsendchan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpsendchan <- data
}

func init() {
	go udpSendProc()
	go udpRecvProc()
}

//完成udp数据发送携程
func udpSendProc() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 0, 255),
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(conn)

	for {
		select {
		case data := <-udpsendchan:
			_, err := conn.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

//完成udp数据接受携程
func udpRecvProc() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(conn)
	for {
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		dispatch(buf[0:n])
	}
}

//后端调度逻辑处理
func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case 1: //私信
		sendMsg(msg.TargetId, data)
		//case 2: //群发
		//	sendGroup()
		//case 3: //广播
		//	sendAllMsg()
		//case 4:
		//
	}
}

func sendMsg(userId int64, msg []byte) {
	rwLocker.RLock()
	node, ok := clientMap[userId]
	rwLocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
