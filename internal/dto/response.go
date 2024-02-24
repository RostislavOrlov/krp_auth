package dto

type AuthResponse struct {
	Email                 string `json:"email"`
	Password              string `json:"password"`
	AccessTokenString     string `json:"access_token_string"`
	AccessTokenExpiresAt  int64  `json:"access_token_expires_at"`
	AccessTokenIssuedAt   int64  `json:"access_token_issued_at"`
	RefreshTokenString    string `json:"refresh_token_string"`
	RefreshTokenExpiresAt int64  `json:"refresh_token_expires_at"`
	RefreshTokenIssuedAt  int64  `json:"refresh_token_issued_at"`
}

type RegisterResponse struct {
	Id         int    `json:"id"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
}

type UpdateAccessTokenResponse struct {
	AccessTokenString  string `json:"access_token_string"`
	RefreshTokenString string `json:"refresh_token_string"`
}
