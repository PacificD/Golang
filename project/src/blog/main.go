/*
 * @Author: Pacific_D
 * @Date: 2022-03-29 15:24:46
 * @LastEditTime: 2022-03-29 16:17:43
 * @LastEditors: Pacific_D
 * @Description:
 * @FilePath: \blog\main.go
 */
package main

import (
	"fmt"
	"net/http"

	"github.com/PacificD/go-gin-example/pkg/setting"
	"github.com/PacificD/go-gin-example/routers"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
