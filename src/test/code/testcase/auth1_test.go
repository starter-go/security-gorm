package testcase

import (
	"context"
	"testing"

	"github.com/starter-go/security-gorm/src/test/code/testboot"
	"github.com/starter-go/security-gorm/src/test/code/testcom"
	"github.com/starter-go/security/auth"
	"github.com/starter-go/security/rbac"
)

func TestAuth1(t *testing.T) {
	p := &testcom.BootParams{}
	p.Callback = func(b *testcom.Bootstrap) {

		ctx := context.Background()
		a1 := &rbac.AuthDTO{
			Mechanism: auth.MechanismPassword,
			Account:   "root",
			Action:    "login",
			Secret:    "bar",
		}
		err := b.RbacAuthService.Handle(ctx, []*rbac.AuthDTO{a1})
		if err != nil {
			panic(err)
		}

	}
	testboot.Boot(p)
}
