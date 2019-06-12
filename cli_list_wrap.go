package btccli

import (
	"encoding/json"
	"github.com/lemon-sunxiansong/btccli/btcjson"
	"os/exec"
	"strconv"
)

// CliListunspent https://bitcoin.org/en/developer-reference#listunspent
func CliListunspent(minconf, maxconf int, addresses []string, includeUnsafe *bool, query *btcjson.ListUnspentQueryOptions) ([]btcjson.ListUnspentResult, error) {
	if includeUnsafe == nil {
		includeUnsafe = btcjson.Bool(false)
	}
	args := []string{
		CmdParamRegtest,
		"listunspent",
		strconv.Itoa(minconf),
		strconv.Itoa(maxconf),
		ToJson(addresses),
		strconv.FormatBool(*includeUnsafe),
	}
	if query != nil {
		args = append(args, ToJson(query))
	}
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, args...,
	))
	var unspents []btcjson.ListUnspentResult
	err := json.Unmarshal([]byte(cmdPrint), &unspents)
	return unspents, err
}
