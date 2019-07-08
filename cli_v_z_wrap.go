package btccli

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

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

// CliValidateaddress .
func CliValidateaddress(addr string) (ValidateAddressResp, error) {
	validateCmd := exec.Command(CmdBitcoinCli, basicParamsWith("validateaddress", addr)...)
	cmdPrint := cmdAndPrint(validateCmd) //auto print result
	var resp ValidateAddressResp
	err := json.Unmarshal([]byte(cmdPrint), &resp)
	if err != nil {
		err = fmt.Errorf("Failed to decode validate address resp,(%s), err: %v", cmdPrint, err)
	}
	return resp, err
}
