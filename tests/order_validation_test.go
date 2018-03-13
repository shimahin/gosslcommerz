package tests

import (
	"fmt"
	"github.com/shimahin/gosslcommerz/core"
	"github.com/shimahin/gosslcommerz/models"
	"testing"
)

func TestOrderValidation(t *testing.T) {

	data := models.OrderValidationRequest{
		StoreId:       "test_shophobe",
		StorePassword: "test_shophobe@ssl",
		ValId:         "1711231900331kHP17lnrr9T8Gt",
		Format:        "json",
		V:             1,
	}

	sslcom := core.GetSslCommerzIPNListener(data.StoreId, data.StorePassword)

	orderResponse, err := sslcom.CheckValidation(&data)

	if err != nil {
		fmt.Println(err.Error())
	}
	t.Log(orderResponse)

}
