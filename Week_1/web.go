package main

import (
	"fmt"
	"net/http"
)

func myWeb(http.ResponseWriter, *http.Request) {
	fmt.Println("Welcome crucal.com")
}

func main() {
	http.HandleFunc("/", myWeb)

	fmt.Println("服务器即将开启，访问网址 http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器开启失败：", err)
	}
}
