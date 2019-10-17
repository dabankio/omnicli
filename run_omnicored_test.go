package omnicli

import (
	"fmt"
	"testing"

	"github.com/dabankio/omnicli/testtool"
)

func TestStartOmnicored(t *testing.T) {
	cli, killomnicored, err := RunOmnicored(&RunOptions{NewTmpDir: true})
	testtool.FailOnErr(t, err, "omnicore start err")
	defer func() {
		killomnicored()
		testtool.FailOnFlag(t, cmdIsPortContainsNameRunning(RPCPortRegtest, "omnicored"), "omnicored should be stopped")
		t.Log("Done")
	}()

	fmt.Println("================to_get_balance=======")

	ret, err := cli.OmniCreatepaloadSimplesend(2, "0.1")
	testtool.FailOnErr(t, err, "")
	expectedRet := "00000000000000020000000000989680"
	testtool.FailOnFlag(t, ret != expectedRet, "not as expected", ret, expectedRet)

	testtool.FailOnFlag(t, !cmdIsPortContainsNameRunning(RPCPortRegtest, "omnicored"), "端口现在应该已经运行")

	_, _, err = RunOmnicored(nil)
	testtool.FailOnFlag(t, err == nil, "再次运行应该返回错误")
}
