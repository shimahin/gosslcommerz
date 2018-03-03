package core

import (
	"encoding/json"
	"github.com/shimahin/gosslcommerz/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"fmt"
)

const SANDBOX_GATEWAY = "https://sandbox.sslcommerz.com"
const SESSION_URI = "gwprocess/v3/api.php"

type SslCommerz struct {
	StorePass string
	StoreId   string
	req       *models.RequestValue
	resp      *models.SessionResponse
}

func GetSslCommerzSession(req *models.RequestValue) *SslCommerz {
	return &SslCommerz{
		req: req,
	}
}

func (s *SslCommerz) CreateSession() (*models.SessionResponse, error) {

	data := url.Values{}

	data.Set("store_id", s.req.StoreId)
	data.Set("store_passwd", s.req.StorePassword)
	data.Set("total_amount", s.req.TotalAmount)
	data.Set("currency", s.req.Currency)
	data.Set("tran_id", s.req.TranID)
	data.Set("success_url", s.req.SuccessURL)
	data.Set("fail_url", s.req.FailUrl)
	data.Set("cancel_url", s.req.CancelURL)
	data.Set("cus_name", s.req.CustomerName)
	data.Set("cus_email", s.req.CustomerEmail)
	data.Set("cus_add1", s.req.CustomerAdd1)
	data.Set("cus_add2", s.req.CustomerAdd2)
	data.Set("cus_city", s.req.CustomerCity)
	data.Set("cus_state", s.req.CustomerState)
	data.Set("cus_postcode", s.req.CustomerPostCode)
	data.Set("cus_country", s.req.CustomerCountry)
	data.Set("cus_phone", s.req.CustomerPhone)
	data.Set("cus_fax", s.req.CustomerFax)
	data.Set("ship_name", s.req.ShipName)
	data.Set("ship_add1 ", s.req.ShipAdd1)
	data.Set("ship_add2", s.req.ShipAdd2)
	data.Set("ship_city", s.req.ShipCity)
	data.Set("ship_state", s.req.ShipState)
	data.Set("ship_postcode", s.req.ShipPostCode)
	data.Set("ship_country", s.req.ShipCountry)
	data.Set("multi_card_name", strings.Join(s.req.MultiCardName, ","))
	data.Set("value_a", s.req.ValueA)
	data.Set("value_b", s.req.ValueB)
	data.Set("value_c", s.req.ValueC)
	data.Set("value_d", s.req.ValueD)

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

	body, _ := ioutil.ReadAll(response.Body)

	// SSL Commerz api is weired , it responses with 200 even if the request fails , WTF!
	// So check the struct not the status code !

	var resp models.SessionResponse
	json.Unmarshal(body, &resp)
	s.resp = &resp

	return &resp, nil
}

func (s *SslCommerz) GetRedirectGatewayURL() string {
	return s.resp.RedirectGatewayURL
}

func (s *SslCommerz) GetDirectPaymentURL() string {
	return s.resp.DirectPaymentURL
}