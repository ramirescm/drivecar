package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ramirescm/drivecar/internal/users"
)

var jwtSecret = "ererfsfs"

type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"user_name"`
	jwt.RegisteredClaims
}

func createToken(user *users.User) (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute)

	claims := &Claims{
		UserID:   user.ID,
		Username: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Auth(rw http.ResponseWriter, r *http.Request) {
	var cred Credentials

	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	u, err := users.Authenticate(cred.Username, cred.Password)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := createToken(u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Write([]byte(token))
}
