package omnicli

import (
	"fmt"
	"os"
)

// bitcoin bin path
var (
	BasePath      = "/Users/some_user/Applications/bitcoin/bin" //see init()
	CmdOmnicored  = BasePath + "/omnicored"
	CmdBitcoinCli = BasePath + "/omnicore-cli"
)

func init() {
	fmt.Println("--init start---")
	defer fmt.Println("--init end---")

	fmt.Println(":Read env OMNI_BIN_PATH to configure command path")
	p := os.Getenv(OmniBinPathEnv)
	if p == "" {
		panic("使用omni需要bin path env: OMNI_BIN_PATH")
	}
	BasePath = p
	fmt.Println(":omni bin path:", BasePath)
	CmdOmnicored = BasePath + "/omnicored"     //windows may need change suffix
	CmdBitcoinCli = BasePath + "/omnicore-cli" //windows may need change suffix

}
