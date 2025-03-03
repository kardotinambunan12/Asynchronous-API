package model

type GeneralResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type WebResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
