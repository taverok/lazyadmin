package jwtutils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/taverok/lazyadmin/pkg/rest"
)

type JwtClaims struct {
	Uuid  string `json:"uuid,omitempty"`
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
	jwt.RegisteredClaims
}

func (it JwtClaims) GenerateJwtToken(secret string) (string, error) {
	claims := &JwtClaims{
		Uuid:  it.Uuid,
		Email: it.Email,
		Role:  it.Role,
		//RegisteredClaims: jwt.RegisteredClaims{
		//	ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(d)},
		//},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetClaims(token *string, secret string) (*JwtClaims, error) {
	if token == nil {
		return nil, rest.ErrNotAuthorized
	}

	claims, err := tokenClaims(*token, secret)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func tokenClaims(strToken string, secret string) (*JwtClaims, error) {
	key := []byte(secret)

	claims := &JwtClaims{}

	token, err := jwt.ParseWithClaims(strToken, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return claims, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func ParseAuthToken(r *http.Request) *string {
	authHeader := r.Header.Get("Authorization")
	bearerToken := strings.Split(authHeader, " ")

	if len(bearerToken) < 2 {
		return nil
	}

	return &bearerToken[1]
}
