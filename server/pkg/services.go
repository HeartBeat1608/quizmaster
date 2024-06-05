package pkg

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
  JWT_KEY []byte = []byte("super_secret_jwt_key")
  JWT_ISSUER string = "quizmaster"
)

func WriteJsonResponse(w http.ResponseWriter, statusCode int, value any) error {
  b, err := json.Marshal(value)
  if err != nil {
    return err
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(statusCode)
  _, err = w.Write(b)
  return err
}

func WriteError(w http.ResponseWriter, statusCode int, err error) error {
  val := make(map[string]string)
  val["error"] = err.Error()
  return WriteJsonResponse(w, statusCode, val)
}

func GenerateToken(claims map[string]interface{}) (string, error) {
  claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
  claims["iss"] = JWT_ISSUER

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
  return token.SignedString(JWT_KEY)
}
