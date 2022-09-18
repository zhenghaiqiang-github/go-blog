package router

import (
	"go-blob/api"
	"go-blob/views"
	"net/http"
)

func Router() {
	//1 页面 views
	//2 api数据 json
	//3 静态资源 static

	http.HandleFunc("/", views.HTML.Index) //处理请求
	//http://127.0.0.1:8080/c/1
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/login", views.HTML.Login)
	//http://127.0.0.1:8080/p/7.html
	http.HandleFunc("/p/", views.HTML.Detail)
	http.HandleFunc("/writing", views.HTML.Writing)
	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)

	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
	http.HandleFunc("/api/v1/qiniu/token", api.API.QiniuToken)

	http.HandleFunc("/api/v1/login", api.API.Login)

	//静态资源处理:我就是没加这玩意，所以什么css格式都没有
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

}
