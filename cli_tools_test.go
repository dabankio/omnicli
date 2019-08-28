package omnicli

import (
	"fmt"
	"github.com/lomocoin/omnicli/testtool"
	"testing"
)

func TestCliToolGetSomeAddrs(t *testing.T) {
	ast := testtool.NewSimpleAssert(t)
	cc, err := StartOmnicored()
	ast.FailOnErr(err, "start d")
	defer func() {
		cc <- struct{}{}
	}()

	addrs, err := CliToolGetSomeAddrs(5)
	ast.FailOnErr(err, "get addrs")
	// fmt.Println(ToJsonIndent(addrs))
	fmt.Printf("%#v", addrs)

}
