package btccli

import (
	"fmt"
	"github.com/lemon-sunxiansong/btccli/btcjson"
	"os/exec"
	"strconv"
	"strings"
)

// CliImportprivkey https://bitcoin.org/en/developer-reference#importprivkey
func CliImportprivkey(cmd btcjson.ImportPrivKeyCmd) error {
	args := []string{
		CmdParamRegtest, "importprivkey", cmd.PrivKey,
	}
	if cmd.Label != nil {
		args = append(args, *cmd.Label)
	} else {
		args = append(args, "")
	}

	if cmd.Rescan != nil {
		args = append(args, strconv.FormatBool(*cmd.Rescan))
	}
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, args...,
	))
	if strings.Contains(cmdPrint, "error") {
		return fmt.Errorf("import privkey return error: %s", cmdPrint)
	}
	return nil
}

// CliImportpubkey https://bitcoin.org/en/developer-reference#importpubkey
func CliImportpubkey(cmd btcjson.ImportPubKeyCmd) error {
	args := []string{
		CmdParamRegtest, "importpubkey", cmd.PubKey,
	}
	if cmd.Label != nil {
		args = append(args, *cmd.Label)
	} else {
		args = append(args, "")
	}

	if cmd.Rescan != nil {
		args = append(args, strconv.FormatBool(*cmd.Rescan))
	}
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, args...,
	))
	if strings.Contains(cmdPrint, "error") {
		return fmt.Errorf("import privkey return error: %s", cmdPrint)
	}
	return nil

}
