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

	// 注：默认GOPATH设置为项目的src目录，此处多加一层目录作为模块目录
	//temp := filepath.Join(os.Getenv("GOPATH"),"templates/*")

	temp := filepath.Join(getwd,"HTML 渲染/多层目录/templates/**/*") // 扫描 template目录下包括文件夹在内所有资源
	router.LoadHTMLGlob(temp)

	// 注： "posts/index.tmpl" 对应 posts/index.tmpl中的 {{ define "posts/index.tmpl" }}
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})

	router.Run(":8080")
}
