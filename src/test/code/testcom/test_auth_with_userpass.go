package testcom

import (
	"context"
	"net/http"

	"github.com/starter-go/application"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/src/test/code/testboot"
	"github.com/starter-go/security/auth"
)

// TestAuthWithUserPassword ... 测试身份验证
type TestAuthWithUserPassword struct {

	//starter:component
	_as func(testboot.BootingRegistry) //starter:as(".")

	AC              application.Context //starter:inject("context")
	RbacAuthService rbac.AuthService    //starter:inject("#")

}

// Boots ...
func (inst *TestAuthWithUserPassword) Boots() []*testboot.Boot {
	bl := testboot.BootList{}
	bl.Handle(http.MethodGet, "/api//", inst.run)
	return bl.List()
}

func (inst *TestAuthWithUserPassword) run() error {

	ctx := context.Background()
	au := []*rbac.AuthDTO{}

	au = append(au, &rbac.AuthDTO{
		Mechanism: auth.MechanismPassword,
		Account:   "foo@bar",
		Secret:    "666",
	})

	return inst.RbacAuthService.Handle(ctx, "test", au)
}
