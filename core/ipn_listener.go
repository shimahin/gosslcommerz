package core

import "net/http"

func(s *SslCommerz) IPNListener(request *http.Request) map[string]interface{}{

	data := request.Form

	var ipnResponse map[string]interface{}

	for k, v := range data {
		ipnResponse[k] = v
	}

	return ipnResponse
}