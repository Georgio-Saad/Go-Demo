package helpers

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateVerificationCode() string {
	var max = big.NewInt(999999)
	n, _ := rand.Int(rand.Reader, max)
	return fmt.Sprintf("%06d", n)
}
