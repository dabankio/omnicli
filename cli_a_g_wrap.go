package btccli

import (
	"encoding/json"
	"github.com/lemon-sunxiansong/btccli/btcjson"
	"os/exec"
	"strconv"
)

// CliCreatemultisig https://bitcoin.org/en/developer-reference#createmultisig
func CliCreatemultisig(nRequired uint8, keys []string, addressType *string) (btcjson.CreateMultiSigResult, error) {
	args := []string{
		CmdParamRegtest, "createmultisig", strconv.Itoa(int(nRequired)), toJson(keys),
	}
	if addressType != nil {
		args = append(args, *addressType)
	}
	cmdPrint := cmdAndPrint(exec.Command(CmdBitcoinCli, args...))
	//TODO validate address
	var resp btcjson.CreateMultiSigResult
	err := json.Unmarshal([]byte(cmdPrint), &resp)
	return resp, err
}

// CliGeneratetoaddress https://bitcoin.org/en/developer-reference#generatetoaddress
func CliGeneratetoaddress(nBlocks uint, address string, maxtriesPtr *uint) ([]string, error) {
	maxtries := 1000000
	if maxtriesPtr != nil {
		maxtries = int(*maxtriesPtr)
	}
	cmd := exec.Command(CmdBitcoinCli, CmdParamRegtest, "generatetoaddress", strconv.Itoa(int(nBlocks)), address, strconv.Itoa(maxtries))
	cmdPrint := cmdAndPrint(cmd)
	var hashs []string
	err := json.Unmarshal([]byte(cmdPrint), &hashs)
	return hashs, err
}

// CliDumpprivkey https://bitcoin.org/en/developer-reference#dumpprivkey
func CliDumpprivkey(addr string) (string, error) {
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, CmdParamRegtest, "dumpprivkey", addr,
	))
	//TODO validate privKey
	return cmdPrint, nil
}
