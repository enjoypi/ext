package ext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkRandomUint64(b *testing.B) {
	m := make(map[uint64]bool)
	for i := 0; i < b.N; i++ {
		m[RandomUint64()] = true
	}
	AssertB(b, b.N == len(m))
}

func BenchmarkRandomInt64(b *testing.B) {
	m := make(map[int64]bool)
	for i := 0; i < b.N; i++ {
		m[RandomInt64()] = true
	}
	AssertB(b, b.N == len(m))
}

func TestEncrypt(t *testing.T) {
	key := []byte("1234567890123456")
	plain := []byte("abcdefghijklmn")

	cipher, err := CBCEncrypt(key, plain)
	assert.NoError(t, err)

	p, err := CBCDecrypt(key, cipher)
	assert.NoError(t, err)

	assert.EqualValues(t, plain, p)
}
