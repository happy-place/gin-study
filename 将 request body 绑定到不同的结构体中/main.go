package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func shouldBindBodyWith(c *gin.Context){
	objA := formA{}
	objB := formB{}
	// 读取 c.Request.Body 并将结果存入上下文。
	if err := c.ShouldBindBodyWith(&objA, binding.JSON); err == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// 这时, 复用存储在上下文中的 body。
	} else if err := c.ShouldBindBodyWith(&objB, binding.JSON); err == nil {
		c.String(http.StatusOK, `the body should be formB JSON`)
		// 可以接受其他格式
	} else if errB := c.ShouldBindBodyWith(&objB, binding.XML); errB == nil {
		c.String(http.StatusOK, `the body should be formB XML`)
	} else {
		c.String(http.StatusOK, `bad body`)
	}
}

func shouldBind(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// c.ShouldBind 使用了 c.Request.Body，不可重用。
	if err := c.ShouldBind(&objA); err == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// 因为现在 c.Request.Body 是 EOF，所以这里会报错。
	} else if err = c.ShouldBind(&objB); err == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		c.String(http.StatusOK, `bad body`)
	}
}

func main() {
	r := gin.Default()

	r.POST("/shouldBind",shouldBind)
	r.POST("/shouldBindBodyWith",shouldBindBodyWith)

	r.Run(":8080")
}