package tests

import (
	"fmt"
	"github.com/shimahin/gosslcommerz/core"
	"github.com/shimahin/gosslcommerz/models"
	"testing"
)

func TestTransactionQuery(t *testing.T) {
	data := models.TransactionQueryRequest{
		StoreId:     "test_shophobe",
		StorePasswd: "test_shophobe@ssl",
		TranId:      "59C2A4F6432F8",
		Format:      "json",
	}

	x := core.SslCommerz{}

	tranResponse, err := x.TransactionQuery(&data)
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Log(tranResponse)
}
