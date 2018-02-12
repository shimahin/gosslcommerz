package core

import "net/http"

func(s *SslCommerz) IPNListener(request *http.Request) map[string]interface{}{

	// IPN sends the data to IPN URL as Query Request URI so ,
	data := request.URL.Query()
	ipnResponse := make(map[string]interface{})

	for k, v := range data {
		ipnResponse[k] = v
	}
	return ipnResponse
}