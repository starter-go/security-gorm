package testcom

import (
	"github.com/starter-go/application"
	"github.com/starter-go/security/auth"
)

// TestAuthWithSMS ... 测试授权
type TestAuthWithSMS struct {

	//starter:component
	_as func(application.Lifecycle) //starter:as(".")

	AC application.Context //starter:inject("context")

	AuthService auth.Service
}

// Life ...
func (inst *TestAuthWithSMS) Life() *application.Life {
	l := &application.Life{}
	l.OnLoop = inst.run
	return l
}

func (inst *TestAuthWithSMS) run() error {

	return nil
}
