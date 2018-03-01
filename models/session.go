package models


type RequestValue struct {
	StoreId          string   `json:"store_id"`
	StorePassword    string   `json:"store_passwd"`
	TotalAmount      string   `json:"total_amount"`
	Currency         string   `json:"currency"`
	TranID           string   `json:"tran_id"`
	SuccessURL       string   `json:"success_url"`
	FailUrl          string   `json:"fail_url"`
	CancelURL        string   `json:"cancel_url"`
	CustomerName     string   `json:"customer_name"`
	CustomerEmail    string   `json:"customer_email"`
	CustomerAdd1     string   `json:"customer_add_1"`
	CustomerAdd2     string   `json:"customer_add_2"`
	CustomerCity     string   `json:"customer_city"`
	CustomerState    string   `json:"customer_state"`
	CustomerPostCode string   `json:"customer_post_code"`
	CustomerCountry  string   `json:"customer_country"`
	CustomerPhone    string   `json:"customer_phone"`
	CustomerFax      string   `json:"customer_fax"`
	ShipName         string   `json:"ship_name"`
	ShipAdd1         string   `json:"ship_add_1"`
	ShipAdd2         string   `json:"ship_add_2"`
	ShipCity         string   `json:"ship_city"`
	ShipState        string   `json:"ship_state"`
	ShipPostCode     string   `json:"ship_post_code"`
	ShipCountry      string   `json:"ship_country"`
	MultiCardName    []string `json:"multi_card_name"`
	ValueA           string   `json:"value_a"`
	ValueB           string   `json:"value_b"`
	ValueC           string   `json:"value_c"`
	ValueD           string   `json:"value_d"`
}
