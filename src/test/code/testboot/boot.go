package testboot

import (
	"os"

	"github.com/starter-go/security-gorm/modules/securitygormtest"
	"github.com/starter-go/security-gorm/src/test/code/testcom"

	"github.com/starter-go/starter"
)

// Boot ...
func Boot(p *testcom.BootParams) {
	i := starter.Init(os.Args)
	i.GetAttributes().SetAttribute(testcom.AttrKeyBootParams, p)
	i.MainModule(securitygormtest.ModuleForTest())
	i.WithPanic(true).Run()
}
