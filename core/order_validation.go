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

func (s *SslCommerz) CheckValidation(valID []string, storeID string, storePass string, v string, format string) (*models.OrderValidationResponse, error) {

	data := url.Values{}

	data.Set("val_id", valID[0])
	data.Set("store_id", storeID)
	data.Set("store_passwd", storePass)
	data.Set("v", v)
	data.Set("format", format)

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
