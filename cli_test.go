package btccli

import (
	"fmt"
	"testing"
)

func TestSimpleTx(t *testing.T) {
	PrintCmdOut = false

	closeChan, err := BitcoindRegtest()
	trueThenFailNow(t, err != nil, "Failed to start bitcoind", err)
	defer func() {
		closeChan <- struct{}{}
		fmt.Println("Done")
	}()

	var (
		newaddr string
	)

	{ // getnewaddress
		newaddr, err = CliGetnewaddress(nil, nil)
		trueThenFailNow(t, err != nil, "Failed to get new address", err)
	}

	{
		leng := 3
		hashs, err := CliGeneratetoaddress(uint(leng), newaddr, nil)
		trueThenFailNow(t, err != nil, "Failed to generate to address", err)
		trueThenFailNow(t, len(hashs) != leng, "Generated hashs len not as expeted", leng, len(hashs))
	}

	scanChain(scanOps{includeGenBlock: true, includeCoinbaseTx: true})
}
