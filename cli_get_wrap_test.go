package omnicli

import (
	"fmt"
	"github.com/lomocoin/omnicli/testtool"
	"testing"
)

func TestCliGetAddressInfo(t *testing.T) {
	cc, err := StartOmnicored()
	testtool.FailOnFlag(t, err != nil, "Failed to start btcd", err)
	defer func() {
		cc <- struct{}{}
	}()

	var newAddr string
	{
		newAddr, err = CliGetnewaddress(nil, nil)
		testtool.FailOnFlag(t, err != nil, "Failed to get new address", err)
	}
	{
		vRes, err := CliValidateaddress(newAddr)
		testtool.FailOnFlag(t, err != nil, "Failed to validate address", err)
		fmt.Println("validate address res:", ToJsonIndent(vRes))
	}
}

func TestCliGetWalletInfo(t *testing.T) {
	cc, err := StartOmnicored()
	testtool.FailOnFlag(t, err != nil, "Failed to start btcd", err)
	defer func() {
		cc <- struct{}{}
	}()

	info, err := CliGetWalletInfo()
	testtool.FailOnErr(t, err, "get wallet info")
	fmt.Println(ToJsonIndent(info))
}
