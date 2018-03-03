package core

import (
	"io/ioutil"
	"net/url"
	"strconv"

	"encoding/json"
	"github.com/shimahin/gosslcommerz/models"
	"net/http"

	"fmt"
)

const TRANSACTION_QUERY_URI = "validator/api/merchantTransIDvalidationAPI.php"

func (s *SslCommerz) TransactionQuery(transactionID []string, storeID string, storePass string) (*models.TransactionQueryResponse, error) {

	data := url.Values{}

	data.Set("tran_id", transactionID[0])
	data.Set("store_id", storeID)
	data.Set("store_passwd", storePass)

	u, _ := url.ParseRequestURI(SANDBOX_GATEWAY)
	u.Path = TRANSACTION_QUERY_URI
	//u.RawQuery = data.Encode()

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

	var resp models.TransactionQueryResponse
	json.Unmarshal(body, &resp)

	return &resp, nil
}
