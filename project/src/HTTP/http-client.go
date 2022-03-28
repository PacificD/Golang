/*
 * @Author: Pacific_D
 * @Date: 2022-03-27 10:21:53
 * @LastEditTime: 2022-03-27 12:06:35
 * @LastEditors: Pacific_D
 * @Description:
 * @FilePath: \src\HTTP\http-client.go
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//http包
	client := http.Client{}
	resp, err := client.Get("http://localhost:8080/name")
	if err != nil {
		fmt.Println("Client.Get err: ", err)
		return
	}

	ct := resp.Header.Get("Content-Type")
	date := resp.Header.Get("Date")
	server := resp.Header.Get("Server")

	url := resp.Request.URL
	code := resp.StatusCode
	status := resp.Status

	body := resp.Body
	readBodyStr, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("read body err: ", err)
		return
	}

	fmt.Println(ct)
	fmt.Println(date)
	fmt.Println(server)
	fmt.Println(url)
	fmt.Println(code)
	fmt.Println(status)
	fmt.Println("body string: ", string(readBodyStr))
}
