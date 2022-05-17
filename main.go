package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("./view/index.html", "./view/upload.html")
	router.Static("/static", "./static")
	router.StaticFS("/upload", http.Dir("upload"))

	router.StaticFile("/favicon.ico", "./static/favicon.ico")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/upload.html", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "upload.html", nil)
	})

	router.POST("/upload", func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file")
		if err := ctx.SaveUploadedFile(file, "./upload/"+file.Filename); err != nil {
			fmt.Println(err)
			ctx.String(http.StatusInternalServerError, "存储失败")
		}
		ctx.String(http.StatusOK, "上传成功")
	})

	router.Run(":3051")
}
