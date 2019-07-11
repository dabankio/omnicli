package omnicli

import (
	"fmt"
	"github.com/lemon-sunxiansong/omnicli/btcjson"
	"os/exec"
	"strings"
)

// CliImportaddress .
func CliImportaddress(cmd btcjson.ImportAddressCmd) error {
	args := []string{
		CmdParamRegtest,
		"importaddress",
		cmd.Address,
		"",     //TODO process label
		"true", //TODO process rescan
		//TODO other options
	}
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, args...,
	))
	if strings.Contains(cmdPrint, "error") {
		return fmt.Errorf("Not null resp: %s", cmdPrint)
	}
	return nil
}
