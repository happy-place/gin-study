package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 使用默认配置
//func main() {
//	r := gin.Default()
//	r.GET("/ping",func(ctx *gin.Context){
//		//time.Sleep(1*time.Second)
//		ctx.String(http.StatusOK,"pong")
//	})
//
//	http.ListenAndServe(":8080", r)
//}

// 使用自定义配置
func main() {
	r := gin.Default()

	r.GET("/ping",func(ctx *gin.Context){
		time.Sleep(2*time.Second)
		ctx.String(http.StatusOK,"pong")
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   3 * time.Second, // 超时配置
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}