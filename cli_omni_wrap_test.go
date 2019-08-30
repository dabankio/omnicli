package omnicli

import (
	"github.com/lomocoin/omnicli/testtool"
	"testing"
)

func TestCliOmniCreatepaloadSimplesend(t *testing.T) {
	cli, killomnicored, err := RunOmnicored(&RunOptions{NewTmpDir: true})
	testtool.FailOnFlag(t, err != nil, "Failed to start d", err)
	defer killomnicored()

	ret, err := cli.OmniCreatepaloadSimplesend(2, "0.1")
	testtool.FailOnErr(t, err, "")
	expectedRet := "00000000000000020000000000989680"
	testtool.FailOnFlag(t, ret != expectedRet, "not as expected", ret, expectedRet)
}
