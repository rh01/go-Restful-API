package main

import (
	"log"
	"net/http"
	. "./routers" // 表示导入当前包
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

