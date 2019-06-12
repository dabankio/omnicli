package btccli

type Addr struct {
	Address string
	Privkey string
	Pubkey  string
}

// CliToolGetSomeAddrs 一次获取n个地址（包含pub-priv key)
func CliToolGetSomeAddrs(n int) ([]Addr, error) {
	var addrs []Addr
	for i := 0; i < n; i++ {
		add, err := CliGetnewaddress(nil, nil)
		if err != nil {
			return nil, err
		}
		info, err := CliGetAddressInfo(add)
		if err != nil {
			return nil, err
		}
		dump, err := CliDumpprivkey(add)
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, Addr{
			Address: add, Privkey: dump, Pubkey: info.Pubkey,
		})
	}
	return addrs, nil
}
