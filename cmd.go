package btccliwrap

import (
	"fmt"
	"io/ioutil"
	"os/exec"
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
	if print {
		fmt.Println(out)
	}

	bo, err := ioutil.ReadAll(stderr)
	panicIf(err, "Failed to read stderr")
	out += string(bo)

	cmd.Wait()
	stdout.Close()
	stderr.Close()
	return out
}


