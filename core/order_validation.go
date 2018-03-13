package core

import (
	"encoding/json"

	"fmt"
	"github.com/shimahin/gosslcommerz/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const ORDER_VALIDATION_URI = "validator/api/validationserverAPI.php"

func (s *SslCommerz) CheckValidation(request *models.OrderValidationRequest) (*models.OrderValidationResponse, error) {

	data := url.Values{}

	data.Set("val_id", request.ValId)
	data.Set("store_id", request.StoreId)
	data.Set("store_passwd", request.StorePassword)
	data.Set("v", strconv.Itoa(request.V))
	data.Set("format", request.Format)

	u, _ := url.ParseRequestURI(SANDBOX_GATEWAY)
	u.Path = ORDER_VALIDATION_URI
	u.RawQuery = data.Encode()

	sessionURL := u.String()

	client := &http.Client{}
	r, err := http.NewRequest("GET", sessionURL, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))

	// SSL Commerz api is weired , it responses with 200 even if the request fails , WTF!
	// So check the struct not the status code !

	var resp models.OrderValidationResponse
	json.Unmarshal(body, &resp)

	return &resp, nil
}

/*
Sample request

curl -X POST localhost:8080 -d 'tran_id=5a16c68b23783&val_id=1711231900331kHP17lnrr9T8Gt&amount=100&card_type=VISA-Dutch Bangla&store_amount=97&card_no=425272XXXXXX3456&
bank_tran_id=1711231900331S0R8atkhAZksmM&status=VALID&tran_date=2017-11-23 18:59:55&currency=BDT&card_issuer=Standard Chartered Bank&card_brand=VISA&
card_issuer_country=Bangladesh&card_issuer_country_code=BD&store_id=testbox&verify_sign=8070c0cefed9e629b01100d8a92afda2&verify_key=amount,bank_tran_id,base_fair,card_brand,card_issuer,card_issuer_country,card_issuer_country_code,card_no,card_type,currency,currency_amount,currency_rate,currency_type,risk_level,risk_title,status,store_amount,store_id,tran_date,tran_id,val_id,value_a,value_b,value_c,value_d&
cus_fax=01711111111&currency_type=BDT&currency_amount=100.00&currency_rate=1.0000&base_fair=0.00&value_a=ref001_A&value_b=ref002_B&value_c=ref003_C&value_d=ref004_D&
risk_level=0&risk_title=Safe'

 */
