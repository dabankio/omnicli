package omnicli

import (
	"fmt"
	"github.com/dabankio/omnicli/testtool"
	"testing"
)

func TestCliToolGetSomeAddrs(t *testing.T) {
	ast := testtool.NewSimpleAssert(t)
	cli, killomnicored, err := RunOmnicored(&RunOptions{NewTmpDir: true})
	ast.FailOnErr(err, "start d")
	defer killomnicored()

	addrs, err := cli.ToolGetSomeAddrs(5)
	ast.FailOnErr(err, "get addrs")
	// fmt.Println(ToJsonIndent(addrs))
	fmt.Printf("%#v", addrs)

}
