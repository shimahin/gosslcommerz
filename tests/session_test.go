package tests

import (
	"fmt"
	"github.com/shimahin/gosslcommerz/core"
	"github.com/shimahin/gosslcommerz/models"
	"testing"
)

func TestCreateSession(t *testing.T) {

	var sessionResponse *models.SessionResponse

	data := models.RequestValue{
		StoreId:          "test_shophobe",
		StorePassword:    "test_shophobe@ssl",
		TotalAmount:      "150",
		Currency:         "BDT",
		TranID:           "CY123",
		SuccessURL:       "http://yoursite.com/success.php",
		FailUrl:          "http://yoursite.com/fail.php",
		CancelURL:        "http://yoursite.com/cancel.php",
		CustomerName:     "Customer Name",
		CustomerEmail:    "cus@yahoo.com",
		CustomerAdd1:     "Dhaka",
		CustomerAdd2:     "Dhaka",
		CustomerCity:     "Dhaka",
		CustomerState:    "Dhaka",
		CustomerPostCode: "1000",
		CustomerCountry:  "Bangladesh",
		CustomerPhone:    "01711111111",
		CustomerFax:      "01711111111",
		ShipName:         "Ship Name",
		ShipAdd1:         "Dhaka",
		ShipAdd2:         "Dhaka",
		ShipCity:         "Dhaka",
		ShipState:        "Dhaka",
		ShipPostCode:     "1000",
		ShipCountry:      "Bangladesh",
		MultiCardName:    []string{"mastercard", "visacard", "amexcard"},
		ValueA:           "ref001_A",
		ValueB:           "ref002_B",
		ValueC:           "ref003_C",
		ValueD:           "ref004_D",
	}

	sslcom := core.GetSslCommerzIPNListener(data.StoreId, data.StorePassword)
	/*
	   data.Set("multi_card_name", "mastercard,visacard,amexcard,brac_visa,dbbl_visa,city_visa,ebl_visa,sbl_visa,brac_master,dbbl_master,city_master,ebl_master,sbl_master,city_amex,qcash,dbbl_nexus,bankasia,abbank,ibbl,mtbl,bkash,dbblmobilebanking,city")
	*/

	fmt.Println(sslcom)

	sessionResponse, err := sslcom.CreateSession()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sessionResponse)

	paymentURL := sslcom.GetDirectPaymentURL()
	if paymentURL == "" {
		fmt.Println("nill")
	} else {
		fmt.Println(paymentURL)
	}
}
