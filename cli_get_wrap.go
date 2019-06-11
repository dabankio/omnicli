package btccli

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// some cmd consts
const (
	CmdParamRegtest = "-regtest"
)

func CliGetbestblockhash() (string, error) {
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, CmdParamRegtest, "getbestblockhash",
	))
	//TODO validate hash
	return cmdPrint, nil
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

func CliGetAddressInfo(addr string) (*GetAddressInfoResp, error) {
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, CmdParamRegtest, "getaddressinfo", addr,
	))
	var resp GetAddressInfoResp
	err := json.Unmarshal([]byte(cmdPrint), &resp)
	return &resp, err
}

func CliGetWalletInfo() map[string]interface{} {
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, CmdParamRegtest, "getwalletinfo",
	))
	var info map[string]interface{}
	json.Unmarshal([]byte(cmdPrint), &info)
	return info
}

func CliGetblockcount() (int, error) {
	cmd := exec.Command(CmdBitcoinCli, CmdParamRegtest, "getblockcount")
	cmdPrint := cmdAndPrint(cmd)
	cmdPrint = strings.TrimSpace(cmdPrint)
	return strconv.Atoi(cmdPrint)
}

func CliGetblockhash(height int) (string, error) {
	cmdPrint := cmdAndPrint(exec.Command(CmdBitcoinCli, CmdParamRegtest, "getblockhash", strconv.Itoa(height)))
	//TODO validate hash
	return strings.TrimSpace(cmdPrint), nil
}

type GetblockRespBase struct {
	Hash              string  `json:"hash,omitempty"`
	Confirmations     int     `json:"confirmations,omitempty"`
	Size              int     `json:"size"`
	Strippedsize      int     `json:"strippedsize,omitempty"`
	Weight            int     `json:"weight"`
	Height            int     `json:"height"`
	Version           int     `json:"version,omitempty"`
	VersionHex        string  `json:"version_hex,omitempty"`
	Merkleroot        string  `json:"merkleroot,omitempty"`
	Time              uint64  `json:"time,omitempty"`
	Mediantime        uint64  `json:"mediantime,omitempty"`
	Nonce             int     `json:"nonce,omitempty"`
	Bits              string  `json:"bits,omitempty"`
	Difficulty        float64 `json:"difficulty,omitempty"`
	Chainwork         string  `json:"chainwork,omitempty"`
	NTx               int     `json:"n_tx,omitempty"`
	Previousblockhash string  `json:"previousblockhash,omitempty"`
	Nextblockhash     string  `json:"nextblockhash,omitempty"`
}
type BlocRespV1 struct {
	GetblockRespBase
	Tx []string `json:"tx,omitempty"`
}
type BlocRespV2 struct {
	GetblockRespBase
	Tx []RawTx `json:"tx,omitempty"`
}
type RawTx struct {
	InActiveChain bool   `json:"in_active_chain,omitempty"`
	Hex           string `json:"hex,omitempty"`
	Txid          string `json:"txid,omitempty"`
	Hash          string `json:"hash,omitempty"`
	Size          uint64 `json:"size,omitempty"`
	Vsize         uint64 `json:"vsize,omitempty"`
	Weight        uint64 `json:"weight,omitempty"`
	Version       uint64 `json:"version,omitempty"`
	Locktime      uint64 `json:"locktime,omitempty"`
	Vin           []Vin  `json:"vin,omitempty"`
	Vout          []Vout `json:"vout,omitempty"`
	Blockhash     string `json:"blockhash,omitempty"`
	Confirmations uint64 `json:"confirmations,omitempty"`
	Blocktime     uint64 `json:"blocktime,omitempty"`
	Time          uint64 `json:"time,omitempty"`
}

type ScriptSig struct {
	Asm string `json:"asm,omitempty"`
	Hex string `json:"hex,omitempty"`
}
type Vin struct {
	Txid        string    `json:"txid,omitempty"`
	Vout        uint64    `json:"vout"`
	ScriptSig   ScriptSig `json:"scriptSig,omitempty"`
	Sequence    uint64    `json:"sequence,omitempty"`
	Txinwitness []string  `json:"txinwitness,omitempty"`
}
type Vout struct {
	Value        float64      `json:"value,omitempty"`
	N            uint64       `json:"n"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey,omitempty"`
}
type ScriptPubKey struct {
	Asm       string   `json:"asm,omitempty"`
	Hex       string   `json:"hex,omitempty"`
	ReqSigs   uint64   `json:"reqSigs,omitempty"`
	Typ       string   `json:"type,omitempty"`
	Addresses []string `json:"addresses,omitempty"`
}

// CliGetblock https://bitcoin.org/en/developer-reference#getblock
func CliGetblock(hash string, verbosity int) (*string, *BlocRespV1, *BlocRespV2, error) {
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, CmdParamRegtest,
		"getblock",
		hash,
		strconv.Itoa(verbosity),
	))
	var (
		hex string
		b   BlocRespV1
		b2  BlocRespV2
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

// CliGetnewaddress https://bitcoin.org/en/developer-reference#getnewaddress
func CliGetnewaddress(labelPtr, addressTypePtr *string) (hexedAddress string, err error) {
	label := ""
	if labelPtr != nil {
		label = *labelPtr
	}
	args := []string{CmdParamRegtest, "getnewaddress", label}
	if addressTypePtr != nil {
		args = append(args, *addressTypePtr)
	}
	cmdPrint := cmdAndPrint(exec.Command(CmdBitcoinCli, args...))
	//TODO validate address
	return cmdPrint, nil
}

type Tx struct {
	Amount            uint64     `json:"amount,omitempty"`              // : x.xxx,        (numeric) The transaction amount in BTC
	Fee               uint64     `json:"fee,omitempty"`                 //: x.xxx,            (numeric) The amount of the fee in BTC. This is negative and only available for the 'send' category of transactions.
	Confirmations     uint64     `json:"confirmations,omitempty"`       // : n,     (numeric) The number of confirmations
	Blockhash         string     `json:"blockhash,omitempty"`           // : "hash",  (string) The block hash
	Blockindex        uint64     `json:"blockindex,omitempty"`          // : xx,       (numeric) The index of the transaction in the block that includes it
	Blocktime         uint64     `json:"blocktime,omitempty"`           // : ttt,       (numeric) The time in seconds since epoch (1 Jan 1970 GMT)
	Txid              string     `json:"txid,omitempty"`                // : "transactionid",   (string) The transaction id.
	Time              uint64     `json:"time,omitempty"`                // : ttt,            (numeric) The transaction time in seconds since epoch (1 Jan 1970 GMT)
	Timereceived      uint64     `json:"timereceived,omitempty"`        // : ttt,    (numeric) The time received in seconds since epoch (1 Jan 1970 GMT)
	Bip125Replaceable string     `json:"bip-125-replaceable,omitempty"` //replaceable": "yes|no|unknown",  (string) Whether this transaction could be replaced due to BIP125 (replace-by-fee); may be unknown for unconfirmed transactions not in the mempool
	Details           []TxDetail `json:"details,omitempty"`
	Hex               string     `json:"hex,omitempty"` // : "data"         (string) Raw data for transaction
}

type TxDetail struct {
	Address  string `json:"address,omitempty"`  // : "address",          (string) The bitcoin address involved in the transaction
	Category string `json:"category,omitempty"` // :                      (string) The transaction category.
	//    "send"                  Transactions sent.
	//    "receive"               Non-coinbase transactions received.
	//    "generate"              Coinbase transactions received with more than 100 confirmations.
	//    "immature"              Coinbase transactions received with 100 or fewer confirmations.
	//    "orphan"                Orphaned coinbase transactions received.
	Amount    uint64 `json:"amount,omitempty"`    // : x.xxx,                 (numeric) The amount in BTC
	Label     string `json:"label,omitempty"`     // : "label",              (string) A comment for the address/transaction, if any
	Vout      uint64 `json:"vout,omitempty"`      // : n,                       (numeric) the vout value
	Fee       uint64 `json:"fee,omitempty"`       //: x.xxx,                     (numeric) The amount of the fee in BTC. This is negative and only available for the 'send' category of transactions.
	Abandoned bool   `json:"abandoned,omitempty"` //                  (bool) 'true' if the transaction has been abandoned (inputs are respendable). Only available for the 'send' category of transactions.
}

// CliGettransaction https://bitcoin.org/en/developer-reference#gettransaction
func CliGettransaction(txid string, includeWatchonly bool) (*Tx, error) {
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, CmdParamRegtest, "gettransaction", txid, strconv.FormatBool(includeWatchonly),
	))
	var tx Tx
	err := json.Unmarshal([]byte(cmdPrint), &tx)
	return &tx, err
}
