package jwt

type Jwt interface {
	GenerateToken(payload *Payload) (string, error)
	ValidateToken(token string) (*Payload, error)
}
