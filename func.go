package omnicli

import (
	"encoding/json"
	"fmt"
	"github.com/lomocoin/omnicli/btcjson"
	"strings"
)

func panicIf(e error, msg string) {
	if e != nil {
		panic(fmt.Errorf("【ERR】 %s %v", msg, e))
	}
}

func jsonStr(i interface{}) string {
	b, _ := json.MarshalIndent(i, "", " ")
	return string(b)
}

func ToJsonIndent(i interface{}) string {
	b, _ := json.MarshalIndent(i, "", " ")
	return string(b)
}

func ToJson(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}

// IfOrString if flag return s ,or s2
func IfOrString(flag bool, trueS, falseS string) string {
	if flag {
		return trueS
	}
	return falseS
}

// ToError 如果含有error字样，则视为error
func ToError(str string) error {
	if strings.Contains(str, "error") {
		return fmt.Errorf("%s", str)
	}
	return nil
}

// WrapJSONDecodeError include raw json in error
func WrapJSONDecodeError(e error, rawJSON string) error {
	if e == nil {
		return nil
	}
	return fmt.Errorf("Decode json error: %v, json:\n%s", e, rawJSON)
}

func dividePrint(msg string) {
	fmt.Printf("\n--------------%s--------------\n", msg)
}

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
