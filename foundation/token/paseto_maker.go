package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

// PasetoMaker is a PASETO token maker
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewPasetoMaker creates a new PasetoMaker
func NewPasetoMaker(symmetricKey string) (*PasetoMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *PasetoMaker) CreateToken(username string, Uuid string) (string, error) {
	now := time.Now()
	exp := now.Add(24 * time.Hour)
	nbt := now

	jsonToken := paseto.JSONToken{
		Audience:   "patreons",
		Issuer:     "farmstyle",
		Jti:        Uuid,
		Subject:    "test_subject",
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  nbt,
	}

	jsonToken.Set(username, "Users name")
	footer := "some footer"

	token, err := maker.paseto.Encrypt(maker.symmetricKey, jsonToken, footer)
	return token, err
}

// VerifyToken checks if the token is valid or not
func (maker *PasetoMaker) VerifyToken(token string) (paseto.JSONToken, string, error) {
	var newjsonToken paseto.JSONToken
	var newFooter string

	err := maker.paseto.Decrypt(token, maker.symmetricKey, &newjsonToken, &newFooter)

	return newjsonToken, newFooter, err
}
