/*
 * @Author: Pacific_D
 * @Date: 2022-03-27 14:01:10
 * @LastEditTime: 2022-03-27 14:01:10
 * @LastEditors: Pacific_D
 * @Description:
 * @FilePath: \src\WebSocket\server.go
 */
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/ws", myws)
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
}
