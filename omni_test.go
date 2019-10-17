package omnicli

import (
	"fmt"
	"testing"

	"github.com/dabankio/omnicli/testtool"

	"github.com/dabankio/omnicli/btcjson"
)

/// TextTx createRawTx, signTx, sendTx
func TestOmniSimpleTx(t *testing.T) {
	cli, killomnicored, err := RunOmnicored(&RunOptions{NewTmpDir: true})
	testtool.FailOnFlag(t, err != nil, "Failed to start btcd", err)
	defer killomnicored()

	var (
		addrs      []Addr
		a0, a1     Addr
		propertyID int
	)

	{
		NoPrintCmd(func() {
			addrs, err = cli.ToolGetSomeAddrs(5)
			testtool.FailOnFlag(t, err != nil, "Failed to get new address", err)
			a0, a1 = addrs[0], addrs[1]
		})
		// fmt.Printf("accounts, %#v\n", addrs)
		fmt.Println("accounts", ToJsonIndent(addrs))
	}
	{ //gen 101 to addr
		NoPrintCmd(func() {
			_, err := cli.Generatetoaddress(101, a0.Address, nil)
			testtool.FailOnFlag(t, err != nil, "Failed to generate to address ", err)
		})
	}
	var unspents []btcjson.ListUnspentResult
	{ // list unspent
		unspents, err = cli.Listunspent(0, 999, []string{a0.Address})
		testtool.FailOnErr(t, err, "Failed to list unspent")
		testtool.FailOnFlag(t, len(unspents) == 0, "no unspent find")
	}

	{ // 创建一个omni代币, a0 为持有人，
		cmd := OmniSenddissuancefixedCmd{
			Fromaddress: a0.Address,
			Ecosystem:   2, //2 fot test
			Typ:         1, // 1 for indivisible
			Previousid:  0, // 0 for new tokens
			Category:    "test_omni",
			Subcategory: "unit_test",
			Name:        "FakeUSDT",
			URL:         "",
			Data:        "",
			Amount:      "10000",
		}
		txHash, err := cli.OmniSendissuancefixed(&cmd)
		testtool.FailOnErr(t, err, "Failed to create omni coin")
		{ //生成几个块，确认刚才的交易
			_, err = cli.Generatetoaddress(2, a0.Address, nil)
			testtool.FailOnErr(t, err, "Failed to generate to address")
		}
		tx, err := cli.OmniGettransaction(txHash)
		testtool.FailOnErr(t, err, "Failed to get tx")
		propertyID = tx.Propertyid

		testtool.FailOnFlag(t, propertyID == 0, "Got property id error", propertyID)
	}

	{ // 代币创建完成后查询代币持有人的余额，应该等于总的发行量
		fmt.Println("-------then balance of new created property-----")
		bal, err := cli.OmniGetbalance(a0.Address, propertyID)
		testtool.FailOnErr(t, err, "Failed to get balance of owner")
		testtool.FailOnFlag(t, bal.Balance != "10000", "余额不符合预期")
	}

	{ // 代币从a0 simple send 到 a1
		unspents, err = cli.Listunspent(0, 999, []string{a0.Address})
		testtool.FailOnErr(t, err, "Failed to list unspent")
		testtool.FailOnFlag(t, len(unspents) == 0, "no unspent find")

		utxo := unspents[0]
		rawTX, err := cli.Createrawtransaction(btcjson.CreateRawTransactionCmd{
			Inputs: []btcjson.TransactionInput{{Txid: utxo.TxID, Vout: utxo.Vout}},
		})
		testtool.FailOnErr(t, err, "Failed to create raw tx")
		cli.DecodeAndPrintTX("Initial create", rawTX)

		payload, err := cli.OmniCreatepaloadSimplesend(propertyID, "233.3")
		testtool.FailOnErr(t, err, "Faied to create payload")

		rawTX, err = cli.OmniCreaterawtxOpreturn(rawTX, payload)
		testtool.FailOnErr(t, err, "Failed to create raw tx op return")

		cli.DecodeAndPrintTX("after opreturn", rawTX)

		pos := 1
		rawTX, err = cli.OmniCreaterawtxReference(rawTX, a1.Address, &pos)
		testtool.FailOnErr(t, err, "Failed to create raw tx reference")
		cli.DecodeAndPrintTX("after refrerence", rawTX)

		pos = 2
		rawTX, err = cli.OmniCreaterawtxChange(rawTX, []btcjson.PreviousDependentTxOutput{{
			TxID:         utxo.TxID,
			Vout:         utxo.Vout,
			Amount:       utxo.Amount,
			Value:        utxo.Amount,
			ScriptPubKey: utxo.ScriptPubKey,
		}}, a0.Address, 0.0006, &pos)
		testtool.FailOnErr(t, err, "Failed to create omni change")
		cli.DecodeAndPrintTX("after change", rawTX)

		signRet, err := cli.Signrawtransaction(btcjson.SignRawTransactionCmd{
			RawTx: rawTX,
			Inputs: &[]btcjson.RawTxInput{{
				Txid: utxo.TxID, Vout: utxo.Vout, ScriptPubKey: utxo.ScriptPubKey,
			}},
			PrivKeys: &[]string{a0.Privkey},
			Prevtxs: []btcjson.PreviousDependentTxOutput{{
				TxID: utxo.TxID, Vout: utxo.Vout, ScriptPubKey: utxo.ScriptPubKey, Amount: utxo.Amount,
			}},
		})
		testtool.FailOnErr(t, err, "Failed to sign tx")

		cli.DecodeAndPrintTX("已签名的tx", signRet.Hex)

		txid, err := cli.Sendrawtransaction(btcjson.SendRawTransactionCmd{
			HexTx: signRet.Hex,
		})
		testtool.FailOnErr(t, err, "Failed to send tx")
		fmt.Println("send txid", txid)
	}

	{ // 生成一个块确认usdt转账
		_, err = cli.Generatetoaddress(1, a0.Address, nil)
		// fmt.Println("应该被确认")
		testtool.FailOnErr(t, err, "Failed to generate to address")
	}

	{ // 确认代币转账成功
		bal, err := cli.OmniGetbalance(a1.Address, propertyID)
		testtool.FailOnErr(t, err, "Failed to get omni balance")
		testtool.FailOnFlag(t, bal.Balance != "233", "wrong balance, not 233")
		fmt.Println("bal:", ToJsonIndent(bal))
	}

}

/// TextTx createRawTx, signTx, sendTx
func TestBtcSimpleTx(t *testing.T) {
	cli, killomnicored, err := RunOmnicored(&RunOptions{NewTmpDir: true})
	testtool.FailOnFlag(t, err != nil, "Failed to start btcd", err)
	defer killomnicored()

	var addrs []Addr
	var a0, a1 Addr
	{
		addrs, err = cli.ToolGetSomeAddrs(5)
		testtool.FailOnFlag(t, err != nil, "Failed to get new address", err)
		a0 = addrs[0]
		a1 = addrs[1]
	}
	{ //gen 101 to addr
		_, err := cli.Generatetoaddress(101, a0.Address, nil)
		testtool.FailOnFlag(t, err != nil, "Failed to generate to address ", err)
	}
	var unspents []btcjson.ListUnspentResult
	{ // list unspent
		unspents, err = cli.Listunspent(0, 999, []string{a0.Address})
		testtool.FailOnFlag(t, err != nil, "Failed to list unspent", err)
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
			Outputs: map[string]interface{}{
				a1.Address: amount,
				a0.Address: unspent.Amount - amount - 0.001,
			},
		}
		rawHex, err := cli.Createrawtransaction(cmd)
		testtool.FailOnFlag(t, err != nil, "Failed to create raw tx", err)

		fmt.Println("Then decode rawHex")
		_, err = cli.Decoderawtransaction(btcjson.DecodeRawTransactionCmd{
			HexTx: rawHex,
		})
		testtool.FailOnFlag(t, err != nil, "Failed to decode raw tx", err)

		keys := []string{a0.Privkey}
		signRes, err := cli.Signrawtransaction(btcjson.SignRawTransactionCmd{
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
		testtool.FailOnFlag(t, err != nil, "Failed to sign with key raw tx", err)
		// fmt.Println("sign res", ToJsonIndent(signRes))

		fmt.Println("Then decode rawHex")
		decodedTxAfterSign, err := cli.Decoderawtransaction(btcjson.DecodeRawTransactionCmd{
			HexTx: signRes.Hex,
		})
		testtool.FailOnFlag(t, err != nil, "Failed to decode raw tx", err)
		fmt.Println("decodedTxAfterSign tx", ToJsonIndent(decodedTxAfterSign))

		sendRes, err := cli.Sendrawtransaction(btcjson.SendRawTransactionCmd{
			HexTx: signRes.Hex,
		})
		testtool.FailOnFlag(t, err != nil, "Failed to send raw tx", err)
		fmt.Println("send res:", sendRes)

		for _, vout := range decodedTxAfterSign.Vout {
			if len(vout.ScriptPubKey.Hex) == 0 {
				continue
			}
			decodeScript, err := cli.Decodescript(vout.ScriptPubKey.Hex)
			testtool.FailOnFlag(t, err != nil, "Failed to decode scriptPubkey", err)
			fmt.Println("vout:", ToJsonIndent(vout), "scriptPubkey decode:", ToJsonIndent(decodeScript))
		}
	}

}

func TestBTCMultisigTx(t *testing.T) {
	cli, killomnicored, err := RunOmnicored(&RunOptions{NewTmpDir: true})
	testtool.FailOnFlag(t, err != nil, "Failed to start btcd", err)
	defer killomnicored()

	var addrs []Addr
	var a0, a1, secondAddr, thirdAddr, fourthAddr Addr
	_ = fourthAddr
	// 0 把钱转给1+2+3多签(2-3)，1+3再转给4
	{
		addrs, err = cli.ToolGetSomeAddrs(5)
		testtool.FailOnFlag(t, err != nil, "Failed to get new address", err)
		a0, a1, secondAddr, thirdAddr, fourthAddr = addrs[0], addrs[1], addrs[2], addrs[3], addrs[4]
		fmt.Println("addrs")
		for _, a := range addrs {
			fmt.Printf("%s,\n", a.String())
		}
	}
	{ //gen 101 to addr
		_, err := cli.Generatetoaddress(101, a0.Address, nil)
		testtool.FailOnFlag(t, err != nil, "Failed to generate to address ", err)
	}

	var (
		// multisigAddres123       string
		createMultisigAddresRes btcjson.CreateMultiSigResult
		spentTx                 *btcjson.RawTx
	)
	_ = spentTx

	{ // 创建多签地址
		createMultisigAddresRes, err = cli.Createmultisig(2, []string{a1.Pubkey, secondAddr.Pubkey, thirdAddr.Pubkey}, nil)
		testtool.FailOnFlag(t, err != nil, "Failed to create multisig address", err)
		fmt.Println("生成多签地址的结果", ToJsonIndent(createMultisigAddresRes))

		//注意需要导入钱包，否则查不到unspent
		err = cli.Importaddress(btcjson.ImportAddressCmd{
			Address: createMultisigAddresRes.Address,
		})
		testtool.FailOnFlag(t, err != nil, "导入多签地址失败", err)

		_, err := cli.Decodescript(createMultisigAddresRes.RedeemScript)
		testtool.FailOnFlag(t, err != nil, "Failed to decode script", err)
		// fmt.Println("decoded redeemScript:", ToJsonIndent(decodeScript))
	}

	{ // 把0的钱交易给多签地址
		// unspents, err := cli.Listunspent(0, 999, []string{a0.Address}, nil, nil)
		// testtool.FailOnFlag(t, err != nil, "Failed to list unspent", err)
		// // fmt.Println("unspents", ToJsonIndent(unspents))

		// unspent := unspents[0]
		// //0->1 3.777 btc
		// amount := float64(17)
		// cmd := btcjson.CreateRawTransactionCmd{
		// 	Inputs: []btcjson.TransactionInput{
		// 		btcjson.TransactionInput{
		// 			Txid: unspent.TxID,
		// 			Vout: unspent.Vout,
		// 		},
		// 	},
		// 	Outputs: []map[string]interface{}{
		// 		map[string]interface{}{
		// 			createMultisigAddresRes.Address: amount,
		// 		},
		// 		map[string]interface{}{
		// 			a0.Address: unspent.Amount - amount - 0.001,
		// 		},
		// 	},
		// }
		// rawHex, err := cli.Createrawtransaction(cmd)
		// testtool.FailOnFlag(t, err != nil, "Failed to create raw tx", err)

		// keys := []string{a0.Privkey}
		// signRes, err := cli.Signrawtransaction(btcjson.SignRawTransactionCmd{
		// 	RawTx:    rawHex,
		// 	PrivKeys: &keys,
		// 	Prevtxs: []btcjson.PreviousDependentTxOutput{
		// 		btcjson.PreviousDependentTxOutput{
		// 			TxID:         unspent.TxID,
		// 			Vout:         unspent.Vout,
		// 			ScriptPubKey: unspent.ScriptPubKey,
		// 			Amount:       unspent.Amount,
		// 		},
		// 	},
		// })
		// testtool.FailOnFlag(t, err != nil, "Failed to sign with key raw tx", err)
		// // fmt.Println("sign res", ToJsonIndent(signRes))

		// sendRes, err := cli.Sendrawtransaction(btcjson.SendRawTransactionCmd{
		// 	HexTx: signRes.Hex,
		// })
		// testtool.FailOnFlag(t, err != nil, "Failed to send raw tx", err)
		// // fmt.Println("send res:", sendRes)

		// // tx, err := cli.Gettransaction(sendRes, true)
		// spentTx, err = cli.Getrawtransaction(btcjson.GetRawTransactionCmd{
		// 	Txid: sendRes,
		// })
		// testtool.FailOnFlag(t, err != nil, "Failed to get tx", err)
		// fmt.Println("to spent tx(mutisig)", ToJsonIndent(spentTx))
	}

	{ //生成一个block来确认下刚才的交易
		_, err = cli.Generatetoaddress(1, a0.Address, nil)
		testtool.FailOnFlag(t, err != nil, "Failed to generate to address", err)
	}

	// {
	// 	unspents, err := cli.Listunspent(0, 999, []string{a0.Address}, btcjson.Bool(true), nil)
	// 	testtool.FailOnFlag(t, err != nil, "Failed to list unspent", err)
	// 	fmt.Println("a0上的UTXO", ToJsonIndent(unspents))
	// }

	{ //现在多签地址里的钱要转给fourthAddr
		// fmt.Println("收取多签地址转来的前的地址", fourthAddr.Address)
		// unspents, err := cli.Listunspent(0, 999, []string{createMultisigAddresRes.Address}, btcjson.Bool(true), nil)
		// testtool.FailOnFlag(t, err != nil, "Failed to list unspent", err)
		// fmt.Println("多签地址上的UTXO", ToJsonIndent(unspents))

		// // amt, _ := cli.Getreceivedbyaddress(createMultisigAddresRes.Address, 0)
		// amt, _ := cli.Getreceivedbyaddress(a0.Address, 0)
		// fmt.Println("Received amt:", amt)

		// spentVout := spentTx.Vout[0]
		// amount := float64(9)
		// cmd := btcjson.CreateRawTransactionCmd{
		// 	Inputs: []btcjson.TransactionInput{
		// 		btcjson.TransactionInput{
		// 			Txid: spentTx.Txid,
		// 			Vout: spentVout.N,
		// 		},
		// 	},
		// 	Outputs: []map[string]interface{}{
		// 		map[string]interface{}{
		// 			fourthAddr.Address: amount,
		// 		},
		// 		map[string]interface{}{
		// 			createMultisigAddresRes.Address: spentVout.Value - amount - 0.001,
		// 		},
		// 	},
		// }
		// rawHex, err := cli.Createrawtransaction(cmd)
		// testtool.FailOnFlag(t, err != nil, "Failed to create raw tx", err)
		// dTx, err := cli.Decoderawtransaction(btcjson.DecodeRawTransactionCmd{
		// 	HexTx: rawHex,
		// })
		// testtool.FailOnFlag(t, err != nil, "Failed to decode rawHex", err)
		// fmt.Println("创建的多签raw tx", ToJsonIndent(dTx))

		// // for _, ke := range []string{a1.Privkey} {
		// // for i, ke := range []string{a1.Privkey, thirdAddr.Privkey} {
		// for i, ke := range []string{thirdAddr.Privkey, a1.Privkey} {
		// 	keys := []string{ke}
		// 	signRes, err := cli.Signrawtransaction(btcjson.SignRawTransactionCmd{
		// 		RawTx:    rawHex,
		// 		PrivKeys: &keys,
		// 		Prevtxs: []btcjson.PreviousDependentTxOutput{
		// 			btcjson.PreviousDependentTxOutput{
		// 				TxID:         spentTx.Txid,
		// 				Vout:         spentVout.N,
		// 				ScriptPubKey: spentVout.ScriptPubKey.Hex,
		// 				Amount:       spentVout.Value,
		// 				RedeemScript: createMultisigAddresRes.RedeemScript,
		// 			},
		// 		},
		// 	})
		// 	rawHex = signRes.Hex
		// 	testtool.FailOnFlag(t, err != nil, "Failed to sign with key raw tx", err)
		// 	fmt.Println("第n次签名的结果", i, ToJsonIndent(signRes))
		// 	deTx, err := cli.Decoderawtransaction(btcjson.DecodeRawTransactionCmd{
		// 		HexTx: rawHex,
		// 	})
		// 	testtool.FailOnFlag(t, err != nil, "Failed to decode raw tx in multisig", err)
		// 	fmt.Println("第n次签名后对rawTx的解码", i, ToJsonIndent(deTx))
		// }

		// multisigTxid, err := cli.Sendrawtransaction(btcjson.SendRawTransactionCmd{
		// 	HexTx: rawHex,
		// })
		// testtool.FailOnFlag(t, err != nil, "Failed to send raw tx", err)
		// fmt.Println("send(multisig) res:", multisigTxid)
		// mtx, err := cli.Getrawtransaction(btcjson.GetRawTransactionCmd{
		// 	Txid:    multisigTxid,
		// 	Verbose: btcjson.Int(1),
		// })
		// testtool.FailOnFlag(t, err != nil, "Failed to get raw multisig tx", err)
		// fmt.Println("raw multisig tx", ToJsonIndent(mtx))
	}

	{
		_, err := cli.Generatetoaddress(1, a0.Address, nil)
		testtool.FailOnFlag(t, err != nil, "Failed to send to addr 0", err)
	}

	{ //列出multisig的unspent
		// unspents, err := cli.Listunspent(0, 999, []string{createMultisigAddresRes.Address}, nil, nil)
		// testtool.FailOnFlag(t, err != nil, "Failed to list unspent", err)
		// fmt.Println("unspent of multisig", ToJsonIndent(unspents))
	}
	{ //最后列出转出的unspent
		// unspents, err := cli.Listunspent(0, 999, []string{fourthAddr.Address}, nil, nil)
		// testtool.FailOnFlag(t, err != nil, "Failed to list unspent", err)
		// fmt.Println("unspent of 4", ToJsonIndent(unspents))
	}

	// PrintCmdOut = false
	// scanChain(scanOps{simpleBlock: true})
}
