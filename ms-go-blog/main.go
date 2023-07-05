package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte("welcome crucal.com!"))
	if err != nil {
		return
	}
}

func main() {
	// 程序入口，一个项目 只能有一个入口
	// Web程序，http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// 响应
	http.HandleFunc("/", index)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
