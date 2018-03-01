package main

import (
	//"github.com/shimahin/gosslcommerz/core"
	"fmt"
	//"net/url"
	"github.com/gin-gonic/gin"
)

func main() {

	//var sslcom core.SslCommerz

	router := gin.Default()

	router.POST("/ipn", func(c *gin.Context) {

		fmt.Println(c.Request.Form)
		fmt.Println(c.Request.URL.Query())

		fmt.Printf("%+v", c.Request)
	})
	router.Run(":8080")
}