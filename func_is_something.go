package omnicli

import (
	"github.com/lemon-sunxiansong/omnicli/btcjson"
)

func isCoinbaseTx(tx *btcjson.GetTransactionResult) bool {
	flag := false
	for _, dtl := range tx.Details {
		if dtl.Category == "immature" || dtl.Category == "generate" {
			flag = true
			break
		}
	}
	return len(tx.Details) == 0 && flag
}
