package password

import (
	"encoding/base64"
	"fmt"

	"github.com/starter-go/application"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/auth"
	"github.com/starter-go/security/random"
	"github.com/starter-go/security/rbac"
)

// RootAuth 提供 username(root)+password(随机生成) 的验证机制
type RootAuth struct {
	//starter:component
	_as func(application.Lifecycle, auth.Registry) //starter:as(".",".")

	Enabled bool           //starter:inject("${security.auth-as-root.enabled}")
	Random  random.Service //starter:inject("#")

	salt lang.Hex
	sum  lang.Hex
	t1   lang.Time // 有效期（起）
	t2   lang.Time // 有效期（止）
}

func (inst *RootAuth) _impl(c application.Lifecycle, a auth.Authenticator, b auth.Registry, d auth.Mechanism) {
	a = inst
	b = inst
	c = inst
	d = inst
}

// Life ...
func (inst *RootAuth) Life() *application.Life {
	return &application.Life{
		OnCreate: inst.init,
	}
}

// ListRegistrations ...
func (inst *RootAuth) ListRegistrations() []*auth.Registration {
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
func (inst *RootAuth) Support(r1 auth.Request) bool {
	now := lang.Now()
	if (now < inst.t1) || (inst.t2 < now) {
		return false
	}
	a1, ok := r1.(auth.Authentication)
	if !ok {
		return false
	}
	account := a1.Account()
	mech := a1.Mechanism()
	return (account == "root") && (mech == auth.MechanismPassword)
}

// Authenticate ...
func (inst *RootAuth) Authenticate(a auth.Authentication) ([]auth.Identity, error) {
	account := a.Account()
	password := a.Secret()
	if account != "root" {
		return nil, fmt.Errorf("for user 'root' only")
	}
	pc := auth.PasswordCalculator{}
	pc.Reset().WithSalt(inst.salt.Bytes()).WithTarget(inst.sum.Bytes())
	err := pc.Verify(password)
	if err != nil {
		return nil, err
	}
	return inst.makeIDs()
}

func (inst *RootAuth) makeIDs() ([]auth.Identity, error) {
	mechanism := auth.MechanismPassword
	user := &rbac.UserDTO{
		Roles:    rbac.RoleNameList(rbac.RoleRoot),
		Name:     "root",
		NickName: "root",
	}
	id := auth.NewUserIdentity(mechanism, user)
	ids := []auth.Identity{id}
	return ids, nil
}

func (inst *RootAuth) makeSalt() []byte {
	return inst.Random.NextBytes(32)
}

func (inst *RootAuth) makeTempPassword() string {
	size := 28
	data := inst.Random.NextBytes(size)
	str := base64.StdEncoding.EncodeToString(data)
	return str[0:size]
}

func (inst *RootAuth) initTempPassword() {
	pass := inst.makeTempPassword()
	salt := inst.makeSalt()
	pc := auth.PasswordCalculator{}
	pc.Reset().WithSalt(salt).WithTarget(nil)
	sum := pc.Compute([]byte(pass))
	fmt.Printf("root account (temporary) password:%s max-age:1h\n", pass)
	inst.sum = lang.HexFromBytes(sum)
	inst.salt = lang.HexFromBytes(salt)
}

func (inst *RootAuth) init() error {
	if inst.Enabled {
		const maxAge = 3600 * 1000 // 1h
		now := lang.Now()
		inst.t1 = now
		inst.t2 = now + maxAge
		inst.initTempPassword()
	}
	return nil
}
