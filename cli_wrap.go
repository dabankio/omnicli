package omnicli

import (
	"encoding/json"
	"fmt"
	"github.com/lomocoin/omnicli/btcjson"
	"os/exec"
	"strconv"
	"strings"
)

// Addmultisigaddress .
func (cli *Cli) Addmultisigaddress(cmd btcjson.AddMultisigAddressCmd) (string, error) {
	args := cli.AppendArgs("addmultisigaddress", strconv.Itoa(int(cmd.NRequired)), ToJson(cmd.Keys))
	if cmd.Label != nil {
		args = append(args, *cmd.Label)
		if cmd.AddressType != nil {
			args = append(args, *cmd.AddressType)
		}
	}

	cmdPrint := cmdAndPrint(exec.Command(CmdOmniCli, args...))
	if strings.Contains(cmdPrint, "error") {
		return "", fmt.Errorf("respone contains error: %s", cmdPrint)
	}
	return cmdPrint, nil
}

// Createmultisig https://bitcoin.org/en/developer-reference#createmultisig
func (cli *Cli) Createmultisig(nRequired uint8, keys []string, addressType *string) (btcjson.CreateMultiSigResult, error) {
	args := cli.AppendArgs("createmultisig", strconv.Itoa(int(nRequired)), ToJson(keys))
	if addressType != nil {
		args = append(args, *addressType)
	}
	cmdPrint := cmdAndPrint(exec.Command(CmdOmniCli, args...))
	//TODO validate address
	var resp btcjson.CreateMultiSigResult
	err := json.Unmarshal([]byte(cmdPrint), &resp)
	return resp, err
}

// Createrawtransaction https://bitcoin.org/en/developer-reference#createrawtransaction
func (cli *Cli) Createrawtransaction(cmd btcjson.CreateRawTransactionCmd) (string, error) {
	args := cli.AppendArgs(
		"createrawtransaction",
		ToJson(cmd.Inputs),
		IfOrString(len(cmd.Outputs) > 0, ToJson(cmd.Outputs), "{}"),
	)
	if cmd.LockTime != nil {
		args = append(args, strconv.Itoa(int(*cmd.LockTime)))
	}
	cmdPrint := cmdAndPrint(exec.Command(
		CmdOmniCli, args...,
	))
	//TODO validate hex
	return cmdPrint, ToError(cmdPrint)
}

// Decoderawtransaction https://bitcoin.org/en/developer-reference#decoderawtransaction
func (cli *Cli) Decoderawtransaction(cmd btcjson.DecodeRawTransactionCmd) (*btcjson.DecodeRawTransactionResult, error) {
	args := cli.AppendArgs(
		"decoderawtransaction",
		cmd.HexTx,
	)
	if cmd.Iswitness != nil {
		args = append(args, strconv.FormatBool(*cmd.Iswitness))
	}
	cmdPrint := cmdAndPrint(exec.Command(CmdOmniCli, args...))
	var res btcjson.DecodeRawTransactionResult
	err := json.Unmarshal([]byte(cmdPrint), &res)
	return &res, WrapJSONDecodeError(err, cmdPrint)
}

// Decodescript https://bitcoin.org/en/developer-reference#decodescript
func (cli *Cli) Decodescript(hex string) (btcjson.DecodeScriptResult, error) {
	args := cli.AppendArgs("decodescript", hex)
	cmdPrint := cmdAndPrint(exec.Command(
		CmdOmniCli, args...,
	))
	var res btcjson.DecodeScriptResult
	err := json.Unmarshal([]byte(cmdPrint), &res)
	return res, err
}

// Dumpprivkey https://bitcoin.org/en/developer-reference#dumpprivkey
func (cli *Cli) Dumpprivkey(addr string) (string, error) {
	cmdPrint := cmdAndPrint(exec.Command(
		CmdOmniCli, cli.AppendArgs("dumpprivkey", addr)...,
	))
	//TODO validate privKey
	return cmdPrint, ToError(cmdPrint)
}

// Generatetoaddress https://bitcoin.org/en/developer-reference#generatetoaddress
func (cli *Cli) Generatetoaddress(nBlocks uint, address string, maxtriesPtr *uint) ([]string, error) {
	maxtries := 1000000
	if maxtriesPtr != nil {
		maxtries = int(*maxtriesPtr)
	}
	cmd := exec.Command(CmdOmniCli, cli.AppendArgs("generatetoaddress", strconv.Itoa(int(nBlocks)), address, strconv.Itoa(maxtries))...)
	cmdPrint := cmdAndPrint(cmd)
	var hashs []string
	err := json.Unmarshal([]byte(cmdPrint), &hashs)
	return hashs, err
}

// Generate https://bitcoin.org/en/developer-reference#generatetoaddress
// func (cli *Cli) Generate(nBlocks uint, maxtriesPtr *uint) ([]string, error) {
// 	maxtries := 1000000
// 	if maxtriesPtr != nil {
// 		maxtries = int(*maxtriesPtr)
// 	}
// 	cmd := exec.Command(CmdOmniCli, CmdParamRegtest, "generatetoaddress", strconv.Itoa(int(nBlocks)), address, strconv.Itoa(maxtries))
// 	cmdPrint := cmdAndPrint(cmd)
// 	var hashs []string
// 	err := json.Unmarshal([]byte(cmdPrint), &hashs)
// 	return hashs, err
// }

// Getbestblockhash .
func (cli *Cli) Getbestblockhash() (string, error) {
	cmdPrint := cmdAndPrint(exec.Command(
		CmdOmniCli, CmdParamRegtest, "getbestblockhash",
	))
	//TODO validate hash
	return cmdPrint, nil
}

//  support
// func (cli *Cli) GetAddressInfo(addr string) (*btcjson.GetAddressInfoResp, error) {
// 	cmdPrint := cmdAndPrint(exec.Command(
// 		CmdOmniCli, cli.AppendArgs("getaddressinfo", addr,)...
// 	))
// 	var resp btcjson.GetAddressInfoResp
// 	err := json.Unmarshal([]byte(cmdPrint), &resp)
// 	if err != nil {
// 		err = fmt.Errorf("json encode err, %v, print: \n %s", err, cmdPrint)
// 	}
// 	return &resp, err
// }

// GetWalletInfo .
func (cli *Cli) GetWalletInfo() (map[string]interface{}, error) {
	cmdPrint := cmdAndPrint(exec.Command(
		CmdOmniCli, cli.AppendArgs("getwalletinfo")...,
	))
	var info map[string]interface{}
	err := json.Unmarshal([]byte(cmdPrint), &info)
	return info, err
}

// Getblockcount .
func (cli *Cli) Getblockcount() (int, error) {
	cmd := exec.Command(CmdOmniCli, CmdParamRegtest, "getblockcount")
	cmdPrint := cmdAndPrint(cmd)
	cmdPrint = strings.TrimSpace(cmdPrint)
	return strconv.Atoi(cmdPrint)
}

// Getblockhash .
func (cli *Cli) Getblockhash(height int) (string, error) {
	cmdPrint := cmdAndPrint(exec.Command(CmdOmniCli, CmdParamRegtest, "getblockhash", strconv.Itoa(height)))
	//TODO validate hash
	return strings.TrimSpace(cmdPrint), nil
}

// Getblock https://bitcoin.org/en/developer-reference#getblock
func (cli *Cli) Getblock(hash string, verbosity int) (*string, *btcjson.GetBlockResultV1, *btcjson.GetBlockResultV2, error) {
	cmdPrint := cmdAndPrint(exec.Command(
		CmdOmniCli, CmdParamRegtest,
		"getblock",
		hash,
		strconv.Itoa(verbosity),
	))
	var (
		hex string
		b   btcjson.GetBlockResultV1
		b2  btcjson.GetBlockResultV2
		err error
	)
	switch verbosity {
	case 0:
		hex = cmdPrint
	case 1:
		err = json.Unmarshal([]byte(cmdPrint), &b)
	case 2:
		err = json.Unmarshal([]byte(cmdPrint), &b2)
	default:
		err = fmt.Errorf("verbosity must one of 0/1/2, got: %d", verbosity)
	}
	return &hex, &b, &b2, err
}

// Getnewaddress https://bitcoin.org/en/developer-reference#getnewaddress
func (cli *Cli) Getnewaddress(labelPtr, addressTypePtr *string) (hexedAddress string, err error) {
	label := ""
	if labelPtr != nil {
		label = *labelPtr
	}
	args := cli.AppendArgs("getnewaddress", label)
	if addressTypePtr != nil {
		args = append(args, *addressTypePtr)
	}
	cmdPrint := cmdAndPrint(exec.Command(CmdOmniCli, args...))
	//TODO validate address
	return cmdPrint, ToError(cmdPrint)
}

// Gettransaction https://bitcoin.org/en/developer-reference#gettransaction
func (cli *Cli) Gettransaction(txid string, includeWatchonly bool) (*btcjson.GetTransactionResult, error) {
	cmdPrint := cmdAndPrint(exec.Command(
		CmdOmniCli, CmdParamRegtest, "gettransaction", txid, strconv.FormatBool(includeWatchonly),
	))
	var tx btcjson.GetTransactionResult
	err := json.Unmarshal([]byte(cmdPrint), &tx)
	return &tx, err
}

// Getrawtransaction .
func (cli *Cli) Getrawtransaction(cmd btcjson.GetRawTransactionCmd) (*btcjson.RawTx, error) {
	args := []string{ //TODO verbose and blockhash process
		CmdParamRegtest,
		"getrawtransaction",
		cmd.Txid,
		strconv.FormatBool(true),
	}
	cmdPrint := cmdAndPrint(exec.Command(
		CmdOmniCli, args...,
	))
	var tx btcjson.RawTx
	err := json.Unmarshal([]byte(cmdPrint), &tx)
	return &tx, err
}

// Getreceivedbyaddress https://bitcoin.org/en/developer-reference#getreceivedbyaddress
func (cli *Cli) Getreceivedbyaddress(addr string, minconf int) (string, error) {
	args := cli.AppendArgs(
		"getreceivedbyaddress",
		addr,
		strconv.Itoa(minconf),
	)
	cmdPrint := cmdAndPrint(exec.Command(
		CmdOmniCli, args...,
	))
	return cmdPrint, ToError(cmdPrint)
}

// Importprivkey https://bitcoin.org/en/developer-reference#importprivkey
func (cli *Cli) Importprivkey(cmd btcjson.ImportPrivKeyCmd) error {
	args := cli.AppendArgs("importprivkey", cmd.PrivKey)
	if cmd.Label != nil {
		args = append(args, *cmd.Label)
	} else {
		args = append(args, "")
	}

	if cmd.Rescan != nil {
		args = append(args, strconv.FormatBool(*cmd.Rescan))
	}
	cmdPrint := cmdAndPrint(exec.Command(
		CmdOmniCli, args...,
	))
	return ToError(cmdPrint)
}

// Importpubkey https://bitcoin.org/en/developer-reference#importpubkey
func (cli *Cli) Importpubkey(cmd btcjson.ImportPubKeyCmd) error {
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
		CmdOmniCli, args...,
	))
	if strings.Contains(cmdPrint, "error") {
		return fmt.Errorf("import privkey return error: %s", cmdPrint)
	}
	return nil

}

// Importaddress .
func (cli *Cli) Importaddress(cmd btcjson.ImportAddressCmd) error {
	args := cli.AppendArgs(
		"importaddress",
		cmd.Address,
		"",     //TODO process label
		"true", //TODO process rescan
		//TODO other options
	)
	cmdPrint := cmdAndPrint(exec.Command(
		CmdOmniCli, args...,
	))
	if strings.Contains(cmdPrint, "error") {
		return fmt.Errorf("Not null resp: %s", cmdPrint)
	}
	return ToError(cmdPrint)
}

// Listunspent https://bitcoin.org/en/developer-reference#listunspent
func (cli *Cli) Listunspent(minconf, maxconf int, addresses []string) ([]btcjson.ListUnspentResult, error) {
	// if includeUnsafe == nil {
	// 	includeUnsafe = btcjson.Bool(false)
	// }
	args := cli.AppendArgs(
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
		CmdOmniCli, args...,
	))
	var unspents []btcjson.ListUnspentResult
	err := json.Unmarshal([]byte(cmdPrint), &unspents)
	if err != nil {
		err = fmt.Errorf("Decode json err, %v,\n%s", err, cmdPrint)
	}
	return unspents, err
}

// Sendtoaddress https://bitcoin.org/en/developer-reference#sendtoaddress
func (cli *Cli) Sendtoaddress(cmd *btcjson.SendToAddressCmd) (string, error) {
	args := cli.AppendArgs(
		"sendtoaddress",
		cmd.Address,
		strconv.FormatFloat(cmd.Amount, 'f', 6, 64),
	)
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
		CmdOmniCli, args...,
	))
	//TODO validate hex
	return cmdPrint, ToError(cmdPrint)
}

// Sendrawtransaction https://bitcoin.org/en/developer-reference#sendrawtransaction
func (cli *Cli) Sendrawtransaction(cmd btcjson.SendRawTransactionCmd) (string, error) {
	args := cli.AppendArgs(
		"sendrawtransaction",
		cmd.HexTx,
	)
	if cmd.AllowHighFees != nil {
		args = append(args, strconv.FormatBool(*cmd.AllowHighFees))
	}
	cmdPrint := cmdAndPrint(exec.Command(
		CmdOmniCli, args...,
	))
	//TODO validate hex
	return cmdPrint, ToError(cmdPrint)
}

// Signrawtransactionwithkey [Disabled in omni current version] https://bitcoin.org/en/developer-reference#signrawtransactionwithkey
func (cli *Cli) Signrawtransactionwithkey(cmd btcjson.SignRawTransactionCmd) (btcjson.SignRawTransactionResult, error) {
	args := cli.AppendArgs(
		"signrawtransactionwithkey",
		cmd.RawTx,
		ToJson(cmd.PrivKeys),
		IfOrString(len(cmd.Prevtxs) > 0, ToJson(cmd.Prevtxs), ""),
	)
	if cmd.Sighashtype != nil {
		args = append(args, *cmd.Sighashtype)
	}
	cmdPrint := cmdAndPrint(exec.Command(
		CmdOmniCli, args...,
	))
	var res btcjson.SignRawTransactionResult
	err := json.Unmarshal([]byte(cmdPrint), &res)
	return res, WrapJSONDecodeError(err, cmdPrint)
}

// Signrawtransaction .
func (cli *Cli) Signrawtransaction(cmd btcjson.SignRawTransactionCmd) (btcjson.SignRawTransactionResult, error) {
	args := cli.AppendArgs(
		"signrawtransaction",
		cmd.RawTx,
		IfOrString(len(cmd.Prevtxs) > 0, ToJson(cmd.Prevtxs), ""),
		ToJson(cmd.PrivKeys),
	)
	if cmd.Sighashtype != nil {
		args = append(args, *cmd.Sighashtype)
	}
	cmdPrint := cmdAndPrint(exec.Command(
		CmdOmniCli, args...,
	))
	var res btcjson.SignRawTransactionResult
	err := json.Unmarshal([]byte(cmdPrint), &res)
	return res, WrapJSONDecodeError(err, cmdPrint)
}

// ValidateAddressResp .
type ValidateAddressResp struct {
	Isvalid      bool   `json:"isvalid"`
	Address      string `json:"address"`
	ScriptPubKey string `json:"scriptPubKey"`
	Isscript     bool   `json:"isscript"`
	Iswitness    bool   `json:"iswitness"`
	Pubkey       string `json:"pubkey"`

	WitnessVersion string `json:"witness_version"` // version   (numeric, optional) The version number of the witness program
	WitnessProgram string `json:"witness_program"` // "hex"     (string, optional) The hex value of the witness program
}

// Validateaddress .
func (cli *Cli) Validateaddress(addr string) (ValidateAddressResp, error) {
	validateCmd := exec.Command(CmdOmniCli, cli.AppendArgs("validateaddress", addr)...)
	cmdPrint := cmdAndPrint(validateCmd) //auto print result
	var resp ValidateAddressResp
	err := json.Unmarshal([]byte(cmdPrint), &resp)
	if err != nil {
		err = fmt.Errorf("Failed to decode validate address resp,(%s), err: %v", cmdPrint, err)
	}
	return resp, err
}
