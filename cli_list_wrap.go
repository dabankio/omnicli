package omnicli

import (
	"encoding/json"
	"fmt"
	"github.com/lemon-sunxiansong/omnicli/btcjson"
	"os/exec"
	"strconv"
)

// CliListunspent https://bitcoin.org/en/developer-reference#listunspent
func CliListunspent(minconf, maxconf int, addresses []string) ([]btcjson.ListUnspentResult, error) {
	// if includeUnsafe == nil {
	// 	includeUnsafe = btcjson.Bool(false)
	// }
	args := basicParamsWith(
		"listunspent",
		strconv.Itoa(minconf),
		strconv.Itoa(maxconf),
		ToJson(addresses),
		// strconv.FormatBool(*includeUnsafe),
	)
	// if query != nil {
	// 	args = append(args, ToJson(query))
	// }
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, args...,
	))
	var unspents []btcjson.ListUnspentResult
	err := json.Unmarshal([]byte(cmdPrint), &unspents)
	if err != nil {
		err = fmt.Errorf("Decode json err, %v,\n%s", err, cmdPrint)
	}
	return unspents, err
}
