package helper

import (
	"crypto/rand"
	"fmt"
)

func GenerateIdempotencyKey() string {
	b := make([]byte, 10)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
