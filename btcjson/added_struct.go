package btcjson

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
