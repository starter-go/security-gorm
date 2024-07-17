package testcom

import (
	"context"
	"encoding/json"

	"github.com/starter-go/application"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/src/test/code/cases"
	"github.com/starter-go/security/auth"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

// TestAuthWithEmail ... 测试授权
type TestAuthWithEmail struct {

	//starter:component
	_as func(units.Units) //starter:as(".")

	AC          application.Context //starter:inject("context")
	AuthService rbac.AuthService    //starter:inject("#")

	// AuthService auth.Service

}

func (inst *TestAuthWithEmail) _impl() units.Units {
	return inst
}

// Units ...
func (inst *TestAuthWithEmail) Units(list []*units.Registration) []*units.Registration {

	list = append(list, &units.Registration{
		Name:    cases.HelpByMail,
		Enabled: true,
		Test:    inst.runHelpByMail,
	})
	list = append(list, &units.Registration{
		Name:    cases.SendcodeByMail,
		Enabled: true,
		Test:    inst.runSendCodeByMail,
	})
	list = append(list, &units.Registration{
		Name:    cases.LoginByMail,
		Enabled: true,
		Test:    inst.runLoginByMail,
	})

	return list
}

func (inst *TestAuthWithEmail) runHelpByMail() error {

	ctx := context.Background()
	action := auth.ActionLogin
	mech := auth.MechanismEmail
	alist := make([]*rbac.AuthDTO, 0)

	alist = append(alist, &rbac.AuthDTO{
		Mechanism: mech,
		Action:    action,
		Step:      rbac.StepHelp,
		Account:   "[helper]",
	})

	alist2, err := inst.AuthService.Handle(ctx, action, alist)
	return inst.logResult(alist2, err)
}

func (inst *TestAuthWithEmail) runSendCodeByMail() error {

	ctx := context.Background()
	action := auth.ActionLogin
	mech := auth.MechanismEmail
	alist := make([]*rbac.AuthDTO, 0)

	alist = append(alist, &rbac.AuthDTO{
		Mechanism: mech,
		Action:    action,
		Step:      rbac.StepSendCode,
		Account:   "demo1@test.example.com",
	})

	alist2, err := inst.AuthService.Handle(ctx, action, alist)
	return inst.logResult(alist2, err)
}

func (inst *TestAuthWithEmail) runLoginByMail() error {

	ctx := context.Background()
	action := auth.ActionLogin
	mech := auth.MechanismEmail
	alist := make([]*rbac.AuthDTO, 0)

	alist = append(alist, &rbac.AuthDTO{
		Mechanism: mech,
		Action:    action,
		Step:      rbac.StepAuth,
		Account:   "demo1@test.example.com",
	})

	alist2, err := inst.AuthService.Handle(ctx, action, alist)
	return inst.logResult(alist2, err)
}

func (inst *TestAuthWithEmail) logResult(alist []*rbac.AuthDTO, err error) error {

	if err != nil {
		return err
	}

	vo := &rbac.AuthVO{
		Auth: alist,
	}

	str, err2 := json.MarshalIndent(vo, "", "\t")
	if err2 != nil {
		return err2
	}

	vlog.Info("auth.result: %s", str)
	return nil
}
