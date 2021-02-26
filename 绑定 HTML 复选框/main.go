package main

import "github.com/gin-gonic/gin"


type myForm struct {
	Colors []string `form:"colors[]"`
}


func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}

// form表单自动将相同key收集为数组
func main() {
	r := gin.Default()
	r.POST("/submit",formHandler)
	r.Run(":8080")
}