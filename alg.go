package ext

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"math/big"
	"os"
)

var (
	maxBigInt = big.NewInt(math.MaxInt64)
)

func RandomBigInt() *big.Int {
	n, err := rand.Int(rand.Reader, maxBigInt)
	ANoError(err)
	return n
}

func RandomUint16() uint16 {
	return uint16(RandomBigInt().Uint64())
}

func RandomUint64() uint64 {
	return RandomBigInt().Uint64()
}

func RandomInt64() int64 {
	return RandomBigInt().Int64()
}

// newUUID generates a random UUID according to RFC 4122
func NewUUID() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	ATrue(n == len(uuid) && err == nil)
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

func DeviceID() string {
	hostname, err := os.Hostname()
	ANoError(err)
	hasher := sha1.New()
	hasher.Write([]byte(hostname))
	return hex.EncodeToString(hasher.Sum(nil))
}
