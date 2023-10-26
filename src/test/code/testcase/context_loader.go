package testcase

import (
	"os"

	"github.com/starter-go/security-gorm/modules/securitygormtest"
	"github.com/starter-go/starter"
)

type testingContextLoader struct {
}

func (inst *testingContextLoader) run(method, path string) {
	i := starter.Init(os.Args)
	i.MainModule(securitygormtest.ModuleForTest())
	props := i.GetParameters()
	props.SetParam("method", method)
	props.SetParam("path", path)
	i.WithPanic(true).Run()
}

////////////////////////////////////////////////////////////////////////////////

// Run ...
func Run() {
	tcl := &testingContextLoader{}
	tcl.run("", "")
}
