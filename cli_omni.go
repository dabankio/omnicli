package omnicli

import (
	"encoding/json"
	"os/exec"
	"strconv"

	"github.com/lomocoin/omnicli/btcjson"
)

// CliOmniCreaterawtxChange https://github.com/OmniLayer/omnicore/blob/master/src/omnicore/doc/rpc-api.md#omni_createrawtx_change
func CliOmniCreaterawtxChange(rawtx string, prevtxs []btcjson.PreviousDependentTxOutput, destination string, fee float64, position *int) (string, error) {
	args := basicParamsWith(
		"omni_createrawtx_change",
		rawtx,
		ToJson(prevtxs),
		destination,
		strconv.FormatFloat(fee, 'f', -1, 32),
	)
	if position != nil {
		args = append(args, strconv.Itoa(*position))
	}

	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, args...,
	))
	return cmdPrint, ToError(cmdPrint)
}

// CliOmniCreaterawtxOpreturn Adds a payload with class C (op-return) encoding to the transaction.
// If no raw transaction is provided, a new transaction is created.
// If the data encoding fails, then the transaction is not modified.
func CliOmniCreaterawtxOpreturn(rawtx string, payload string) (string, error) {
	args := basicParamsWith(
		"omni_createrawtx_opreturn", rawtx, payload,
	)
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, args...,
	))
	return cmdPrint, ToError(cmdPrint)
}

// CliOmniCreaterawtxReference Adds a reference output to the transaction.
// If no raw transaction is provided, a new transaction is created.
// The output value is set to at least the dust threshold.
func CliOmniCreaterawtxReference(rawtx, destination string, amount *int) (string, error) {
	args := basicParamsWith(
		"omni_createrawtx_reference", rawtx, destination,
	)
	if amount != nil {
		args = append(args, strconv.Itoa(*amount))
	}
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, args...,
	))
	return cmdPrint, ToError(cmdPrint)
}

// CliOmniCreatepaloadSimplesend .
func CliOmniCreatepaloadSimplesend(propertyID int, amount string) (string, error) {
	args := basicParamsWith(
		"omni_createpayload_simplesend",
		strconv.Itoa(propertyID),
		amount,
	)

	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, args...,
	))
	return cmdPrint, ToError(cmdPrint)
}

// OmniGetbalanceResult .
type OmniGetbalanceResult struct {
	Balance  string `json:"balance"`
	Reserved string `json:"reserved"`
	Frozen   string `json:"frozen"`
}

// CliOmniGetbalance https://github.com/OmniLayer/omnicore/blob/master/src/omnicore/doc/rpc-api.md#omni_getbalance
func CliOmniGetbalance(address string, propertyid int) (*OmniGetbalanceResult, error) {
	cmdPrint := cliResult("omni_getbalance", address, strconv.Itoa(propertyid))
	var ret OmniGetbalanceResult
	err := json.Unmarshal([]byte(cmdPrint), &ret)
	return &ret, WrapJSONDecodeError(err, cmdPrint)
}

// OmniGettransactionResult .
type OmniGettransactionResult struct {
	Txid             string `json:"txid"`             // (string) the hex-encoded hash of the transaction
	Sendingaddress   string `json:"sendingaddress"`   // (string) the Bitcoin address of the sender
	Referenceaddress string `json:"referenceaddress"` // (string) a Bitcoin address used as reference (if any)
	Ismine           bool   `json:"ismine"`           // (boolean) whether the transaction involes an address in the wallet
	Confirmations    int    `json:"confirmations"`    // (number) the number of transaction confirmations
	Fee              string `json:"fee"`              // (string) the transaction fee in bitcoins
	Blocktime        int    `json:"blocktime"`        // (number) the timestamp of the block that contains the transaction
	Valid            bool   `json:"valid"`            // (boolean) whether the transaction is valid
	Positioninblock  int    `json:"positioninblock"`  // (number) the position (index) of the transaction within the block
	Version          int    `json:"version"`          // (number) the transaction version
	TypeInt          int    `json:"type_int"`         // (number) the transaction type as number
	Type             string `json:"type"`             // (string) the transaction type as string
	//other
	Propertyid int `json:"propertyid"`
}

// CliOmniGettransaction .
func CliOmniGettransaction(txHash string) (*OmniGettransactionResult, error) {
	cmdPrint := cliResult("omni_gettransaction", txHash)
	var ret OmniGettransactionResult
	err := json.Unmarshal([]byte(cmdPrint), &ret)
	return &ret, WrapJSONDecodeError(err, cmdPrint)
}

// OmniSenddissuancefixedCmd .
type OmniSenddissuancefixedCmd struct {
	Fromaddress                                    string
	Ecosystem                                      int
	Typ                                            int
	Previousid                                     int
	Category, Subcategory, Name, URL, Data, Amount string
}

// CliOmniSendissuancefixed https://github.com/OmniLayer/omnicore/blob/master/src/omnicore/doc/rpc-api.md#omni_sendissuancefixed
func CliOmniSendissuancefixed(cmd *OmniSenddissuancefixedCmd) (string, error) {
	cmdPrint := cliResult(
		"omni_sendissuancefixed",
		cmd.Fromaddress, strconv.Itoa(cmd.Ecosystem), strconv.Itoa(cmd.Typ), strconv.Itoa(cmd.Previousid),
		cmd.Category, cmd.Subcategory, cmd.Name, cmd.URL, cmd.Data, cmd.Amount,
	)
	return cmdPrint, ToError(cmdPrint)
}
