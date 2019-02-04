package main

import (
	"fmt"
	"github.com/gorilla/websocket"//这里使用的是 gorilla 的 websocket 库
	"log"
	"net/http"
	"time"
)

func main() {
	//websocket 的升级接口
	upgrader := websocket.Upgrader{}
	//go 标准库的 http 处理，这里处理的是根路径
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//通过 upgrader 将 http 连接升级为 websocket 连接
		connect, err := upgrader.Upgrade(writer, request, nil)
		if nil != err {
			log.Println(err)
			return
		}
		//离开作用域关闭连接，go 的常规操作
		defer connect.Close()

		//定时向客户端发送数据
		go tickWriter(connect)

		//启动数据读取循环，读取客户端发送来的数据
		for {
			//从 websocket 中读取数据
			//messageType 消息类型，websocket 标准
			//messageData 消息数据
			messageType, messageData, err := connect.ReadMessage()
			if nil != err {
				log.Println(err)
				break
			}
			switch messageType {
			case websocket.TextMessage://文本数据
				fmt.Println(string(messageData))
			case websocket.BinaryMessage://二进制数据
				fmt.Println(messageData)
			case websocket.CloseMessage://关闭
			case websocket.PingMessage://Ping
			case websocket.PongMessage://Pong
			default:

			}
		}
	})

	//go 标准库监听http
	err := http.ListenAndServe(":80", nil)
	if nil != err {
		log.Println(err)
		return
	}
}

func tickWriter(connect *websocket.Conn) {
	for {
		//向客户端发送类型为文本的数据
		err := connect.WriteMessage(websocket.TextMessage, []byte("from server to client"))
		if nil != err {
			log.Println(err)
			break
		}
		//休息一秒
		time.Sleep(time.Second)
	}
}