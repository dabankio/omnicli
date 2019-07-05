package btccli

import (
	"github.com/lemon-sunxiansong/btccli/testtool"
	"fmt"
	"testing"
)

func TestCliGetAddressInfo(t *testing.T) {
	cc, err := BitcoindRegtest()
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
		addrInfo, err := CliGetAddressInfo(newAddr)
		testtool.FailOnFlag(t, err != nil, "Failed to get address info", err)
		fmt.Println("address info", ToJsonIndent(addrInfo))
	}
	{
		vRes, err := CliValidateaddress(newAddr)
		testtool.FailOnFlag(t, err != nil, "Failed to validate address", err)
		fmt.Println("validate address res:", ToJsonIndent(vRes))
	}
}
