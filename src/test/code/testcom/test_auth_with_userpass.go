package testcom

import (
	"context"

	"github.com/starter-go/application"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/src/test/code/cases"
	"github.com/starter-go/security/auth"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

// TestAuthWithUserPassword ... 测试身份验证
type TestAuthWithUserPassword struct {

	//starter:component
	_as func(units.Units) //starter:as(".")

	AC              application.Context //starter:inject("context")
	RbacAuthService rbac.AuthService    //starter:inject("#")

}

func (inst *TestAuthWithUserPassword) _impl() units.Units {
	return inst
}

// Units ...
func (inst *TestAuthWithUserPassword) Units(list []*units.Registration) []*units.Registration {

	// bl := testboot.BootList{}
	// bl.Handle(http.MethodGet, "/api//", inst.run)
	// return bl.List()

	list = append(list, &units.Registration{
		Name:    cases.LoginWithPassword,
		Enabled: true,
		Test:    inst.run,
	})

	return list
}

func (inst *TestAuthWithUserPassword) run() error {

	ctx := context.Background()
	au := []*rbac.AuthDTO{}

	au = append(au, &rbac.AuthDTO{
		Mechanism: auth.MechanismPassword,
		Account:   "foo@bar",
		Secret:    "666",
	})

	alist2, err := inst.RbacAuthService.Handle(ctx, "test", au)
	if err != nil {
		return err
	}

	for _, item := range alist2 {
		vlog.Debug("%s", item)
	}
	return nil
}
