package dto

type LoginRequest struct {
  Email string `json:"email"`
  Password string `json:"password"`
}

type RegisterRequest struct {
  Email string `json:"email" db:"email"`
  Password string `json:"password" db:"password"`
}
