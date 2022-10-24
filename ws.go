package main

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	NewConnection()
}

func NewConnection() {
	conn, resp, err := websocket.DefaultDialer.Dial("ws://localhost:5005/ws", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)

	go (func() {
		for {
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				fmt.Println(err)
				return
			}
			time.Sleep(time.Second * 30)
		}
	})()

	go (func() {
		time.Sleep(time.Second * 5)
		err = conn.WriteMessage(websocket.TextMessage, []byte("echo"))
		if err != nil {
			fmt.Println(err)
		}
	})()

	for {
		messageType, bz, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(messageType)
		fmt.Println(string(bz))
	}
}
