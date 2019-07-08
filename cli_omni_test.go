package btccli

import (
	"github.com/lemon-sunxiansong/btccli/testtool"
	"testing"
)

func TestCliOmniCreatepaloadSimplesend(t *testing.T) {
	closeChan, err := StartOmnicored()
	testtool.FailOnFlag(t, err != nil, "Failed to start d", err)
	defer func() {
		closeChan <- struct{}{}
	}()

	ret, err := CliOmniCreatepaloadSimplesend(2, "0.1")
	testtool.FailOnErr(t, err, "")
	expectedRet := "00000000000000020000000000989680"
	testtool.FailOnFlag(t, ret != expectedRet, "not as expected", ret, expectedRet)
}
