package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	Id int64 `json:"id"`
	jwt.RegisteredClaims
}

func BuildCustomClaims(id int64, duration, issuer, subject string) (*CustomClaims, error) {
	parseDuration, err := time.ParseDuration(duration)
	if err != nil {
		return nil, err
	}
	return &CustomClaims{
		id, jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(parseDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    issuer,
			Subject:   subject,
		},
	}, nil
}

type JWT struct {
	SigningKey []byte
}

func NewJWT(key string) *JWT {
	return &JWT{SigningKey: []byte(key)}
}

func (j *JWT) CreateToken(claims *CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, *claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
