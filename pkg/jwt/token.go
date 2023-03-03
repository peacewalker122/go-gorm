package jwt

import (
	"github.com/golang-jwt/jwt"
)

type JwtStruct struct {
	SecretKey []byte
}

func NewJwt(secretKey string) Jwt {
	return &JwtStruct{
		SecretKey: []byte(secretKey),
	}
}

func (j *JwtStruct) GenerateToken(payload *Payload) (string, error) {

	newJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := newJwt.SignedString(j.SecretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (j *JwtStruct) ValidateToken(token string) (*Payload, error) {

	payload := &Payload{}

	_, err := jwt.ParseWithClaims(token, payload, func(token *jwt.Token) (interface{}, error) {
		return j.SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return payload, nil
}
