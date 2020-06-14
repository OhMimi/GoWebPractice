package model

import (
	"fmt"
	"os"
	"testing"
)

// 簡易測試指令: go test
// 詳細測試指令: go test -v

// TestMain 函數可以在測試函數執行前做一些其他操作
func TestMain(m *testing.M) {
	fmt.Println("測試開始:")
	// 透過m.Run()來執行測試函數
	m.Run()
	os.Exit(0)
}

func TestUser(t *testing.T) {
	fmt.Println("開始測試User中的相關方法")
	// 透過t.Run()來執行子測試程序
	t.Run("測試添加用戶: ", testAddUser)
	t.Run("測試獲取一筆用戶: ", testGetUserById)
	t.Run("測試獲取所有用戶: ", testGetUsers)
}

// 如果函數名不是以TestXxx開頭,那該函數不執行
func testAddUser(t *testing.T) {
	fmt.Println("測試添加用戶: ")
	// user := &User{}
	// 調用添加用戶的方法
	// err := user.AddUser()
	// if err != nil {
	// 	fmt.Println("user.AddUser 有錯誤 err = ", err)
	// return
	// }
	// err = user.AddUser2()
	// if err != nil {
	// 	fmt.Println("user.AddUser2 有錯誤 err = ", err)
	// return
	// }
}

// // 如果函數名不是以TestXxx開頭,那該函數不執行
// func TestAddUser(t *testing.T) {
// 	fmt.Println("測試添加用戶: ")
// 	user := &User{}
// 	// 調用添加用戶的方法
// 	err := user.AddUser()
// 	if err != nil {
// 		fmt.Println("user.AddUser 有錯誤 err = ", err)
// return
// 	}
// 	err = user.AddUser2()
// 	if err != nil {
// 		fmt.Println("user.AddUser2 有錯誤 err = ", err)
// return
// 	}
// }

// 測試獲取一筆用戶紀錄
func testGetUserById(t *testing.T) {
	fmt.Println("測試查詢一筆紀錄: ")
	user := &User{
		ID: 1,
	}
	// 調用獲取User的方法
	findUser, err := user.GetUserById(user.ID)
	if err != nil {
		fmt.Println("user.GetUserById 有錯誤 err = ", err)
		return
	}
	fmt.Printf("得到的User信息是: ID=%d 用戶名=%s 密碼=%s email=%s\n", findUser.ID, findUser.Username, findUser.Password, findUser.Email)
}

// 測試獲取所有用戶紀錄
func testGetUsers(t *testing.T) {
	fmt.Println("測試查詢所有紀錄: ")
	user := &User{}
	// 調用 獲取所有用戶方法
	users, err := user.GetUsers()
	if err != nil {
		fmt.Println("user.GetUsers 有錯誤 err = ", err)
		return
	}
	// 遍歷切片
	for i, u := range users {
		fmt.Printf("得到的第%d個User信息是: ID=%d 用戶名=%s 密碼=%s email=%s\n", i+1, u.ID, u.Username, u.Password, u.Email)
	}
}
