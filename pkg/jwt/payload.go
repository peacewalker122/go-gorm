package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Payload struct {
	jwt.StandardClaims
}

func (p *Payload) GetID() uuid.UUID {
	return uuid.Must(uuid.Parse(p.Id))
}
