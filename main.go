package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
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

	//这个处理可以兼容单文件和多文件
	router.POST("/upload", func(ctx *gin.Context) {
		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.String(http.StatusInternalServerError, "获取信息错误")
			return
		}

		files := form.File["file"]
		fmt.Println("文件数量", files, len(files))

		if len(files) == 0 {
			ctx.String(http.StatusBadRequest, "文件不能为空")
			return
		}

		for _, file := range files {
			fmt.Println("看看信息", file.Size)
			fmt.Println("看看信息2", file.Header)
			fmt.Println("看看信息3", file.Header.Get("Content-Type"))
			fmt.Println("看看信息4", file.Header.Get("Content-Disposition"))
			//1MB
			if file.Size > 1<<20 {
				ctx.String(http.StatusInternalServerError, "文件过大")
				return
			}
			if !lo.Contains[string]([]string{"image/png", "text/xml"}, file.Header.Get("Content-Type")) {
				ctx.String(http.StatusInternalServerError, "只能上传图片")
				return
			}
			if err := ctx.SaveUploadedFile(file, "./upload/"+file.Filename); err != nil {
				fmt.Println(err)
				ctx.String(http.StatusInternalServerError, "存储失败")
				return
			}
		}

		ctx.String(http.StatusOK, "上传成功")
	})

	router.Run(":3051")
}
