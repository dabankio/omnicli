package btccli

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// ValidateAddressResp .
type ValidateAddressResp struct {
	Isvalid      bool   `json:"isvalid,omitempty"`
	Address      string `json:"address,omitempty"`
	ScriptPubKey string `json:"scriptPubKey,omitempty"`
	Isscript     bool   `json:"isscript,omitempty"`
	Iswitness    bool   `json:"iswitness,omitempty"`
}

// CliValidateaddress .
func CliValidateaddress(addr string) (ValidateAddressResp, error) {
	validateCmd := exec.Command(CmdBitcoinCli, CmdParamRegtest, "validateaddress", addr)
	cmdPrint := cmdAndPrint(validateCmd) //auto print result
	var resp ValidateAddressResp
	err := json.Unmarshal([]byte(cmdPrint), &resp)
	if err != nil {
		err = fmt.Errorf("Failed to decode validate address resp,(%s), err: %v", cmdPrint, err)
	}
	return resp, err
}
