/*
 * @Author: Pacific_D
 * @Date: 2022-03-30 14:05:25
 * @LastEditTime: 2022-03-30 14:26:56
 * @LastEditors: Pacific_D
 * @Description:
 * @FilePath: \blog\models\auth.go
 */
package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}
