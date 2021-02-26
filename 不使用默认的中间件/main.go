package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建router
	r := gin.Default() // Default默认使用Logger （彩色日志） 和 Recovery中间件
	//r := gin.New()

	// 注册handler
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		// 为非ascii 文字提供编码支持
		c.AsciiJSON(http.StatusOK, data)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

