package btccli

import (
	"fmt"
	"github.com/lemon-sunxiansong/btccli/btcjson"
	"os/exec"
)

type Addr struct {
	Address string
	Privkey string
	Pubkey  string
}

func (ad *Addr) String() string {
	return fmt.Sprintf("{Address: \"%s\", Privkey: \"%s\", Pubkey: \"%s\"}", ad.Address, ad.Privkey, ad.Pubkey)
}

// CliToolGetSomeAddrs 一次获取n个地址（包含pub-priv key)
func CliToolGetSomeAddrs(n int) ([]Addr, error) {
	var addrs []Addr
	for i := 0; i < n; i++ {
		add, err := CliGetnewaddress(nil, nil)
		if err != nil {
			return nil, err
		}

		vr, err := CliValidateaddress(add)
		if err != nil {
			return nil, err
		}
		_ = vr

		dump, err := CliDumpprivkey(add)
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, Addr{
			Address: add, Privkey: dump, Pubkey: vr.Pubkey,
		})
	}
	return addrs, nil
}

func cliResult(method string, args ...string) string {
	withMethod := append([]string{method}, args...)
	return cmdAndPrint(exec.Command(
		CmdBitcoinCli, basicParamsWith(withMethod...)...,
	))
}

// DecodeAndPrintTX panic on error
func DecodeAndPrintTX(title, rawtx string) {
	PrintCmdOut = false
	defer func() {
		PrintCmdOut = true
	}()
	tx, err := CliDecoderawtransaction(btcjson.DecodeRawTransactionCmd{
		HexTx: rawtx,
	})
	if err != nil {
		panic(fmt.Errorf("failed to decode rawtx, %v", err))
	}
	fmt.Printf("----[TX]%s------\n%s\n%s\n", title, rawtx, ToJsonIndent(tx))
}
