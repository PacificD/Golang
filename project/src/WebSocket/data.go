/*
 * @Author: Pacific_D
 * @Date: 2022-03-28 14:28:21
 * @LastEditTime: 2022-03-28 14:28:23
 * @LastEditors: Pacific_D
 * @Description:
 * @FilePath: \WebSocket\data.go
 */
package main

type Data struct {
	Ip       string   `json:"ip"`
	User     string   `json:"user"`
	From     string   `json:"from"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	UserList []string `json:"user_list"`
}
