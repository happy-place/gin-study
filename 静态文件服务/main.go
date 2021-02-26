package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	router := gin.Default()
	getwd, _ := os.Getwd()
	base := filepath.Join(getwd,"静态文件服务/static")
	router.Static("/assets", base+"/assets")
	router.StaticFS("/more_static", http.Dir(filepath.Join(getwd,"静态文件服务/fs"))) // 静态资源服务器
	router.StaticFile("/favicon.ico", filepath.Join(getwd,"静态文件服务/static/favicon.ico")) // 图标

	router.GET("/test",func(ctx *gin.Context){
		ctx.String(http.StatusOK,"hello")
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8080")
}
