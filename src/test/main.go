package main

import (
	"github.com/starter-go/security-gorm/src/test/code/testboot"
	"github.com/starter-go/security-gorm/src/test/code/testcom"
	"github.com/starter-go/vlog"
)

func main() {
	p := &testcom.BootParams{}
	p.Callback = func(b *testcom.Bootstrap) {
		vlog.Info("hello, testcase1")
	}
	testboot.Boot(p)
}
