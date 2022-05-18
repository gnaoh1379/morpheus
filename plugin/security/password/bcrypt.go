package password

import (
	"golang.org/x/crypto/bcrypt"
)

type BCrypt struct {
	cost int
}

func BCryptEncoder(cost int) Encoder {
	e := &BCrypt{cost: bcrypt.DefaultCost}
	if cost >= bcrypt.MinCost || cost <= bcrypt.MaxCost {
		e.cost = cost
	}
	return e
}

func (e *BCrypt) Encode(raw string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(raw), e.cost)
	if err != nil {
		return "", err
	}
	return string(password), nil
}

func (e *BCrypt) Compare(raw string, encoded string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encoded), []byte(raw)) == nil
}
