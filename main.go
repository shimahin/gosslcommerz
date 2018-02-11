package main

import (
	"github.com/shimahin/gosslcommerz/core"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	var sslcom core.SslCommerz
/*
	data := url.Values{}

	data.Set("store_id", "testbox")
	data.Set("store_passwd", "qwerty")
	data.Set("total_amount", "100")
	data.Set("currency", "BDT")
	data.Set("tran_id", "REF123")
	data.Set("success_url", "http://yoursite.com/success.php")
	data.Set("fail_url", "http://yoursite.com/fail.php")
	data.Set("cancel_url", "http://yoursite.com/cancel.php")
	data.Set("cus_name", "Customer Name")
	data.Set("cus_email", "cust@yahoo.com")
	data.Set("cus_add1", "Dhaka")
	data.Set("cus_add2", "Dhaka")
	data.Set("cus_city", "Dhaka")
	data.Set("cus_state", "Dhaka")
	data.Set("cus_postcode", "1000")
	data.Set("cus_country", "Bangladesh")
	data.Set("cus_phone", "01711111111")
	data.Set("cus_fax", "01711111111")
	data.Set("ship_name", "Customer Name")
	data.Set("ship_add1 ", "Dhaka")
	data.Set("ship_add2", "Dhaka")
	data.Set("ship_city", "Dhaka")
	data.Set("ship_state", "Dhaka")
	data.Set("ship_postcode", "1000")
	data.Set("ship_country", "Bangladesh")
	data.Set("multi_card_name", "mastercard,visacard,amexcard,brac_visa,dbbl_visa,city_visa,ebl_visa,	sbl_visa,brac_master,dbbl_master,city_master,ebl_master,sbl_master,city_amex,qcash,dbbl_nexus,bankasia,abbank,ibbl,mtbl,bkash,dbblmobilebanking,city")
	data.Set("value_a", "ref001_A")
	data.Set("value_b", "ref002_B")
	data.Set("value_c", "ref003_C")
	data.Set("value_d", "ref004_D")

	body , err := sslcom.CreateSession(&data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(body)*/

	router := gin.Default()

	router.POST("/ipn", func(c *gin.Context) {

		data := sslcom.IPNListener(c.Request)

		fmt.Println(data)
	})
	router.Run(":8080")
}