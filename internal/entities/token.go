package entities

type Token struct {
	TokenString string `json:"token_string"`
	ExpiresAt   int64  `json:"expires_at"`
	IssuedAt    int64  `json:"issued_at"`
}
