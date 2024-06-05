package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"quizmaster-server/pkg"
	"quizmaster-server/pkg/database"
	"quizmaster-server/pkg/dto"
	"quizmaster-server/pkg/models"
	"quizmaster-server/pkg/utils"
)

var (
  ErrInvalidCredentials = errors.New("invalid email or password")
  ErrEmailAlreadyExists = errors.New("email already registered")
  ErrFailedToGenerateToken = errors.New("failed to generate token")
)

type AuthHandler struct{
  db *database.AppDatabase
}

func NewAuthHandler(db *database.AppDatabase) *AuthHandler {
	return &AuthHandler{db}
}

func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
  var loginRequest dto.LoginRequest

  if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		pkg.WriteError(w, 500, err)
    return
	}

  var user models.User
  err := h.db.QueryRowx("SELECT * FROM users WHERE email=$1", loginRequest.Email).StructScan(&user)
  if err != nil {
    pkg.WriteError(w, 500, ErrInvalidCredentials)
    return
  }

  if !utils.CheckPassword(loginRequest.Password, user.Password) {
    pkg.WriteError(w, 500, ErrInvalidCredentials)
    return
  }

  userMap := make(map[string]interface{})
  err = h.db.QueryRowx("SELECT id, email FROM users WHERE email=$1", loginRequest.Email).MapScan(userMap)
  if err != nil {
		pkg.WriteError(w, 400, err)
    return
  }

  token, err := pkg.GenerateToken(userMap)
  if err != nil {
		pkg.WriteError(w, 400, ErrFailedToGenerateToken)
    return
  }

  pkg.WriteJsonResponse(w, 200, map[string]interface{} {
		"message": "Login successful",
    "token": token,
    "user": userMap,
  })
}

func (h *AuthHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
  var registerRequest dto.RegisterRequest

  if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		pkg.WriteError(w, 500, err)
    return
	}

  existingId := -1
  err := h.db.QueryRowx("SELECT id from users WHERE email=$1", registerRequest.Email).Scan(&existingId)

  log.Printf("err: %v | existingId: %d", err, existingId)

  if existingId > -1 {
		pkg.WriteError(w, 400, ErrEmailAlreadyExists)
    return
  }

  hashedPw, err := utils.HashPassword(registerRequest.Password)
  if err != nil {
		pkg.WriteError(w, 400, ErrEmailAlreadyExists)
    return
  }

  registerRequest.Password = hashedPw

  query := `
  INSERT INTO users (email, password, created_at, active)
  VALUES (:email, :password, NOW(), true)
  `

  _, err = h.db.NamedExec(query, registerRequest)
  if err != nil {
		pkg.WriteError(w, 400, err)
    return
  }

  userMap := make(map[string]interface{})
  err = h.db.QueryRowx("SELECT id, email FROM users WHERE email=$1", registerRequest.Email).MapScan(userMap)
  if err != nil {
		pkg.WriteError(w, 400, err)
    return
  }

  token, err := pkg.GenerateToken(userMap)
  if err != nil {
		pkg.WriteError(w, 400, ErrFailedToGenerateToken)
    return
  }

  pkg.WriteJsonResponse(w, 200, map[string]interface{} {
    "message": "Registered Successfully",
    "token": token,
    "user": userMap,
  })
}
