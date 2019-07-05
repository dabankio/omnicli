package btccli

import (
	"github.com/lemon-sunxiansong/btccli/testtool"
	"fmt"
	"github.com/lemon-sunxiansong/btccli/btcjson"
	"testing"
)

func TestCliListunspent(t *testing.T) {
	closeChan, err := BitcoindRegtest()
	testtool.FailOnFlag(t, err != nil, "Failed to start d", err)
	defer func() {
		closeChan <- struct{}{}
	}()

	var newaddr string
	{
		newaddr, err = CliGetnewaddress(nil, nil)
		testtool.FailOnFlag(t, err != nil, "Failed to get new address", err)
	}
	{
		const leng = 102
		hashs, err := CliGeneratetoaddress(leng, newaddr, nil)
		testtool.FailOnFlag(t, err != nil, "Failed to gen to addr", err)
		testtool.FailOnFlag(t, len(hashs) != leng, "len not equal", leng, hashs)
	}
	{
		unspents, err := CliListunspent(0, 999, []string{newaddr}, btcjson.Bool(true), nil)
		testtool.FailOnFlag(t, err != nil, "Fail on listunspent", err)
		fmt.Println(jsonStr(unspents))
	}
}
