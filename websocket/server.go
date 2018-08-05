package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		// 定义允许跨域请求
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		conn *websocket.Conn
		err  error
		data []byte
	)
	// 1、完成握手应答
	conn, err = upgrader.Upgrade(w, r, nil)

	if err != nil {
		//http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	// 2、建立连接
	for {
		// 读取消息
		//msgType,data,err = conn.ReadMessage()
		_, data, err = conn.ReadMessage() // 没有使用到的变量可以使用 “_”代替
		if err != nil {
			// 关闭连接
			goto ERR
		}
		// 打印接受到的消息
		fmt.Printf("%s \n", []byte(data))
		// 发送(写入)消息
		err = conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			goto ERR
		}
	}
ERR:
	// 安全线程
	conn.Close()
}

func main() {
	// http://127.0.0.1:8888/ws
	fmt.Println("websocket is start ...")
	http.HandleFunc("/ws", wsHandler)
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
