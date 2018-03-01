package models

type SessionResponse struct {
	Status       string `json:"status"`
	Failedreason string `json:"failedreason"`
	Sessionkey   string `json:"sessionkey"`
	Gw           struct {
		Visa            string `json:"visa"`
		Master          string `json:"master"`
		Amex            string `json:"amex"`
		Othercards      string `json:"othercards"`
		Internetbanking string `json:"internetbanking"`
		Mobilebanking   string `json:"mobilebanking"`
	} `json:"gw"`
	RedirectGatewayURL       string `json:"redirectGatewayURL"`
	DirectPaymentURLBank     string `json:"directPaymentURLBank"`
	DirectPaymentURLCard     string `json:"directPaymentURLCard"`
	DirectPaymentURL         string `json:"directPaymentURL"`
	RedirectGatewayURLFailed string `json:"redirectGatewayURLFailed"`
	GatewayPageURL           string `json:"GatewayPageURL"`
	StoreBanner              string `json:"storeBanner"`
	StoreLogo                string `json:"storeLogo"`
	Desc                     []struct {
		Name string `json:"name"`
		Type string `json:"type"`
		Logo string `json:"logo"`
		Gw   string `json:"gw"`
	} `json:"desc"`
	IsDirectPayEnable string `json:"is_direct_pay_enable"`
}
