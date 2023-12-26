package hash

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(s *string) error {
	if s == nil {
		return errors.New("Reference provided for hashing is nil")
	}

	sBytes := []byte(*s)

	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	*s = string(hashedBytes[:])
	return nil
}

func CheckPassword(existingHash, incomingPass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(existingHash),
		[]byte(incomingPass)) == nil
}
