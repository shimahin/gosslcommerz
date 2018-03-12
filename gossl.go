package gosslcommerz

import (
	"github.com/shimahin/gosslcommerz/models"
	"net/http"
)

type PaymentService interface {

	// Create session
	CreateSession() (*models.SessionResponse, error)

	// Set up IPN Listener
	IPNListener(request *http.Request) (map[string]interface{}, error)

	// 	Order Validation
	CheckValidation(valID []string, storeID string, storePass string, v string, format string) (*models.OrderValidationResponse, error)

	// Transaction
	TransactionQuery(transactionID []string, storeID string, storePass string) (*models.TransactionQueryResponse, error)
}
