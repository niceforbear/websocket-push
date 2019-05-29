package main

import (
	"fmt"
	"net/http"
)


func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	if _, err = w.Write([]byte("hello")); err != nil {
		fmt.Println("error")
	}

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
