package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
	"os"
	"github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})

	getwd, _ := os.Getwd()
	temp := filepath.Join(getwd, "HTML 渲染/自定义模板/templates/raw.tmpl")
	router.LoadHTMLFiles(temp)

	// Date: {[{.now | formatAsDate}]} 模板中控制调用函数
	router.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})

	router.Run(":8080")
}
