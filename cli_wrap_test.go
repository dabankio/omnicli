package omnicli

import (
	"fmt"
	"github.com/lomocoin/omnicli/testtool"
	"testing"

	"github.com/lomocoin/omnicli/btcjson"
)

func TestCliCreatemultisig(t *testing.T) {
	cli, killomnicored, err := RunOmnicored(&RunOptions{NewTmpDir: true})
	testtool.FailOnFlag(t, err != nil, "Failed to start btcd", err)
	defer killomnicored()

	type addrinfo struct {
		addr, privkey, pubkey string
	}
	var addrs [3]addrinfo
	{ //获取几个新地址
		for i := 0; i < len(addrs); i++ {
			add, err := cli.Getnewaddress(nil, nil)
			testtool.FailOnFlag(t, err != nil, "Failed to get new address", err)
			addrs[i] = addrinfo{addr: add}

			validateResp, err := cli.Validateaddress(add)
			testtool.FailOnErr(t, err, "failed to validate address")
			addrs[i].pubkey = validateResp.Pubkey

			privkey, err := cli.Dumpprivkey(add)
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
		multisigResp, err = cli.Createmultisig(2, keys, nil)
		testtool.FailOnFlag(t, err != nil, "Failed to create multi sig", err)
		fmt.Println("keys", keys)
		fmt.Println("multisig address:", jsonStr(multisigResp))
	}

	{
		vRes, err := cli.Validateaddress(multisigResp.Address)
		testtool.FailOnFlag(t, err != nil, "Failed to validate address info", err)
		fmt.Println("validate multisig address", ToJsonIndent(vRes))
	}

}

func TestCliAddmultisigaddress(t *testing.T) {
	cli, killomnicored, err := RunOmnicored(&RunOptions{NewTmpDir: true})
	testtool.FailOnFlag(t, err != nil, "Failed to start btcd", err)
	defer killomnicored()

	type addrinfo struct {
		addr, privkey, pubkey string
	}
	var addrs [3]addrinfo
	{ //获取几个新地址
		for i := 0; i < len(addrs); i++ {
			add, err := cli.Getnewaddress(nil, nil)
			testtool.FailOnFlag(t, err != nil, "Failed to get new address", err)
			addrs[i] = addrinfo{addr: add}

			validateResp, err := cli.Validateaddress(add)
			testtool.FailOnErr(t, err, "failed to validate address")
			addrs[i].pubkey = validateResp.Pubkey

			privkey, err := cli.Dumpprivkey(add)
			testtool.FailOnFlag(t, err != nil, "Failed to dump privkey", err)
			addrs[i].privkey = privkey
		}
		fmt.Println("addrs", addrs)
	}

	{ //create multisig address
		var keys []string
		for _, info := range addrs {
			keys = append(keys, info.pubkey)
		}

		multisigResp, err := cli.Addmultisigaddress(btcjson.AddMultisigAddressCmd{
			NRequired: 2, Keys: keys,
		})
		testtool.FailOnErr(t, err, "Failed to add multi sig address")
		fmt.Println("multisig address:", multisigResp)
	}

}

func TestCliGetAddressInfo(t *testing.T) {
	cli, killomnicored, err := RunOmnicored(&RunOptions{NewTmpDir: true})
	testtool.FailOnFlag(t, err != nil, "Failed to start btcd", err)
	defer killomnicored()

	var newAddr string
	{
		newAddr, err = cli.Getnewaddress(nil, nil)
		testtool.FailOnFlag(t, err != nil, "Failed to get new address", err)
	}
	{
		vRes, err := cli.Validateaddress(newAddr)
		testtool.FailOnFlag(t, err != nil, "Failed to validate address", err)
		fmt.Println("validate address res:", ToJsonIndent(vRes))
	}
}

func TestCliGetWalletInfo(t *testing.T) {
	cli, killomnicored, err := RunOmnicored(&RunOptions{NewTmpDir: true})
	testtool.FailOnFlag(t, err != nil, "Failed to start btcd", err)
	defer killomnicored()

	info, err := cli.GetWalletInfo()
	testtool.FailOnErr(t, err, "get wallet info")
	fmt.Println(ToJsonIndent(info))
}

func TestCliListunspent(t *testing.T) {
	cli, killomnicored, err := RunOmnicored(&RunOptions{NewTmpDir: true})
	testtool.FailOnFlag(t, err != nil, "Failed to start d", err)
	defer killomnicored()

	var newaddr string
	{
		newaddr, err = cli.Getnewaddress(nil, nil)
		testtool.FailOnFlag(t, err != nil, "Failed to get new address", err)
	}
	{
		const leng = 102
		hashs, err := cli.Generatetoaddress(leng, newaddr, nil)
		testtool.FailOnFlag(t, err != nil, "Failed to gen to addr", err)
		testtool.FailOnFlag(t, len(hashs) != leng, "len not equal", leng, hashs)
	}
	{
		unspents, err := cli.Listunspent(0, 999, []string{newaddr})
		testtool.FailOnFlag(t, err != nil, "Fail on listunspent", err)
		fmt.Println(jsonStr(unspents))
	}
}
