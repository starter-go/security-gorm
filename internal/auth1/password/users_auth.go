package password

import (
	"context"

	"github.com/starter-go/application"
	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/security/auth"
	"github.com/starter-go/security/random"
)

// UsersAuth 提供 username+password 的验证机制
type UsersAuth struct {

	//starter:component
	_as func(application.Lifecycle, auth.Registry) //starter:as(".",".")

	Enabled bool                 //starter:inject("${security.auth-as-root.enabled}")
	Random  random.Service       //starter:inject("#")
	UserDao rbacdb.UserDAO       //starter:inject("#")
	UserCvt rbacdb.UserConvertor //starter:inject("#")

}

func (inst *UsersAuth) _impl(c application.Lifecycle, a auth.Authenticator, b auth.Registry, d auth.Mechanism) {
	a = inst
	b = inst
	c = inst
	d = inst
}

// Life ...
func (inst *UsersAuth) Life() *application.Life {
	return &application.Life{
		// OnCreate: inst.init,
	}
}

// ListRegistrations ...
func (inst *UsersAuth) ListRegistrations() []*auth.Registration {
	r1 := &auth.Registration{
		Authenticator: inst,
		Authorizer:    nil,
		Mechanism:     inst,
		Enabled:       inst.Enabled,
		Priority:      0,
	}
	list := []*auth.Registration{}
	if inst.Enabled {
		list = append(list, r1)
	}
	return list
}

// Support ...
func (inst *UsersAuth) Support(r1 auth.Request) bool {
	a1, ok := r1.(auth.Authentication)
	if !ok {
		return false
	}
	mech := a1.Mechanism()
	account := a1.Account()
	if account == "root" {
		return false // 对于 root 用户, 必须使用单独的验证机制
	}
	return (mech == auth.MechanismPassword)
}

// Authenticate ...
func (inst *UsersAuth) Authenticate(a1 auth.Authentication) ([]auth.Identity, error) {

	account := a1.Account() // username|phone|email
	password := a1.Secret()

	user, err := inst.UserDao.FindByAny(nil, account)
	if err != nil {
		return nil, err
	}

	pc := auth.PasswordCalculator{}
	pc.Reset().WithSalt(user.Salt.Bytes()).WithTarget(user.Password.Bytes())
	err = pc.Verify(password)
	if err != nil {
		return nil, err
	}

	return inst.makeIDs(a1, user)
}

func (inst *UsersAuth) makeIDs(a1 auth.Authentication, user1 *rbacdb.UserEntity) ([]auth.Identity, error) {
	ctx := context.Background()
	user2, err := inst.UserCvt.ConvertE2D(ctx, user1)
	if err != nil {
		return nil, err
	}
	id := auth.NewUserIdentity(a1, user2)
	ids := []auth.Identity{id}
	return ids, nil
}
