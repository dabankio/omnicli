package omnicli

import (
	"fmt"
)

type scanOps struct {
	includeGenBlock   bool // 打印 generated block
	includeCoinbaseTx bool // 打印fee o tx
	simpleBlock       bool // 打印block的少数字段
}

func scanAll() {
	scanChain(scanOps{})
}
func scanChain(op scanOps) {
	dividePrint("迭代所有的块")

	count, err := CliGetblockcount()
	panicIf(err, "Failed to get block count")
	fmt.Println("Total height:", count)

	for i := count - 1; i >= 0; i-- {
		blchash, _ := CliGetblockhash(i)
		_, _, blc, err := CliGetblock(blchash, 2)
		panicIf(err, "Failed to get block")

		if !op.includeGenBlock && len(blc.Tx) < 2 {
			fmt.Println("TX len < 2, skiped", i)
			continue
		}

		fmt.Println("------------")
		if op.simpleBlock {
			m := map[string]interface{}{
				"hash":    blc.Hash,
				"confirm": blc.Confirmations,
				"tx":      blc.Tx,
			}
			fmt.Println("hight", i, jsonStr(m))
		} else {
			fmt.Println("hight", i, jsonStr(blc))
		}
		fmt.Println()

		for ti, txHash := range blc.Tx {
			tx, _ := CliGettransaction(txHash.Hash, true)

			if tx.TxID == "" {
				fmt.Printf("skiped empty tx %d/%d\n", i, ti)
				continue
			}

			if !op.includeCoinbaseTx && isCoinbaseTx(tx) {
				fmt.Println("skiped coinbase tx")
				continue
			}
			fmt.Println("tx", ti, jsonStr(tx))
		}
	}
}
