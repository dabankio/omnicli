package btccli

import (
	"testing"
)

func trueThenFailNow(t *testing.T, flag bool, msg string, others ...interface{}) {
	if flag {
		t.Fatalf("[TRUE] %s, %v", msg, others)
	}
}
