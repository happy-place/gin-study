package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// callback 是 x
		// 将输出：x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
		// 使用 JSONP 向不同域的服务器请求数据。如果查询参数存在回调，则将回调添加到响应体中。
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}