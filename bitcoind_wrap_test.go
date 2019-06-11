package btccli

import (
	"testing"
)

func TestBitcoindRegtest(t *testing.T) {
	closeChan, err := BitcoindRegtest()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		closeChan <- struct{}{}

		if cmdIsPortContainsNameRunning(RPCPortRegtest, "bitcoin") {
			t.Fatal("bitcoind should be stoped")
		}
		t.Log("Done")
	}()

	if !cmdIsPortContainsNameRunning(RPCPortRegtest, "bitcoin") {
		t.Fatal("port not running error")
	}

	_, err = BitcoindRegtest()
	if err == nil {
		t.Fatal("再次运行应该返回错误")
	}
}
