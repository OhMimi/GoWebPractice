package model

import (
	"GoWebPractice/web01_db/utils"
	"fmt"
)

// User 結構體
type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

// AddUser 添加User的方法(預編譯方式)
func (user *User) AddUser() error {
	// 1.寫sql語句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	// 2.預編譯
	instmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("預編譯出現異常:", err)
		return err
	}
	// 3. 執行
	_, err = instmt.Exec("admin", "123456", "admin@gmail.com")
	if err != nil {
		fmt.Println("執行出現異常:", err)
		return err
	}
	return nil
}

// AddUser2 添加User的方法(不使用預編譯方式)
func (user *User) AddUser2() error {
	// 1.寫sql語句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	// 2.執行
	_, err := utils.Db.Exec(sqlStr, "admin2", "888888", "admin2@gmail.com")
	if err != nil {
		fmt.Println("執行出現異常:", err)
		return err
	}
	return nil
}
