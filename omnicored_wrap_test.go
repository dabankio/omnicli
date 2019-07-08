package btccli

import (
	"fmt"
	"testing"

	"github.com/lemon-sunxiansong/btccli/testtool"
)

func TestStartOmnicored(t *testing.T) {
	closeChan, err := StartOmnicored()
	testtool.FailOnErr(t, err, "bitcoind start err")
	defer func() {
		closeChan <- struct{}{}
		testtool.FailOnFlag(t, cmdIsPortContainsNameRunning(RPCPortRegtest, "bitcoin"), "bitcoind should be stoped")
		fmt.Println("Done")
	}()

	testtool.FailOnFlag(t, !cmdIsPortContainsNameRunning(RPCPortRegtest, "bitcoin"), "端口现在应该已经运行")

	_, err = StartOmnicored()
	testtool.FailOnFlag(t, err == nil, "再次运行应该返回错误")
}
