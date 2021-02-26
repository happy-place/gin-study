package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// 简单的路由组: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", handler)
		v1.POST("/submit", handler)
		v1.POST("/read", handler)
	}

	// 简单的路由组: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", handler)
		v2.POST("/submit", handler)
		v2.POST("/read", handler)
	}

	router.Run(":8080")
}

func handler(ctx *gin.Context){
	ctx.String(http.StatusOK,ctx.Request.URL.Path)
}