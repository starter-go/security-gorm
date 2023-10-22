package testcom

import (
	"github.com/starter-go/application"
	"github.com/starter-go/security/rbac"
)

// Bootstrap ...
type Bootstrap struct {

	//starter:component
	_as func(application.Lifecycle) //starter:as(".")

	AC              application.Context //starter:inject("context")
	RbacAuthService rbac.AuthService    //starter:inject("#")

}

// Life ...
func (inst *Bootstrap) Life() *application.Life {
	return &application.Life{
		OnLoop: inst.loop,
	}
}

func (inst *Bootstrap) loop() error {
	p := inst.AC.GetAttributes().GetAttribute(AttrKeyBootParams).(*BootParams)
	p.Callback(inst)
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// AttrKeyBootParams ...
const AttrKeyBootParams = "theAttrKeyBootParams"

// BootParams ...
type BootParams struct {
	Callback func(b *Bootstrap)
}

////////////////////////////////////////////////////////////////////////////////
