package main

import (
	"yutooaisdk/define"
	"yutooaisdk/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 定义一个简单的GET请求处理器
	r.GET(define.GetVersion, internal.GetVersion)

	r.POST(define.Init, internal.Init)

	r.POST(define.UnInit, internal.UnInit)

	r.POST(define.Activate, internal.Activate)

	// 启动HTTP服务
	r.Run(":8080")
}
