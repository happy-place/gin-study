package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)



func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	getwd, _ := os.Getwd()
	temp := filepath.Join(getwd,"多模板/simple")
	r.AddFromFiles("index", temp+"/templates/base.html", temp+"/templates/index.html")
	r.AddFromFiles("article", temp+"/templates/base.html", temp+"/templates/index.html", temp+"/templates/article.html")
	return r
}

func main() {
	router := gin.Default()
	router.HTMLRender = createMyRender()
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"title": "Html5 Template Engine",
		})
	})
	router.GET("/article", func(c *gin.Context) {
		c.HTML(200, "article", gin.H{
			"title": "Html5 Article Engine",
		})
	})
	router.Run(":8080")
}
