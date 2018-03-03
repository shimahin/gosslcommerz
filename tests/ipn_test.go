package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shimahin/gosslcommerz/core"
	"testing"
)

func TestIPN(t *testing.T) {

	sslcom := core.GetSslCommerzIPNListener("test_shophobe", "test_shophobe@ssl")

	router := gin.Default()

	router.POST("/ipn", func(c *gin.Context) {

		fmt.Println(c.Request.Form)
		fmt.Println(c.Request.URL.Query())

		fmt.Printf("%+v", c.Request)

		data, err := sslcom.IPNListener(c.Request)
		if err != nil {
			t.Error(err.Error())
		}

		x := core.SslCommerz{}

		orderResp, err := x.CheckValidation(data["val_id"].([]string), sslcom.StoreId, sslcom.StorePass, "1", "json")

		if err != nil {
			t.Error(err.Error())
		}

		t.Log(orderResp)

		transactionQuery , err := x.TransactionQuery(data["tran_id"].([]string), sslcom.StoreId, sslcom.StorePass)
		fmt.Println(transactionQuery)

	})
	router.Run(":8080")
}
