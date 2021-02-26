package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
	// 注：默认GOPATH设置为项目的src目录，此处多加一层目录作为模块目录
	router := gin.Default()

	getwd, _ := os.Getwd()
	temp := filepath.Join(getwd,"HTML 渲染/单层目录/templates/*") // 扫描 template目录下资源
	//temp := filepath.Join(os.Getenv("GOPATH"),"templates/*")

	router.LoadHTMLGlob(temp) // 通配方式加载
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html") // 枚举方式加载
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.Run(":8080")
}
