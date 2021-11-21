package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//注册路由
	http.HandleFunc("/user",func(writer http.ResponseWriter, request *http.Request){
		fmt.Println("user用户请求详情request: ",request)
		_,_ = io.WriteString(writer,"/user")
	})

	http.HandleFunc("/name",func(writer http.ResponseWriter, request *http.Request){
		fmt.Println("name用户请求详情request: ",request)
		_,_ = io.WriteString(writer,"/name")
	})

	if err := http.ListenAndServe("localhost:8080",nil); err != nil{
		fmt.Println("http server start fail,err :",err)
		return
	}
}
