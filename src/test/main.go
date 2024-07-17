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

	args := []string{}
	props := map[string]string{
		"debug.enabled":        "1",
		"debug.log-properties": "1",
	}
	mod := securitygorm.ModuleForTest()

	err := units.Run(&units.Config{
		Args:       args,
		Cases:      cname,
		Module:     mod,
		Properties: props,
		T:          t,
		UsePanic:   false,
	})

	if err != nil {
		t.Error(err)
	}
}
