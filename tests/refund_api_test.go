package tests

import (
	"testing"
	"github.com/shimahin/gosslcommerz/models"
	"github.com/shimahin/gosslcommerz/core"
	"fmt"
)

func TestRefundAPI(t *testing.T)  {

	data := models.RefundApiRequest{
		BankTranId: "1709162345070ANJdZV8LyI4cMw",
		StoreID: "test_shophobe",
		StorePasswd:"test_shophobe@ssl",
		RefundAmount:1000,
		RefundRemarks:"Remarks string",
		RefId:"Sample referance",
		Format:"json",
	}

	sslcom := core.GetSslCommerzIPNListener(data.StoreID , data.StorePasswd)

	fmt.Println(sslcom)

	refundResp , err := sslcom.InitiateRefunding(data.BankTranId,data.StoreID,data.StorePasswd,data.RefundAmount,data.RefundRemarks,data.RefId,data.Format)
	if err != nil{
		fmt.Println(err.Error())
	}

	t.Log(refundResp)


}
