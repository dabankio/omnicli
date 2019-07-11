package omnicli

import (
	"fmt"
	"github.com/lemon-sunxiansong/omnicli/testtool"
	"testing"
)

func TestCliToolGetSomeAddrs(t *testing.T) {
	ast := testtool.NewSimpleAssert(t)
	cc, err := StartOmnicored()
	ast.FailOnErr(err, "start d")
	defer func() {
		cc <- struct{}{}
	}()

	addrs, err := CliToolGetSomeAddrs(10)
	ast.FailOnErr(err, "get addrs")
	// fmt.Println(ToJsonIndent(addrs))
	for _, add := range addrs {
		fmt.Println(add.String())
	}

}
