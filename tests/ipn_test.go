package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shimahin/gosslcommerz/core"
	"testing"
	//"github.com/shimahin/gosslcommerz/models"
)

func TestIPN(t *testing.T) {

	sslcom := core.GetSslCommerzIPNListener("test_shophobe", "test_shophobe@ssl")

	router := gin.Default()

	router.POST("/ipn", func(c *gin.Context) {

		fmt.Println(c.Request.Form)
		fmt.Println(c.Request.URL.Query())

		fmt.Printf("%+v", c.Request)

		ipnResp, err := sslcom.IPNListener(c.Request)
		if err != nil {
			t.Error(err.Error())
		}

		fmt.Println("IPN TESTING RESPONSE : ", ipnResp)

	})
	router.Run(":8080")
}
