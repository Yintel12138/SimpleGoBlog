package main

import (
	"log"
	"net/http"
	"simpleGoBlog/common"
	"simpleGoBlog/router"
)

func init() {
	//模板加载
	common.LoadTemplate()
}

func main() {
	// 程序入口
	Server := http.Server{
		Addr: "localhost:8080",
	}
	// 调用路由包下的路由函数，实现路由功能
	router.Router()
	if err := Server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
