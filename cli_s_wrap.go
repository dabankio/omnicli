package btccli

import (
	"github.com/lemon-sunxiansong/btccli/btcjson"
	"os/exec"
	"strconv"
)

// CliSendtoaddress https://bitcoin.org/en/developer-reference#sendtoaddress
func CliSendtoaddress(cmd *btcjson.SendToAddressCmd) (string, error) {
	args := []string{
		CmdParamRegtest,
		"sendtoaddress",
		cmd.Address,
		strconv.FormatFloat(cmd.Amount, 'f', 6, 64),
	}
	if cmd.Comment != nil {
		args = append(args, *cmd.Comment)
	} else {
		args = append(args, "")
	}

	if cmd.CommentTo != nil {
		args = append(args, *cmd.CommentTo)
	} else {
		args = append(args, "")
	}
	//TODO support other params
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, args...,
	))
	//TODO validate hex
	return cmdPrint, nil
}
