package main

import (
	"testing"

	"github.com/starter-go/security-gorm/modules/securitygorm"
	"github.com/starter-go/security-gorm/src/test/code/testcase"
	"github.com/starter-go/units"
)

func main() {
	testcase.Run()
}

func testWithCaseName(cname string, t *testing.T) {

	a := []string{}
	p := map[string]string{
		"debug.enabled":        "1",
		"debug.log-properties": "1",
	}
	m := securitygorm.ModuleForTest()

	c := &units.Context{
		Arguments:  a,
		Module:     m,
		Properties: p,
		T:          t,
		UsePanic:   true,
		Selector:   cname,
	}

	err := units.Run(c)
	if err != nil {
		t.Error(err)
	}
}
