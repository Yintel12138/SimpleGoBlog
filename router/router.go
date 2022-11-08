package router

import (
	"net/http"
	"simpleGoBlog/api"
	"simpleGoBlog/views"
)

func Router() {
	// 1.页面 viewers
	http.HandleFunc("/", views.HTML.Index)
	// 2.数据
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	// 3.静态资源
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
