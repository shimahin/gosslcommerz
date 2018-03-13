package tests

import (
	"testing"
	"github.com/shimahin/gosslcommerz/models"
	"github.com/shimahin/gosslcommerz/core"
	"fmt"
)

func TestTransactionQuery(t *testing.T)  {
	data := models.TransactionQueryRequest{
		StoreId : "test_shophobe",
		StorePasswd : "test_shophobe@ssl",
		TranId : "5a16c68b23783",
	}

	sslcom := core.GetSslCommerzIPNListener(data.StoreId , data.StorePasswd)

	tranResponse , err := sslcom.TransactionQuery(&data)
	if err != nil{
		fmt.Println(err.Error())
	}
	t.Log(tranResponse)
}
