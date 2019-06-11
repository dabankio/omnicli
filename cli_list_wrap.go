package btccli

import (
	"encoding/json"
	"os/exec"
	"strconv"
)

// ListUnspentQueryOptions .
type ListUnspentQueryOptions struct {
	MinimumAmount    float64 `json:"minimumAmount"`    //       (numeric or string, optional, default=0) Minimum value of each UTXO in BTC
	MaximumAmount    float64 `json:"maximumAmount"`    //       (numeric or string, optional, default=unlimited) Maximum value of each UTXO in BTC
	MaximumCount     int     `json:"maximumCount"`     //,             (numeric, optional, default=unlimited) Maximum number of UTXOs
	MinimumSumAmount float64 `json:"minimumSumAmount"` //    (numeric or string, optional, default=unlimited) Minimum sum value of all UTXOs in BTC
}

// Unspent .
type Unspent struct {
	Txid          string  `json:"txid,omitempty"`          //          (string) the transaction id
	Vout          int     `json:"vout"`                    //               (numeric) the vout value
	Address       string  `json:"address,omitempty"`       //    (string) the bitcoin address
	Label         string  `json:"label,omitempty"`         //        (string) The associated label, or "" for the default label
	ScriptPubKey  string  `json:"scriptPubKey,omitempty"`  //   (string) the script key
	Amount        float64 `json:"amount"`                  //xxx,         (numeric) the transaction output amount in BTC
	Confirmations int     `json:"confirmations"`           //      (numeric) The number of confirmations
	RedeemScript  string  `json:"redeemScript,omitempty"`  // (string) The redeemScript if scriptPubKey is P2SH
	WitnessScript string  `json:"witnessScript,omitempty"` // (string) witnessScript if the scriptPubKey is P2WSH or P2SH-P2WSH
	Spendable     bool    `json:"spendable,omitempty"`     //        (bool) Whether we have the private keys to spend this output
	Solvable      bool    `json:"solvable,omitempty"`      //         (bool) Whether we know how to spend this output, ignoring the lack of keys
	Desc          string  `json:"desc,omitempty"`          //             (string, only when solvable) A descriptor for spending this output
	Safe          bool    `json:"safe,omitempty"`          //             (bool) Whether this output is considered safe to spend. Unconfirmed transactions from outside keys and unconfirmed replacement transactions are considered unsafe and are not eligible for spending by fundrawtransaction and sendtoaddress.
}

// CliListunspent https://bitcoin.org/en/developer-reference#listunspent
func CliListunspent(minconf, maxconf int, addresses []string, includeUnsafe bool, query *ListUnspentQueryOptions) ([]Unspent, error) {
	args := []string{
		CmdParamRegtest,
		"listunspent",
		strconv.Itoa(minconf),
		strconv.Itoa(maxconf),
		toJson(addresses),
		strconv.FormatBool(includeUnsafe),
	}
	if query != nil {
		args = append(args, toJson(query))
	}
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, args...,
	))
	var unspents []Unspent
	err := json.Unmarshal([]byte(cmdPrint), &unspents)
	return unspents, err
}
