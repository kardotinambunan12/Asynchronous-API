package request

type Merchant struct {
	MerchantID   string `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	MerchantCity string `json:"merchant_city"`
	PetugasRekam string `json:"petugas_rekam"`
	TglRekam     string `json:"tgl_rekam"`
}
