package core

import (
	"net/url"
	"strconv"
	"net/http"
	"github.com/shimahin/gosslcommerz/models"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

const REFUNDING_URI  = "validator/api/merchantTransIDvalidationAPI.php"

func (s *SslCommerz) InitiateRefunding(bank_tran_id string , store_id string , store_passwd string , refund_amount int , refund_remarks string , ref_id string , format string) (models.RefundResponse,error) {

	data := url.Values{}

	data.Set("bank_tran_id",bank_tran_id)
	data.Set("store_id",store_id)
	data.Set("store_passwd",store_passwd)
	data.Set("refund_amount",strconv.Itoa(refund_amount))
	data.Set("refund_remarks",refund_remarks)
	data.Set("refe_id",ref_id)
	data.Set("format",format)

	u , _ := url.ParseRequestURI(SANDBOX_GATEWAY)
	u.Path = REFUNDING_URI
	u.RawQuery = data.Encode()

	sessionURL := u.String()

	client := &http.Client{}

	var resp models.RefundResponse

	r , err := http.NewRequest("GET",sessionURL,nil)
	if err != nil{

		return resp,err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	response, err := client.Do(r)
	if err != nil {
		return resp , err
	}

	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))

	// SSL Commerz api is weired , it responses with 200 even if the request fails , WTF!
	// So check the struct not the status code !


	json.Unmarshal(body, &resp)


	return resp, nil

}