package btccliwrap

import (
	"strconv"
	"encoding/json"
	"fmt"
	"os/exec"
)

const (
	CmdParamRegtest = "-regtest"
	CliGenerate     = "generate"
)

func cliGetbestblockhash() (string, error) {
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, CmdParamRegtest, "getbestblockhash",
	))
	//TODO validate hash
	return cmdPrint, nil
}

type ValidateAddressResp struct {
	Isvalid      bool   `json:"isvalid,omitempty"`
	Address      string `json:"address,omitempty"`
	ScriptPubKey string `json:"scriptPubKey,omitempty"`
	Isscript     bool   `json:"isscript,omitempty"`
	Iswitness    bool   `json:"iswitness,omitempty"`
}

func cliValidateaddress(addr string) (ValidateAddressResp, error) {
	validateCmd := exec.Command(CmdBitcoinCli, CmdParamRegtest, "validateaddress", addr)
	cmdPrint := cmdAndPrint(validateCmd) //auto print result
	var resp ValidateAddressResp
	err := json.Unmarshal([]byte(cmdPrint), &resp)
	if err != nil {
		err = fmt.Errorf("Failed to decode validate address resp,(%s), err: %v", cmdPrint, err)
	}
	return resp, err
}

type GetAddressInfoResp struct {
	Address             string   `json:"address,omitempty"`
	ScriptPubKey        string   `json:"scriptPubKey,omitempty"`
	Ismine              bool     `json:"ismine,omitempty"`
	Iswatchonly         bool     `json:"iswatchonly,omitempty"`
	Solvable            bool     `json:"solvable,omitempty"`
	Desc                string   `json:"desc,omitempty"`
	Isscript            bool     `json:"isscript,omitempty"`
	Ischange            bool     `json:"ischange,omitempty"`
	Iswitness           bool     `json:"iswitness,omitempty"`
	WitnessProgram      string   `json:"witness_program,omitempty"`
	Script              string   `json:"script,omitempty"`
	Hex                 string   `json:"hex,omitempty"`
	Pubkeys             []string `json:"pubkeys,omitempty"`
	Pubkey              string   `json:"pubkey,omitempty"`
	Iscompressed        bool     `json:"iscompressed,omitempty"`
	Label               string   `json:"label,omitempty"`
	Hdkeypath           string   `json:"hdkeypath,omitempty"`
	Hdseedid            string   `json:"hdseedid,omitempty"`
	Hdmasterfingerprint string   `json:"hdmasterfingerprint,omitempty"`
	// "witness_version" : version   (numeric, optional) The version number of the witness program
	// "sigsrequired" : xxxxx        (numeric, optional) Number of signatures required to spend multisig output (only if "script" is "multisig")
	// "embedded" : {...},           (object, optional) Information about the address embedded in P2SH or P2WSH, if relevant and known. It includes all getaddressinfo output fields for the embedded address, excluding metadata ("timestamp", "hdkeypath", "hdseedid") and relation to the wallet ("ismine", "iswatchonly").
	// "timestamp" : timestamp,      (number, optional) The creation time of the key if available in seconds since epoch (Jan 1 1970 GMT)
	// "labels"                      (object) Array of labels associated with the address.
	//   [
	// 	{ (json object of label data)
	// 	  name string
	// 	  purpose string
	// 	},...
	//   ]
}

func cliGetAddressInfo(addr string) (*GetAddressInfoResp, error) {
	cmd := exec.Command(CmdBitcoinCli, CmdParamRegtest, "getaddressinfo", addr)
	cmdPrint := cmdAndPrint(cmd)
	var resp GetAddressInfoResp
	err := json.Unmarshal([]byte(cmdPrint), &resp)
	return &resp, err
}

func cliGetWalletInfo() map[string]interface{} {
	cmd := exec.Command(CmdBitcoinCli, CmdParamRegtest, "getwalletinfo")
	cmdPrint := cmdAndPrint(cmd)
	var info map[string]interface{}
	json.Unmarshal([]byte(cmdPrint), &info)
	return info
}

func cliGetblockcount() (int, error) {
	cmd := exec.Command(CmdBitcoinCli, CmdParamRegtest)
	cmdPrint := cmdAndPrint(cmd)
	return strconv.Atoi(cmdPrint)
}

func cliGetblockhash(height int) (string, error) {
	cmdPrint := cmdAndPrint(exec.Command(CmdBitcoinCli, CmdParamRegtest, "getblockhash", strconv.Itoa(height)))
	//TODO validate hash
	return cmdPrint, nil
}


type GetblockResp struct {
	HexHash string
	
}
func cliGetBlock(hash string, verbosity int) (map[string]interface{}, error) {
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, CmdParamRegtest,
		"getblock",
		strconv.Itoa(verbosity),
	))
	var res map[string]interface{}
	err := json.Unmarshal([]byte(cmdPrint), &res)
	return res, err
}