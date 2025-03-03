package request

type CustomerRequest struct {
	CostumerId   string `json:"costumer_id"`
	CustomerPAN  string `json:"customer_pan"`
	CustomerName string `json:"customer_name"`
	PetugasRekam string `json:"petugas_rekam"`
}
