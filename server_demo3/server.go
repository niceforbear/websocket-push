package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"push_by_websocket/server_demo3/impl"
	"time"
)

var (
	upgrader = websocket.Upgrader {
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err error
		conn *impl.Connection
		data []byte
	)
	// Upgrade: ws
	// get long conn
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	if conn, err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}

	go func() {
		var (
			err error
		)

		for {
			if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(time.Second * 1)
		}
	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}

		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}

	ERR:
		//TODO
}

func main() {
	var (
		err error
	)

	fmt.Print("start server")
	http.HandleFunc("/ws", wsHandler)

	if err = http.ListenAndServe("0.0.0.0:7777", nil); err != nil {
		fmt.Print(err)
	}
}
