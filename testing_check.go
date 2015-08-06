package ext

import (
	"testing"
)

func CheckEqual(t *testing.T, l interface{}, r interface{}) {

	if l == r {
		return
	}

	t.Fatalf("%v != %v", l, r)
}
