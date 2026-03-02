package auth

import (
	"jwt-auth/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret     []byte
	expiration time.Duration
}

func NewJWT(cfg *config.Config) *JWT {
	return &JWT{
		secret:     []byte(cfg.JWTSecret),
		expiration: cfg.JWTExpiration,
	}
}

func (j *JWT) GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(j.expiration).Unix(),
	})

	return token.SignedString(j.secret)
}

func (j *JWT) ValidateToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}
