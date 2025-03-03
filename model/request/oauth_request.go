package request

type GoogleTokenInfo struct {
	Audience  string `json:"audience"`
	UserID    string `json:"user_id"`
	ExpiresIn int    `json:"expires_in"`
}
