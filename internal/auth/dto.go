package auth

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	Type        string `json:"type"`
	ExpiresIn   int64  `json:"expires_in"`
}
