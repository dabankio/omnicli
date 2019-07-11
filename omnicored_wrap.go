package btccli

import (
	"fmt"
	"os/exec"
	"time"
)

// StartOmnicored 启动bitcoind -regtest 用以测试.
// return close channel (cc)
func StartOmnicored() (closeChan chan struct{}, err error) {
	if cmdIsPortContainsNameRunning(RPCPortRegtest, "bitcoin") {
		return nil, fmt.Errorf("bitcoind 似乎已经运行在18443端口了,不先杀掉的话数据可能有问题")
	}

	closeChan = make(chan struct{})

	//bitcoin/share/rpcauth$ python3 rpcauth.py rpcusr 233
	//String to be appended to bitcoin.conf:
	//rpcauth=rpcusr:656f9dabc62f0eb697c801369617dc60$422d7fca742d4a59460f941dc9247c782558367edcbf1cd790b2b7ff5624fc1b
	//Your password:
	//233
	cmd := exec.Command(CmdOmnicored,
		"-regtest",
		// "-testnet",
		// "-deprecatedrpc=generate",
		"-txindex",
		"-rpcport=18443",
		"-rpcuser=rpcusr",
		"-rpcpassword=233",
	)
	fmt.Println(cmd.Args)
	err = cmd.Start()
	if err != nil {
		return
	}
	go func() {
		fmt.Println("Wait for message to kill bitcoind")
		<-closeChan
		fmt.Println("Received message,killing bitcoind regtest")

		if e := cmd.Process.Kill(); e != nil {
			fmt.Println("关闭 bitcoind 时发生异常", e)
		}
		fmt.Println("关闭 bitcoind 完成")
		closeChan <- struct{}{}
	}()

	// err = cmd.Wait()
	fmt.Println("等待1.8秒,让 bitcoind 启动")
	time.Sleep(time.Millisecond * 1800)
	return
}