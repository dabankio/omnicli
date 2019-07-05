package btccli

import (
	"fmt"
	"testing"

	"github.com/lemon-sunxiansong/btccli/btcjson"
)

/// TextTx createRawTx, signTx, sendTx
func TestSimpleTx(t *testing.T) {
	cc, err := BitcoindRegtest()
	trueThenFailNow(t, err != nil, "Failed to start btcd", err)
	defer func() {
		cc <- struct{}{}
	}()

	var addrs []Addr
	var zeroAddr, firstAddr Addr
	{
		addrs, err = CliToolGetSomeAddrs(5)
		trueThenFailNow(t, err != nil, "Failed to get new address", err)
		zeroAddr = addrs[0]
		firstAddr = addrs[1]
	}
	{ //gen 101 to addr
		_, err := CliGeneratetoaddress(101, zeroAddr.Address, nil)
		trueThenFailNow(t, err != nil, "Failed to generate to address ", err)
	}
	var unspents []btcjson.ListUnspentResult
	{ // list unspent
		unspents, err = CliListunspent(0, 999, []string{zeroAddr.Address}, nil, nil)
		trueThenFailNow(t, err != nil, "Failed to list unspent", err)
		fmt.Println("unspents", ToJsonIndent(unspents))
	}

	{ // simple tx, o把btc转给1
		unspent := unspents[0]
		//0->1 3.777 btc
		amount := float64(17)
		cmd := btcjson.CreateRawTransactionCmd{
			Inputs: []btcjson.TransactionInput{
				btcjson.TransactionInput{
					Txid: unspent.TxID,
					Vout: unspent.Vout,
				},
			},
			Outputs: []map[string]interface{}{
				map[string]interface{}{
					firstAddr.Address: amount,
				},
				map[string]interface{}{
					zeroAddr.Address: unspent.Amount - amount - 0.001,
				},
			},
		}
		rawHex, err := CliCreaterawtransaction(cmd)
		trueThenFailNow(t, err != nil, "Failed to create raw tx", err)


		fmt.Println("Then decode rawHex")
		_, err = CliDecoderawtransaction(btcjson.DecodeRawTransactionCmd{
			HexTx: rawHex,
		})
		trueThenFailNow(t, err != nil, "Failed to decode raw tx", err)

		keys := []string{zeroAddr.Privkey}
		signRes, err := CliSignrawtransactionwithkey(btcjson.SignRawTransactionCmd{
			RawTx:    rawHex,
			PrivKeys: &keys,
			Prevtxs: []btcjson.PreviousDependentTxOutput{
				btcjson.PreviousDependentTxOutput{
					TxID:         unspent.TxID,
					Vout:         unspent.Vout,
					ScriptPubKey: unspent.ScriptPubKey,
					Amount:       unspent.Amount,
				},
			},
		})
		trueThenFailNow(t, err != nil, "Failed to sign with key raw tx", err)
		// fmt.Println("sign res", ToJsonIndent(signRes))

		fmt.Println("Then decode rawHex")
		decodedTxAfterSign, err := CliDecoderawtransaction(btcjson.DecodeRawTransactionCmd{
			HexTx: signRes.Hex,
		})
		trueThenFailNow(t, err != nil, "Failed to decode raw tx", err)
		fmt.Println("decodedTxAfterSign tx", ToJsonIndent(decodedTxAfterSign))

		sendRes, err := CliSendrawtransaction(btcjson.SendRawTransactionCmd{
			HexTx: signRes.Hex,
		})
		trueThenFailNow(t, err != nil, "Failed to send raw tx", err)
		fmt.Println("send res:", sendRes)

		for _, vout := range decodedTxAfterSign.Vout {
			if len(vout.ScriptPubKey.Hex) == 0 {
				continue
			}
			decodeScript, err := CliDecodescript(vout.ScriptPubKey.Hex)
			trueThenFailNow(t, err != nil, "Failed to decode scriptPubkey", err)
			fmt.Println("vout:", ToJsonIndent(vout), "scriptPubkey decode:", ToJsonIndent(decodeScript))
		}
	}

}

func TestMultisigTx(t *testing.T) {
	cc, err := BitcoindRegtest()
	trueThenFailNow(t, err != nil, "Failed to start btcd", err)
	defer func() {
		cc <- struct{}{}
	}()

	var addrs []Addr
	var zeroAddr, firstAddr, secondAddr, thirdAddr, fourthAddr Addr
	// 0 把钱转给1+2+3多签(2-3)，1+3再转给4
	{
		addrs, err = CliToolGetSomeAddrs(5)
		trueThenFailNow(t, err != nil, "Failed to get new address", err)
		zeroAddr, firstAddr, secondAddr, thirdAddr, fourthAddr = addrs[0], addrs[1], addrs[2], addrs[3], addrs[4]
		fmt.Println("addrs")
		for _, a := range addrs {
			fmt.Printf("%s,\n", a.String())
		}
	}
	{ //gen 101 to addr
		_, err := CliGeneratetoaddress(101, zeroAddr.Address, nil)
		trueThenFailNow(t, err != nil, "Failed to generate to address ", err)
	}

	var (
		// multisigAddres123       string
		createMultisigAddresRes btcjson.CreateMultiSigResult
		spentTx                 *btcjson.RawTx
	)

	{ // 创建多签地址
		createMultisigAddresRes, err = CliCreatemultisig(2, []string{firstAddr.Pubkey, secondAddr.Pubkey, thirdAddr.Pubkey}, nil)
		trueThenFailNow(t, err != nil, "Failed to create multisig address", err)
		fmt.Println("生成多签地址的结果", ToJsonIndent(createMultisigAddresRes))

		//注意需要导入钱包，否则查不到unspent
		err = CliImportaddress(btcjson.ImportAddressCmd{
			Address: createMultisigAddresRes.Address,
		})
		trueThenFailNow(t, err != nil, "导入多签地址失败", err)

		_, err := CliDecodescript(createMultisigAddresRes.RedeemScript)
		trueThenFailNow(t, err != nil, "Failed to decode script", err)
		// fmt.Println("decoded redeemScript:", ToJsonIndent(decodeScript))
	}

	{ // 把0的钱交易给多签地址
		unspents, err := CliListunspent(0, 999, []string{zeroAddr.Address}, nil, nil)
		trueThenFailNow(t, err != nil, "Failed to list unspent", err)
		// fmt.Println("unspents", ToJsonIndent(unspents))

		unspent := unspents[0]
		//0->1 3.777 btc
		amount := float64(17)
		cmd := btcjson.CreateRawTransactionCmd{
			Inputs: []btcjson.TransactionInput{
				btcjson.TransactionInput{
					Txid: unspent.TxID,
					Vout: unspent.Vout,
				},
			},
			Outputs: []map[string]interface{}{
				map[string]interface{}{
					createMultisigAddresRes.Address: amount,
				},
				map[string]interface{}{
					zeroAddr.Address: unspent.Amount - amount - 0.001,
				},
			},
		}
		rawHex, err := CliCreaterawtransaction(cmd)
		trueThenFailNow(t, err != nil, "Failed to create raw tx", err)

		keys := []string{zeroAddr.Privkey}
		signRes, err := CliSignrawtransactionwithkey(btcjson.SignRawTransactionCmd{
			RawTx:    rawHex,
			PrivKeys: &keys,
			Prevtxs: []btcjson.PreviousDependentTxOutput{
				btcjson.PreviousDependentTxOutput{
					TxID:         unspent.TxID,
					Vout:         unspent.Vout,
					ScriptPubKey: unspent.ScriptPubKey,
					Amount:       unspent.Amount,
				},
			},
		})
		trueThenFailNow(t, err != nil, "Failed to sign with key raw tx", err)
		// fmt.Println("sign res", ToJsonIndent(signRes))

		sendRes, err := CliSendrawtransaction(btcjson.SendRawTransactionCmd{
			HexTx: signRes.Hex,
		})
		trueThenFailNow(t, err != nil, "Failed to send raw tx", err)
		// fmt.Println("send res:", sendRes)

		// tx, err := CliGettransaction(sendRes, true)
		spentTx, err = CliGetrawtransaction(btcjson.GetRawTransactionCmd{
			Txid: sendRes,
		})
		trueThenFailNow(t, err != nil, "Failed to get tx", err)
		// fmt.Println("to spent tx(mutisig)", ToJsonIndent(spentTx))
	}

	{ //生成一个block来确认下刚才的交易
		_, err = CliGeneratetoaddress(1, zeroAddr.Address, nil)
		trueThenFailNow(t, err != nil, "Failed to generate to address", err)
	}

	// {
	// 	unspents, err := CliListunspent(0, 999, []string{zeroAddr.Address}, btcjson.Bool(true), nil)
	// 	trueThenFailNow(t, err != nil, "Failed to list unspent", err)
	// 	fmt.Println("zeroAddr上的UTXO", ToJsonIndent(unspents))
	// }

	{ //现在多签地址里的钱要转给fourthAddr
		fmt.Println("收取多签地址转来的前的地址", fourthAddr.Address)
		unspents, err := CliListunspent(0, 999, []string{createMultisigAddresRes.Address}, btcjson.Bool(true), nil)
		trueThenFailNow(t, err != nil, "Failed to list unspent", err)
		fmt.Println("多签地址上的UTXO", ToJsonIndent(unspents))

		// amt, _ := CliGetreceivedbyaddress(createMultisigAddresRes.Address, 0)
		amt, _ := CliGetreceivedbyaddress(zeroAddr.Address, 0)
		fmt.Println("Received amt:", amt)

		spentVout := spentTx.Vout[0]
		amount := float64(9)
		cmd := btcjson.CreateRawTransactionCmd{
			Inputs: []btcjson.TransactionInput{
				btcjson.TransactionInput{
					Txid: spentTx.Txid,
					Vout: spentVout.N,
				},
			},
			Outputs: []map[string]interface{}{
				map[string]interface{}{
					fourthAddr.Address: amount,
				},
				map[string]interface{}{
					createMultisigAddresRes.Address: spentVout.Value - amount - 0.001,
				},
			},
		}
		rawHex, err := CliCreaterawtransaction(cmd)
		trueThenFailNow(t, err != nil, "Failed to create raw tx", err)
		dTx, err := CliDecoderawtransaction(btcjson.DecodeRawTransactionCmd{
			HexTx: rawHex,
		})
		trueThenFailNow(t, err != nil, "Failed to decode rawHex", err)
		fmt.Println("创建的多签raw tx", ToJsonIndent(dTx))

		// for _, ke := range []string{firstAddr.Privkey} {
		// for i, ke := range []string{firstAddr.Privkey, thirdAddr.Privkey} {
		for i, ke := range []string{thirdAddr.Privkey, firstAddr.Privkey} {
			keys := []string{ke}
			signRes, err := CliSignrawtransactionwithkey(btcjson.SignRawTransactionCmd{
				RawTx:    rawHex,
				PrivKeys: &keys,
				Prevtxs: []btcjson.PreviousDependentTxOutput{
					btcjson.PreviousDependentTxOutput{
						TxID:         spentTx.Txid,
						Vout:         spentVout.N,
						ScriptPubKey: spentVout.ScriptPubKey.Hex,
						Amount:       spentVout.Value,
						RedeemScript: createMultisigAddresRes.RedeemScript,
					},
				},
			})
			rawHex = signRes.Hex
			trueThenFailNow(t, err != nil, "Failed to sign with key raw tx", err)
			fmt.Println("第n次签名的结果", i, ToJsonIndent(signRes))
			deTx, err := CliDecoderawtransaction(btcjson.DecodeRawTransactionCmd{
				HexTx: rawHex,
			})
			trueThenFailNow(t, err != nil, "Failed to decode raw tx in multisig", err)
			fmt.Println("第n次签名后对rawTx的解码", i, ToJsonIndent(deTx))
		}

		multisigTxid, err := CliSendrawtransaction(btcjson.SendRawTransactionCmd{
			HexTx: rawHex,
		})
		trueThenFailNow(t, err != nil, "Failed to send raw tx", err)
		fmt.Println("send(multisig) res:", multisigTxid)
		mtx, err := CliGetrawtransaction(btcjson.GetRawTransactionCmd{
			Txid:    multisigTxid,
			Verbose: btcjson.Int(1),
		})
		trueThenFailNow(t, err != nil, "Failed to get raw multisig tx", err)
		fmt.Println("raw multisig tx", ToJsonIndent(mtx))
	}

	{
		_, err := CliGeneratetoaddress(1, zeroAddr.Address, nil)
		trueThenFailNow(t, err != nil, "Failed to send to addr 0", err)
	}

	{ //列出multisig的unspent
		unspents, err := CliListunspent(0, 999, []string{createMultisigAddresRes.Address}, nil, nil)
		trueThenFailNow(t, err != nil, "Failed to list unspent", err)
		fmt.Println("unspent of multisig", ToJsonIndent(unspents))
	}
	{ //最后列出转出的unspent
		unspents, err := CliListunspent(0, 999, []string{fourthAddr.Address}, nil, nil)
		trueThenFailNow(t, err != nil, "Failed to list unspent", err)
		fmt.Println("unspent of 4", ToJsonIndent(unspents))
	}

	// PrintCmdOut = false
	// scanChain(scanOps{simpleBlock: true})
}
