package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	fmt.Println("これはテストです")
	r :=gin.Default()
	r.LoadHTMLFiles("./template/index.html")
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"message":"pong",
		})
	})

	//サーバーを開く
	r.Run()
}