package testcom

import (
	"github.com/starter-go/application"
	"github.com/starter-go/security/auth"
)

// TestAuthWithEmail ... 测试授权
type TestAuthWithEmail struct {

	//starter:component
	_as func(application.Lifecycle) //starter:as(".")

	AC application.Context //starter:inject("context")

	AuthService auth.Service
}

// Life ...
func (inst *TestAuthWithEmail) Life() *application.Life {
	l := &application.Life{}
	l.OnLoop = inst.run
	return l
}

func (inst *TestAuthWithEmail) run() error {

	return nil
}
