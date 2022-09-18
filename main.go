package main

import (
	"go-blob/common"
	"go-blob/server"
)

func init() {
	//模板加载
	common.LoadTemplate()
}
func main() {
	server.App.Start("127.0.0.1", "8080")
}
