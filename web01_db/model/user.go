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

// GetUserById 根據用戶的id從數據庫中查詢一筆紀錄
func (user *User) GetUserById(id int) (findUser *User, err error) {
	// 1.寫sql語句
	sqlStr := "select id,username,password,email from users where id = ?"
	// 2.執行
	row := utils.Db.QueryRow(sqlStr, user.ID)
	// 聲明變量
	var (
		username string
		password string
		email    string
	)

	err = row.Scan(&id, &username, &password, &email)
	if err != nil {
		fmt.Println("row.Scan() err = ", err)
		return
	}
	// 創建一個USER實例,
	findUser = &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}

	return
}

// GetUsers 從數據庫中查詢所有紀錄
func (user *User) GetUsers() (findUsers []*User, err error) {
	// 1.寫sql語句
	sqlStr := "select id,username,password,email from users"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		fmt.Println("utils.Db.Query() err = ", err)
		return
	}
	// rows.Next 可查看下一筆資料是否存在
	for rows.Next() {
		// 聲明變量
		var (
			id       int
			username string
			password string
			email    string
		)
		err = rows.Scan(&id, &username, &password, &email)
		if err != nil {
			fmt.Println("rows.Scan() err = ", err)
			return
		}
		// 創建一個USER實例,
		findUser := &User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
		}
		findUsers = append(findUsers, findUser)
	}

	return
}
