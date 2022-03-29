/*
 * @Author: Pacific_D
 * @Date: 2022-03-29 14:46:27
 * @LastEditTime: 2022-03-29 15:19:45
 * @LastEditors: Pacific_D
 * @Description:
 * @FilePath: \blog\pkg\util\pagination.go
 */
package util

import (
	"github.com/PacificD/go-gin-example/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
