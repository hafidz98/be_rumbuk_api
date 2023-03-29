package service

type AuthLoginRequest struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
