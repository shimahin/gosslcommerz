package core

import (
	"io/ioutil"
	"net/url"
	"strconv"

	"encoding/json"
	"fmt"
	"github.com/shimahin/gosslcommerz/models"
	"net/http"
	"strings"
)

const TRANSACTION_QUERY_URI = "validator/api/merchantTransIDvalidationAPI.php"

func (s *SslCommerz) TransactionQuery(request *models.TransactionQueryRequest) (*models.TransactionQueryResponse, error) {

	data := url.Values{}

	data.Set("tran_id", request.TranId)
	data.Set("store_id", request.StoreId)
	data.Set("store_passwd", request.StorePasswd)
	data.Set("v", request.V)
	data.Set("format", request.Format)

	u, _ := url.ParseRequestURI(SANDBOX_GATEWAY)
	u.Path = TRANSACTION_QUERY_URI
	u.RawQuery = data.Encode()

	sessionURL := u.String()

	client := &http.Client{}

	var resp models.TransactionQueryResponse

	r, err := http.NewRequest("GET", sessionURL, strings.NewReader(data.Encode()))
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

	//fmt.Println(string(body))

	// SSL Commerz api is weired , it responses with 200 even if the request fails , WTF!
	// So check the struct not the status code !

	json.Unmarshal(body, &resp)

	fmt.Println(resp)
	return &resp, nil
}
