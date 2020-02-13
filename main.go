package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	addr := "0.0.0.0:8080"

	// 创建gin实例
	r := gin.Default()

	// 处理静态资源
	r.Static("/css", "./css")
	r.Static("/js", "./js")
	r.Static("/file", "./file")
	r.StaticFile("/digital.ttf", "./digital.ttf")
	r.StaticFile("/favicon.ico", "./favicon.ico")
	r.StaticFile("/", "./index.html")

	// 启动服务
	log.Fatalln(r.Run(addr))
}
