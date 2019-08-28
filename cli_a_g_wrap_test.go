package omnicli

import (
	"fmt"
	"github.com/lomocoin/omnicli/testtool"
	"testing"

	"github.com/lomocoin/omnicli/btcjson"
)

func TestCliCreatemultisig(t *testing.T) {
	closeChan, err := StartOmnicored()
	testtool.FailOnFlag(t, err != nil, "Failed to start btcd", err)
	defer func() {
		closeChan <- struct{}{}
	}()

	type addrinfo struct {
		addr, privkey, pubkey string
	}
	var addrs [3]addrinfo
	{ //获取几个新地址
		for i := 0; i < len(addrs); i++ {
			add, err := CliGetnewaddress(nil, nil)
			testtool.FailOnFlag(t, err != nil, "Failed to get new address", err)
			addrs[i] = addrinfo{addr: add}

			privkey, err := CliDumpprivkey(add)
			testtool.FailOnFlag(t, err != nil, "Failed to dump privkey", err)
			addrs[i].privkey = privkey
		}
		fmt.Println("addrs", addrs)
	}

	var multisigResp btcjson.CreateMultiSigResult
	{ //create multisig address
		var keys []string
		for _, info := range addrs {
			keys = append(keys, info.pubkey)
			// keys = append(keys, info.addr)
		}
		multisigResp, err = CliCreatemultisig(2, keys, nil)
		testtool.FailOnFlag(t, err != nil, "Failed to create multi sig", err)
		fmt.Println("keys", keys)
		fmt.Println("multisig address:", jsonStr(multisigResp))
	}

	{
		vRes, err := CliValidateaddress(multisigResp.Address)
		testtool.FailOnFlag(t, err != nil, "Failed to validate address info", err)
		fmt.Println("validate multisig address", ToJsonIndent(vRes))
	}

}

func TestCliAddmultisigaddress(t *testing.T) {
	closeChan, err := StartOmnicored()
	testtool.FailOnFlag(t, err != nil, "Failed to start btcd", err)
	defer func() {
		closeChan <- struct{}{}
	}()

	type addrinfo struct {
		addr, privkey, pubkey string
	}
	var addrs [5]addrinfo
	{ //获取几个新地址
		for i := 0; i < len(addrs); i++ {
			add, err := CliGetnewaddress(nil, nil)
			testtool.FailOnFlag(t, err != nil, "Failed to get new address", err)
			addrs[i] = addrinfo{addr: add}

			privkey, err := CliDumpprivkey(add)
			testtool.FailOnFlag(t, err != nil, "Failed to dump privkey", err)
			addrs[i].privkey = privkey
		}
		fmt.Println("addrs", addrs)
	}

	var multisigResp btcjson.CreateMultiSigResult
	{ //create multisig address
		var keys []string
		for _, info := range addrs {
			keys = append(keys, info.pubkey)
		}

		multisigResp, err = CliAddmultisigaddress(btcjson.AddMultisigAddressCmd{
			NRequired: 3, Keys: keys,
		})
		testtool.FailOnFlag(t, err != nil, "Failed to add multi sig address", err)
		fmt.Println("multisig address:", jsonStr(multisigResp))
	}

}
