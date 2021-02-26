package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	router := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		getwd, _ := os.Getwd()
		dst := filepath.Join(getwd,"单文件上传", file.Filename)

		filepath.Clean(dst)

		// 上传文件至指定目录
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			panic(err)
		}

		data := map[string]interface{}{
			"status": fmt.Sprintf("'%s' uploaded!", file.Filename),
		}
		c.AsciiJSON(http.StatusOK,data)
	})
	router.Run(":8080")
}
