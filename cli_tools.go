package omnicli

import (
	"fmt"
	"github.com/lomocoin/omnicli/btcjson"
	"os/exec"
)

// Addr .
type Addr struct {
	Address string
	Privkey string
	Pubkey  string
}

func (ad *Addr) String() string {
	return fmt.Sprintf("{Address: \"%s\", Privkey: \"%s\", Pubkey: \"%s\"}", ad.Address, ad.Privkey, ad.Pubkey)
}

// ToolGetSomeAddrs 一次获取n个地址（包含pub-priv key)
func (cli *Cli) ToolGetSomeAddrs(n int) ([]Addr, error) {
	var addrs []Addr
	for i := 0; i < n; i++ {
		add, err := cli.Getnewaddress(nil, nil)
		if err != nil {
			return nil, err
		}

		vr, err := cli.Validateaddress(add)
		if err != nil {
			return nil, err
		}
		_ = vr

		dump, err := cli.Dumpprivkey(add)
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, Addr{
			Address: add, Privkey: dump, Pubkey: vr.Pubkey,
		})
	}
	return addrs, nil
}

func (cli *Cli) cliResult(method string, args ...string) string {
	withMethod := append([]string{method}, args...)
	return cmdAndPrint(exec.Command(
		CmdOmniCli, cli.AppendArgs(withMethod...)...,
	))
}

// DecodeAndPrintTX panic on error
func (cli *Cli) DecodeAndPrintTX(title, rawtx string) {
	PrintCmdOut = false
	defer func() {
		PrintCmdOut = true
	}()
	tx, err := cli.Decoderawtransaction(btcjson.DecodeRawTransactionCmd{
		HexTx: rawtx,
	})
	if err != nil {
		panic(fmt.Errorf("failed to decode rawtx, %v", err))
	}
	fmt.Printf("----[TX]%s------\n%s\n%s\n", title, rawtx, ToJsonIndent(tx))
}
