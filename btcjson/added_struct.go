package btcjson

type GetAddressInfoResp struct {
	Address             string                   `json:"address"`
	ScriptPubKey        string                   `json:"scriptPubKey"`
	Ismine              bool                     `json:"ismine"`
	Iswatchonly         bool                     `json:"iswatchonly"`
	Solvable            bool                     `json:"solvable"`
	Desc                string                   `json:"desc"`
	Isscript            bool                     `json:"isscript"`
	Ischange            bool                     `json:"ischange"`
	Iswitness           bool                     `json:"iswitness"`
	WitnessProgram      string                   `json:"witness_program"`
	Script              string                   `json:"script"`
	Hex                 string                   `json:"hex"`
	Pubkeys             []string                 `json:"pubkeys"`
	Pubkey              string                   `json:"pubkey"`
	Iscompressed        bool                     `json:"iscompressed"`
	Label               string                   `json:"label"`
	Hdkeypath           string                   `json:"hdkeypath"`
	Hdseedid            string                   `json:"hdseedid"`
	Hdmasterfingerprint string                   `json:"hdmasterfingerprint"`
	WitnessVersion      int                      `json:"witness_version"` // : version   (numeric, optional) The version number of the witness program
	Sigsrequired        uint                     `json:"sigsrequired"`    //        (numeric, optional) Number of signatures required to spend multisig output (only if "script" is "multisig")
	Embedded            map[string]interface{}   `json:"embedded"`        //,           (object, optional) Information about the address embedded in P2SH or P2WSH, if relevant and known. It includes all getaddressinfo output fields for the embedded address, excluding metadata ("timestamp", "hdkeypath", "hdseedid") and relation to the wallet ("ismine", "iswatchonly").
	Timestamp           uint64                   `json:"timestamp"`       // : timestamp,      (number, optional) The creation time of the key if available in seconds since epoch (Jan 1 1970 GMT)
	Labels              []map[string]interface{} `json:"labels"`          //                      (object) Array of labels associated with the address.
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

type DecodeRawTransactionResult struct {
	TxID     string `json:"txid"`     //        (string) The transaction id
	Hash     string `json:"hash"`     //        (string) The transaction hash (differs from txid for witness transactions)
	Size     int    `json:"size"`     //             (numeric) The transaction size
	Vsize    int    `json:"vsize"`    //            (numeric) The virtual transaction size (differs from size for witness transactions)
	Weight   int    `json:"weight"`   //           (numeric) The transaction's weight (between vsize*4 - 3 and vsize*4)
	Version  int    `json:"version"`  //          (numeric) The version
	Locktime int    `json:"locktime"` //       (numeric) The lock time
	Vin      []Vin  `json:"vin"`
	Vout     []Vout `json:"vout"`
}
