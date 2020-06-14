package main

import (
	"fmt"
	"net/http"
)

// 創建處理器函數
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "測試Http協議")
}

func main() {
	// 調用處理器請求
	http.HandleFunc("/http", handler)
	// 路由
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http.ListenAndServe err=%v\n", err)
	}
}
