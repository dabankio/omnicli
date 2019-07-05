package testtool

import (
	"fmt"
	"runtime/debug"
	"testing"
)

// SimpleAssert .
type SimpleAssert struct {
	t *testing.T
}

// NewSimpleAssert .
func NewSimpleAssert(t *testing.T) *SimpleAssert {
	return &SimpleAssert{t: t}
}

// FailOnErr used in testing assert
func (asrt *SimpleAssert) FailOnErr(e error, msg string) {
	if e != nil {
		fmt.Printf("[Fail] on error, %s, %v\n", msg, e)
		debug.PrintStack()
		asrt.t.FailNow()
	}
}

// FailOnFlag false时t.Fatal
func (asrt *SimpleAssert) FailOnFlag(flag bool, params ...interface{}) {
	if flag {
		fmt.Printf("[Fail] on flag, %v\n", params)
		debug.PrintStack()
		asrt.t.FailNow()
	}
}
