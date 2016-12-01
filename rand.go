package ext

import (
	"encoding/hex"
	"math/rand"
)

func RandomString(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err == nil {
		return hex.EncodeToString(b)
	}
	return ""
}
