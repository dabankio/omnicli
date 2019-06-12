package btccli

import (
	"fmt"
	"testing"
)

func TestCliGetAddressInfo(t *testing.T) {
	cc, err := BitcoindRegtest()
	trueThenFailNow(t, err != nil, "Failed to start btcd", err)
	defer func() {
		cc <- struct{}{}
	}()

	var newAddr string
	{
		newAddr, err = CliGetnewaddress(nil, nil)
		trueThenFailNow(t, err != nil, "Failed to get new address", err)
	}
	{
		addrInfo, err := CliGetAddressInfo(newAddr)
		trueThenFailNow(t, err != nil, "Failed to get address info", err)
		fmt.Println("address info", ToJsonIndent(addrInfo))
	}
	{
		vRes, err := CliValidateaddress(newAddr)
		trueThenFailNow(t, err != nil, "Failed to validate address", err)
		fmt.Println("validate address res:", ToJsonIndent(vRes))
	}
}
