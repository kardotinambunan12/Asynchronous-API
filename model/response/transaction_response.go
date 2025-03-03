package response

type DataCustomerResponse struct {
	CostumerId   string `json:"costumer_id"`
	CustomerPAN  string `json:"customer_pan"`
	CustomerName string `json:"customer_name"`
	TglRekam     string `json:"tgl_rekam"`
	PetugasRekam string `json:"petugas_rekam"`
}

type DataMerchantResponse struct {
	MerchantID   string `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	MerchantCity string `json:"merchant_city"`
	PetugasRekam string `json:"petugas_rekam"`
	TglRekam     string `json:"tgl_rekam"`
}

type TransactionStatusResponse struct {
	Code                string `json:"code,omitempty"`
	Message             string `json:"message,omitempty"`
	RequestID           string `json:"request_id"`
	CustomerPAN         string `json:"customer_pan"`
	Amount              string `json:"amount"`
	TransactionDatetime string `json:"transaction_datetime"`
	RRN                 string `json:"rrn"`
	BillNumber          string `json:"bill_number"`
	CustomerName        string `json:"customer_name"`
	MerchantID          string `json:"merchant_id"`
	MerchantName        string `json:"merchant_name"`
	MerchantCity        string `json:"merchant_city"`
	CurrencyCode        string `json:"currency_code"`
	PaymentStatus       string `json:"payment_status"`
	PaymentDescription  string `json:"payment_description"`
}

type GenerateCodeResponse struct {
	BillNumber string `json:"bill_number"`
}
