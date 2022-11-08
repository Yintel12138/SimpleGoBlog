package router

import "net/http"

func Router() {
	// 优化一下路由功能
	http.HandleFunc("/", index)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
