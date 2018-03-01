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

type IpnResponse struct {
	VerifyKey             []string `json:"verify_key"`
	CardIssuer            string   `json:"card_issuer"`
	TranDate              string   `json:"tran_date"`
	ValueA                string   `json:"value_a"`
	ValueB                string   `json:"value_b"`
	ValueC                string   `json:"value_c"`
	ValueD                string   `json:"value_d"`
	RiskTitle             string   `json:"risk_title"`
	Code                  string   `json:"code"`
	CardIssuerCountryCode string   `json:"card_issuer_country_code"`
	CurrencyAmount        string   `json:"currency_amount"`
	RiskLevel             string   `json:"risk_level"`
	StoreAmount           string   `json:"store_amount"`
	ValID                 string   `json:"val_id"`
	Amount                string   `json:"amount"`
	CardBrand             string   `json:"card_brand"`
	VerifySignSha2        string   `json:"verify_sign_sha_2"`
	CrrencyRate           string   `json:"crrency_rate"`
	VerifySign            string   `json:"verify_sign"`
	StoreId               string   `json:"store_id"`
	BaseFair              string   `json:"base_fair"`
	Currency              string   `json:"currency"`
	CardIssuerCountry     string   `json:"card_issuer_country"`
	CardNo                string   `json:"card_no"`
	TranId                string   `json:"tran_id"`
	CurrencyType          string   `json:"currency_type"`
	BankTranId            string   `json:"bank_tran_id"`
}
