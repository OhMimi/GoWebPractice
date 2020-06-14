package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct{}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "通過自己創建的處理器處理請求!", r.URL.Path)
	fmt.Fprintln(w, "通過自己詳細配置創建的處理器處理請求!", r.URL.Path)
}

func main() {
	myHandler := MyHandler{}

	// http.Handle("/myHandler", &myHandler)
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	fmt.Printf("http.ListenAndServe err=%v\n", err)
	// }

	// 創建Server結構體
	sever := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler,
		ReadTimeout: 2 * time.Second,
	}
	err := sever.ListenAndServe()
	if err != nil {
		fmt.Printf("sever.ListenAndServe err=%v\n", err)
	}
}
