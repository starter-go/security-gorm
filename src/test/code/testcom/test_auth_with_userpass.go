package testcom

import (
	"context"

	"github.com/starter-go/application"
	"github.com/starter-go/security/auth"
	"github.com/starter-go/security/rbac"
)

// TestAuthWithUserPassword ... 测试身份验证
type TestAuthWithUserPassword struct {

	//starter:component
	_as func(application.Lifecycle) //starter:as(".")

	AC              application.Context //starter:inject("context")
	RbacAuthService rbac.AuthService    //starter:inject("#")

	Enabled bool //starter:inject("${testcase.enable.auth-with-password}")

}

// Life ...
func (inst *TestAuthWithUserPassword) Life() *application.Life {
	l := &application.Life{}
	if inst.Enabled {
		l.OnLoop = inst.run
	}
	return l
}

func (inst *TestAuthWithUserPassword) run() error {

	ctx := context.Background()
	au := []*rbac.AuthDTO{}

	au = append(au, &rbac.AuthDTO{
		Mechanism: auth.MechanismPassword,
		Account:   "foo@bar",
		Secret:    "666",
	})

	return inst.RbacAuthService.Handle(ctx, au)
}
