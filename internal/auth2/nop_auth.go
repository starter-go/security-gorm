package auth2

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/security/auth"
)

// NopAuth ...
type NopAuth struct {

	//starter:component

	_as func(auth.Registry) //starter:as(".")

}

func (inst *NopAuth) _impl() (auth.Authorizer, auth.Mechanism, auth.Registry) {
	return inst, inst, inst
}

// ListRegistrations ...
func (inst *NopAuth) ListRegistrations() []*auth.Registration {

	r1 := &auth.Registration{
		Authorizer: inst,
		Mechanism:  inst,
		Priority:   -100,
		Enabled:    true,
	}

	return []*auth.Registration{r1}
}

// Support ...
func (inst *NopAuth) Support(r auth.Request) bool {
	req, ok := r.(auth.Authorization)
	if ok {
		step := req.Step()
		if step == rbac.StepHelp {
			return true
		}
		if step == rbac.StepSendCode {
			return true
		}
	}
	return false
}

// Authorize ...
func (inst *NopAuth) Authorize(a auth.Authorization) error {
	return nil // nop
}
