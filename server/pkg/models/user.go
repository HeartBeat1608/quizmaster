package models

import "time"

type User struct {
  ID int32 `json:"id" db:"id"`
  Email string `json:"email" db:"email"`
  Password string `json:"password" db:"password"`
  Active bool `json:"active" db:"active"`
  CreatedAt time.Time `json:"created_at" db:"created_at"`
}
