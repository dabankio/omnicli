package omnicli

import (
	"encoding/json"
	"github.com/lemon-sunxiansong/omnicli/btcjson"
	"os/exec"
	"strconv"
)

func CliAddmultisigaddress(cmd btcjson.AddMultisigAddressCmd) (btcjson.CreateMultiSigResult, error) {
	args := []string{
		CmdParamRegtest, "addmultisigaddress", strconv.Itoa(int(cmd.NRequired)), ToJson(cmd.Keys),
	}
	if cmd.Label != nil {
		args = append(args, *cmd.Label)

		if cmd.AddressType != nil {
			args = append(args, *cmd.AddressType)
		}
	}

	cmdPrint := cmdAndPrint(exec.Command(CmdBitcoinCli, args...))
	//TODO validate address
	var resp btcjson.CreateMultiSigResult
	err := json.Unmarshal([]byte(cmdPrint), &resp)
	return resp, err
}

// CliCreatemultisig https://bitcoin.org/en/developer-reference#createmultisig
func CliCreatemultisig(nRequired uint8, keys []string, addressType *string) (btcjson.CreateMultiSigResult, error) {
	args := []string{
		CmdParamRegtest, "createmultisig", strconv.Itoa(int(nRequired)), ToJson(keys),
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

// CliCreaterawtransaction https://bitcoin.org/en/developer-reference#createrawtransaction
func CliCreaterawtransaction(cmd btcjson.CreateRawTransactionCmd) (string, error) {
	args := basicParamsWith(
		"createrawtransaction",
		ToJson(cmd.Inputs),
		IfOrString(len(cmd.Outputs) > 0, ToJson(cmd.Outputs), "{}"),
	)
	if cmd.LockTime != nil {
		args = append(args, strconv.Itoa(int(*cmd.LockTime)))
	}
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, args...,
	))
	//TODO validate hex
	return cmdPrint, ToError(cmdPrint)
}

// CliDecoderawtransaction https://bitcoin.org/en/developer-reference#decoderawtransaction
func CliDecoderawtransaction(cmd btcjson.DecodeRawTransactionCmd) (*btcjson.DecodeRawTransactionResult, error) {
	args := basicParamsWith(
		"decoderawtransaction",
		cmd.HexTx,
	)
	if cmd.Iswitness != nil {
		args = append(args, strconv.FormatBool(*cmd.Iswitness))
	}
	cmdPrint := cmdAndPrint(exec.Command(CmdBitcoinCli, args...))
	var res btcjson.DecodeRawTransactionResult
	err := json.Unmarshal([]byte(cmdPrint), &res)
	return &res, WrapJSONDecodeError(err, cmdPrint)
}

// CliDecodescript https://bitcoin.org/en/developer-reference#decodescript
func CliDecodescript(hex string) (btcjson.DecodeScriptResult, error) {
	args := basicParamsWith("decodescript", hex)
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, args...,
	))
	var res btcjson.DecodeScriptResult
	err := json.Unmarshal([]byte(cmdPrint), &res)
	return res, err
}

// CliDumpprivkey https://bitcoin.org/en/developer-reference#dumpprivkey
func CliDumpprivkey(addr string) (string, error) {
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, basicParamsWith("dumpprivkey", addr)...,
	))
	//TODO validate privKey
	return cmdPrint, nil
}

// CliGeneratetoaddress https://bitcoin.org/en/developer-reference#generatetoaddress
func CliGeneratetoaddress(nBlocks uint, address string, maxtriesPtr *uint) ([]string, error) {
	maxtries := 1000000
	if maxtriesPtr != nil {
		maxtries = int(*maxtriesPtr)
	}
	cmd := exec.Command(CmdBitcoinCli, basicParamsWith("generatetoaddress", strconv.Itoa(int(nBlocks)), address, strconv.Itoa(maxtries))...)
	cmdPrint := cmdAndPrint(cmd)
	var hashs []string
	err := json.Unmarshal([]byte(cmdPrint), &hashs)
	return hashs, err
}

// CliGenerate https://bitcoin.org/en/developer-reference#generatetoaddress
// func CliGenerate(nBlocks uint, maxtriesPtr *uint) ([]string, error) {
// 	maxtries := 1000000
// 	if maxtriesPtr != nil {
// 		maxtries = int(*maxtriesPtr)
// 	}
// 	cmd := exec.Command(CmdBitcoinCli, CmdParamRegtest, "generatetoaddress", strconv.Itoa(int(nBlocks)), address, strconv.Itoa(maxtries))
// 	cmdPrint := cmdAndPrint(cmd)
// 	var hashs []string
// 	err := json.Unmarshal([]byte(cmdPrint), &hashs)
// 	return hashs, err
// }
