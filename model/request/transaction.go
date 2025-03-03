package request

type Transaksi struct {
	RequestId           string `json:"request_id"`
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
	PetugasRekam        string `json:"petugas_rekam"`
}
type TransaksiStatus struct {
	RequestId  string `json:"request_id"`
	BillNumber string `json:"bill_number"`
}
