package tests

import (
	"github.com/shimahin/gosslcommerz/models"
	"testing"
	"github.com/shimahin/gosslcommerz/core"
	"fmt"
)

func TestOrderValidation(t *testing.T) {

	data := models.OrderValidationRequest{
		StoreId:       "test_shophobe",
		StorePassword: "test_shophobe@ssl",
		ValId:         "1711231900331kHP17lnrr9T8Gt",
		Format:        "json",
		V:             1,
	}

	sslcom := core.GetSslCommerzIPNListener(data.StoreId , data.StorePassword)

	orderResponse , err := sslcom.CheckValidation(&data)

	if err != nil{
		fmt.Println(err.Error())
	}
	t.Log(orderResponse)

}
