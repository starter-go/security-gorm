package sms

import (
	"fmt"
	"strings"

	"github.com/starter-go/security/auth"
)

const (
	myStepSendCode = "sendcode"
	myStepExecute  = "execute"
	myStepHelp     = "help"
)

// UserPhoneAuth ...
type UserPhoneAuth struct {

	//starter:component

	_as func(auth.Registry) //starter:as(".")

}

func (inst *UserPhoneAuth) _impl() (auth.Authenticator, auth.Registry, auth.Mechanism) {
	return inst, inst, inst
}

// ListRegistrations ...
func (inst *UserPhoneAuth) ListRegistrations() []*auth.Registration {
	r1 := &auth.Registration{
		Authenticator: inst,
		Authorizer:    nil,
		Mechanism:     inst,
		Priority:      0,
		Enabled:       true,
	}
	return []*auth.Registration{r1}
}

// Support ...
func (inst *UserPhoneAuth) Support(r auth.Request) bool {
	a := r.(auth.Authentication)
	want := a.Mechanism()
	have := auth.MechanismSMS
	return want == have
}

// Authenticate ...
func (inst *UserPhoneAuth) Authenticate(a auth.Authentication) ([]auth.Identity, error) {
	step := a.Step()
	switch step {
	case myStepSendCode:
		return inst.doSendCode(a)
	case myStepExecute:
		return inst.doExecute(a)
	default:
		return inst.doHelp(a)
	}
}

func (inst *UserPhoneAuth) doSendCode(a auth.Authentication) ([]auth.Identity, error) {

	return nil, fmt.Errorf("no impl")
}

func (inst *UserPhoneAuth) doExecute(a auth.Authentication) ([]auth.Identity, error) {

	return nil, fmt.Errorf("no impl ")
}

func (inst *UserPhoneAuth) doHelp(a auth.Authentication) ([]auth.Identity, error) {
	sep := ""
	builder := &strings.Builder{}
	steps := []string{myStepHelp, myStepSendCode, myStepExecute}
	for _, step := range steps {
		builder.WriteString(sep)
		builder.WriteString(step)
		sep = ", "
	}
	return nil, fmt.Errorf("help: UserPhoneAuth.steps = [%s]", builder.String())
}
