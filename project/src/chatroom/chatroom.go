package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp",":8848")
	if err != nil{
		fmt.Println("net.Listen err: ",err)
		return
	}
	fmt.Println("服务器启动成功!")

	for {
		fmt.Println("======> 主go程监听中...")
		//监听
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("listener.Accept err: ",err)
			return
		}
		fmt.Println("建立连接成功...")

		//启动处理业务的go程
		go handler(conn)
	}
}

//处理具体业务
func handler(conn net.Conn){
	for {
		fmt.Println("启动业务...")
		buf := make([]byte, 1024)
		//读取客户端发送过来的数据
		cnt, err := conn.Read(buf)
		if err != nil{
			fmt.Println("conn.Read err: ",err)
			return
		}
		fmt.Println("服务器接收客户端发送过来的数据为: ",string(buf[:cnt]), ", cnt: ",cnt)
	}
}

