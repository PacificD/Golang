package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//创建监听
	ip := "localhost"
	port := 8848
	address := fmt.Sprintf("%s:%d",ip,port)
	listener, err := net.Listen("tcp",address)
	//net.Listen("tcp",":8848") //简写，冒号前面默认是本机:127.0.0.1

	if err != nil{
		fmt.Println("net.Listen err: ",err)
		return
	}

	//server接收多个连接，每个连接可以接收处理多轮数据
	//主go程负责监听，子go程负责数据处理
	for{
		fmt.Println("监听中...")

		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("listener.Accpet err: ",err)
			return
		}
		fmt.Println("连接建立成功！")

		go handleFunc(conn)
	}
}

//处理具体业务的逻辑，需要将conn传递进来。
//每次一新连接，conn是不同的。
func handleFunc(conn net.Conn){
	for{
		//创建一个容器，用于接收读取到的数据
		buf := make([]byte,1204)
		//cnt: 真正读取client发来的数据的长度
		cnt, err := conn.Read(buf)

		if err != nil{
			fmt.Println("conn.Read err: ",err)
			return
		}
		fmt.Println("Client ===> Server,长度: ",cnt,"数据: ",string(buf[0:cnt]))

		//服务端对客户端请求进行响应
		//将数据转成大写
		upperData := strings.ToUpper(string(buf[0:cnt]))

		cnt, err = conn.Write([]byte(upperData))
		fmt.Println("Client <=== Server,长度: ",cnt,"数据: ",upperData)
	}
	conn.Close()
}