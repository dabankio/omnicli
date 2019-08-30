package omnicli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// RunOptions .
type RunOptions struct {
	NewTmpDir bool //创建一个临时的目录，并在完成后立即删除这个目录
	RPCPort   uint
	Args      []string
}

type killHook func() error

// RunOmnicored .
func RunOmnicored(optionsPtr *RunOptions) (*Cli, func(), error) {
	killHooks := []killHook{}

	var options RunOptions
	if optionsPtr == nil {
		options = RunOptions{}
	} else {
		options = *optionsPtr
	}

	if options.RPCPort == 0 {
		options.RPCPort = RPCPortRegtest
	}

	var dataDir string
	if options.NewTmpDir {
		for _, arg := range options.Args {
			if strings.Contains(arg, "-datadir=") {
				return nil, nil, fmt.Errorf("参数里似乎已经指定了-datadir  >> %v", arg)
			}
		}

		tmpDir := strings.TrimRight(os.TempDir(), "/")
		dataDir = tmpDir + "/omnicli_omni_data_tmp_" + time.Now().Format(time.RFC3339) + "/"
		err := os.MkdirAll(dataDir, 0777)
		if err != nil {
			return nil, nil, fmt.Errorf("cannot create tmp dir: %v, err: %v", dataDir, err)
		}
		options.Args = append(options.Args, "-datadir="+dataDir)

		killHooks = append(killHooks, func() error {
			return os.RemoveAll(dataDir)
		})
	}

	if cmdIsPortContainsNameRunning(options.RPCPort, "omnicore") {
		return nil, nil, fmt.Errorf("omnicore 似乎已经运行在%d端口了", options.RPCPort)
	}

	closeChan := make(chan struct{})

	args := []string{
		"-regtest",
		"-txindex",
		"-rpcuser=rpcusr",
		"-rpcpassword=233",
		fmt.Sprintf("-rpcport=%d", options.RPCPort),
	}
	args = append(args, options.Args...)

	cli, err := NewCliFromRunningOmnicored(RunningOmnicoredOptions{
		RPCPort:     options.RPCPort,
		RPCUser:     "rpcusr",
		RPCPassword: "233",
		DataDir:     dataDir,
		NetID:       NetRegtest,
	})
	if err != nil {
		return nil, nil, err
	}

	cmd := exec.Command(CmdOmnicored, args...)
	fmt.Println(cmd.Args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		return nil, nil, err
	}
	go func() {
		fmt.Println("Wait for message to kill omnicored")
		<-closeChan
		fmt.Println("Received message,killing omnicored regtest")

		if e := cmd.Process.Kill(); e != nil {
			fmt.Println("关闭 omnicored 时发生异常", e)
		}
		fmt.Println("关闭 omnicored 完成")
		closeChan <- struct{}{}
	}()

	// err = cmd.Wait()
	fmt.Println("等待1.5秒,让 omnicored 启动")
	time.Sleep(time.Millisecond * 1500)
	return cli, func() {
		closeChan <- struct{}{}

	}, nil
}
