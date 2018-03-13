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

const REFUNDING_URI = "validator/api/merchantTransIDvalidationAPI.php"

func (s *SslCommerz) InitiateRefunding(request *models.RefundApiRequest) (models.RefundResponse, error) {

	data := url.Values{}

	data.Set("bank_tran_id", request.BankTranId)
	data.Set("store_id", request.StoreID)
	data.Set("store_passwd", request.StorePasswd)
	data.Set("refund_amount", strconv.Itoa(request.RefundAmount))
	data.Set("refund_remarks", request.RefundRemarks)
	data.Set("refe_id", request.RefId)
	data.Set("format", request.Format)

	u, _ := url.ParseRequestURI(SANDBOX_GATEWAY)
	u.Path = REFUNDING_URI
	u.RawQuery = data.Encode()

	sessionURL := u.String()

	client := &http.Client{}

	var resp models.RefundResponse

	r, err := http.NewRequest("GET", sessionURL, nil)
	if err != nil {
		return resp, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	response, err := client.Do(r)
	if err != nil {
		return resp, err
	}

	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))

	// SSL Commerz api is weired , it responses with 200 even if the request fails , WTF!
	// So check the struct not the status code !

	json.Unmarshal(body, &resp)

	return resp, nil

}
