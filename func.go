package btccli

import (
	"encoding/json"
	"fmt"
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

func toJson(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}

func ToJson(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}

func dividePrint(msg string) {
	fmt.Printf("\n--------------%s--------------\n", msg)
}
