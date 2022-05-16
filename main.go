package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello,work!")
	})
	r.POST("/testPost", getting)
	r.PUT("/testPut", putting)
	r.Run(":3510")
}

func getting(ctx *gin.Context) {
	fmt.Println("收到了Post请求")
	ctx.String(http.StatusOK, "收到了Post请求")
}
func putting(ctx *gin.Context) {
	fmt.Println("收到了Put请求")
	ctx.String(http.StatusOK, "收到了Put请求")
}
