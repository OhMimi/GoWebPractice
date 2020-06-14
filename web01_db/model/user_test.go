package model

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	fmt.Println("測試添加用戶: ")
	user := &User{}
	// 調用添加用戶的方法
	err := user.AddUser()
	if err != nil {
		fmt.Println("user.AddUser 有錯誤 err = ", err)
	}
	err = user.AddUser2()
	if err != nil {
		fmt.Println("user.AddUser2 有錯誤 err = ", err)
	}
}
