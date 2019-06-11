package btccli

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// global vars
var (
	PrintCmdOut = true
)

func cmdAndPrint(cmd *exec.Cmd) string {
	return cmdThenPrint(cmd, true)
}

// 执行cmd,然后将输出打印到控制台并return
func cmdThenPrint(cmd *exec.Cmd, print bool) string {
	if print {
		fmt.Println("[CMD]", cmd.Args)
	}
	stderr, err := cmd.StderrPipe()
	panicIf(err, "Failed to get stderr pip ")

	stdout, err := cmd.StdoutPipe()
	panicIf(err, fmt.Sprintf("Failed to get stdout pipe %v", err))
	err = cmd.Start()
	panicIf(err, fmt.Sprintf("Failed to start cmd %v", err))
	b, err := ioutil.ReadAll(stdout)
	panicIf(err, fmt.Sprintf("Failed to read cmd (%v) stdout, %v", cmd, err))
	out := string(b)
	if PrintCmdOut && print {
		fmt.Println(out)
	}

	bo, err := ioutil.ReadAll(stderr)
	panicIf(err, "Failed to read stderr")
	out += string(bo)

	cmd.Wait()
	stdout.Close()
	stderr.Close()
	return strings.TrimSpace(out)
}

// 进程名包含$name的端口$port是否在运行中
func cmdIsPortContainsNameRunning(port uint, name string) bool {
	if strings.Contains(runtime.GOOS, "linux") {
		checkPortCmd := exec.Command("netstat", "-ntpl")
		cmdPrint := cmdAndPrint(checkPortCmd)
		if strings.Contains(cmdPrint, strconv.Itoa(int(port))) && strings.Contains(cmdPrint, name) {
			return true
		}
		return false
	} else if strings.Contains(runtime.GOOS, "darwin") {
		checkPortCmd := exec.Command("lsof", "-i", "tcp:18443")
		cmdPrint := cmdAndPrint(checkPortCmd)
		if strings.Contains(cmdPrint, strconv.Itoa(int(port))) && strings.Contains(cmdPrint, "bitcoin") {
			return true
		}
		return false
	} else {
		panic("其他平台尚未實現")
	}
}
