package btccli

func isCoinbaseTx(tx interface{}) bool {
	return false
	// var flag bool
	// for _, dtl := range tx.Details {
	// 	if dtl.Category == "immature" || dtl.Category == "generate" {
	// 		flag = true
	// 		break
	// 	}
	// }
	// return len(tx.Details) == 0 && flag
}
