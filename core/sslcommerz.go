package core

import (
	"net/url"
	"strings"
	"strconv"
	"net/http"
	"io/ioutil"
	"github.com/shimahin/gosslcommerz/models"
	"encoding/json"
)

const SANDBOX_GATEWAY =  "https://sandbox.sslcommerz.com"
const SESSION_URI = "gwprocess/v3/api.php"

type SslCommerz struct {
}

func(s *SslCommerz) CreateSession(data *url.Values) (*models.SessionResponse, error){

	u, _ := url.ParseRequestURI(SANDBOX_GATEWAY)
	u.Path = SESSION_URI
	sessionURL := u.String()

	client := &http.Client{}
	r, err := http.NewRequest("POST", sessionURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	body , _:= ioutil.ReadAll(response.Body)

	// SSL Commerz api is weired , it responses with 200 even if the request fails , WTF!
	// So check the struct not the status code !

	var resp *models.SessionResponse
	json.Unmarshal(body, &resp)

	return resp, nil

}
