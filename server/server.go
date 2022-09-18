package server

import (
	"go-blob/router"
	"log"
	"net/http"
)

var App = &BlogServer{}

type BlogServer struct {
}

func (*BlogServer) Start(ip, port string) {
	server := http.Server{
		Addr: ip + ":" + port,
	}
	//路由
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
