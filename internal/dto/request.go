package dto

type RegisterRequest struct {
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	Email      string `json:"email"`
	Passport   string `json:"passport"`
	Inn        string `json:"inn"`
	Snils      string `json:"snils"`
	Birthday   string `json:"birthday"`
	Role       string `json:"role"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateAccessTokenRequest struct {
	//Id         int    `json:"id"`
	//LastName   string `json:"last_name"`
	//FirstName  string `json:"first_name"`
	//MiddleName string `json:"middle_name"`
	//Email      string `json:"email"`
	//Role       string `json:"role"`
	//AccessTokenString  string `json:"access_token_string"`
	//RefreshTokenString string `json:"refresh_token_string"`
}
