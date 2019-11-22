package omnicli

import (
	"log"
	"os"
)

// omnicore bin path
var (
	BasePath     = "/Users/some_user/Applications/omnicore/bin" //see init()
	CmdOmnicored = BasePath + "/omnicored"
	CmdOmniCli   = BasePath + "/omnicore-cli"
)

func init() {
	const xxx = "========================"
	
	log.Printf("%s omnicli init start%s\n", xxx, xxx)
	defer log.Printf("%s omnicli init end  %s\n", xxx, xxx)

	log.Println("使用omnicli你需要:配置环境变量 OMNI_BIN_PATH 指向omnicore/bin目录")
	log.Println(":Read env OMNI_BIN_PATH to configure command path")
	p := os.Getenv(OmniBinPathEnv)
	if p == "" {
		panic("使用omni需要bin path env: OMNI_BIN_PATH")
	}
	BasePath = p
	log.Println(":omni bin path:", BasePath)
	CmdOmnicored = BasePath + "/omnicored"  //windows may need change suffix
	CmdOmniCli = BasePath + "/omnicore-cli" //windows may need change suffix
	log.Println("server cmd:", CmdOmnicored)
	log.Println("cli cmd:", CmdOmniCli)
}
