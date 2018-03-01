package main

import (
	//"github.com/shimahin/gosslcommerz/core"
	"fmt"
	//"net/url"
	"github.com/gin-gonic/gin"
	"github.com/shimahin/gosslcommerz/core"
)

func main() {

	var sslcom core.SslCommerz
	sslcom.StorePass = "test_shophobe"
	sslcom.StoreID = "test_shophobe@ssl"

	router := gin.Default()

	router.POST("/ipn", func(c *gin.Context) {

		fmt.Println(c.Request.Form)
		fmt.Println(c.Request.URL.Query())

		ipn , err := sslcom.IPNListener(c.Request)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("%+v", ipn)
	})
	router.Run(":8080")
}