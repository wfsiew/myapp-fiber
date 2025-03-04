package model

type User struct {
	ID           uint   `json:"id"`
    Email        string `json:"email"`
    PasswordHash string `json:"password_hash,omitempty"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}