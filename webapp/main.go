package main

import (
	"fmt"
	"net/http"
)

// 創建處理器函數
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!", r.URL.Path)
}

func main() {
	// 創建多路復用器
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	// http.HandleFunc("/", handler)

	// 創建路由
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("http.ListenAndServe err=%v\n", err)
	}
}
