package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	server := gin.Default()

	// 🤔 中间件之间的跳转你理解了么？
	server.Use(
		func(c *gin.Context) {
			fmt.Println(1)
			c.Next()
			fmt.Println(2)
		},
		func(c *gin.Context) {
			fmt.Println(3)
			c.Next()
			fmt.Println(4)
		},
		func(c *gin.Context) {
			fmt.Println(5)
			c.Next()
			fmt.Println(6)
		},
	)

	server.GET("/health", func(c *gin.Context) {
		fmt.Println("health")
		h := gin.H{
			"msg": "checked!",
		}
		c.JSON(http.StatusOK, h)
	})

	server.Run()
}
