package btccliwrap

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

func dividePrint(msg string) {
	fmt.Printf("\n--------------%s--------------\n", msg)
}
