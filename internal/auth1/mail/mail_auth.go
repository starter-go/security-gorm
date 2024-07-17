package mail

import (
	"strings"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/i18n"
	"github.com/starter-go/keyvalues"
	"github.com/starter-go/mails"
	"github.com/starter-go/mails/templates"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/internal/auth1/vericode"
	"github.com/starter-go/security/auth"
	"github.com/starter-go/security/random"
)

const (
	myStepSendCode = rbac.StepSendCode
	myStepAuth     = rbac.StepAuth
	myStepHelp     = rbac.StepHelp
)

// UserMailAuth ...
type UserMailAuth struct {

	//starter:component

	_as func(auth.Registry) //starter:as(".")

	KVS       keyvalues.Service         //starter:inject("#")
	Mails     mails.Service             //starter:inject("#")
	Random    random.Service            //starter:inject("#")
	Templates templates.TemplateManager //starter:inject("#")

	MaxAgeSec int64  //starter:inject("${security.verification.mail.max-age-sec}")
	TemplName string //starter:inject("${security.verification.mail.template}")

	vcc *vericode.Context
}

func (inst *UserMailAuth) _impl() (auth.Authenticator, auth.Registry, auth.Mechanism) {
	return inst, inst, inst
}

func (inst *UserMailAuth) getVCConfig() *vericode.Config {

	const (
		ns    = "github.com/starter-go/security-gorm/internal/auth1/mail"
		alias = "VeriMail"
	)

	maxage := inst.MaxAgeSec

	cfg := &vericode.Config{
		KVS:              inst.KVS,
		Mails:            inst.Mails,
		Random:           inst.Random,
		MessageTemplates: inst.Templates,

		NS:                  ns,
		Alias:               alias,
		MaxAge:              lang.Milliseconds(maxage * 1000),
		MessageTemplateName: inst.TemplName,
	}

	return cfg
}

func (inst *UserMailAuth) getVCC() (*vericode.Context, error) {

	// try-get
	vcc := inst.vcc
	if vcc != nil {
		return vcc, nil
	}

	// load
	cfg := inst.getVCConfig()
	ctx, err := vericode.New(cfg)
	if err != nil {
		return nil, err
	}
	vcc = ctx

	// hold
	inst.vcc = vcc
	return vcc, nil
}

// ListRegistrations ...
func (inst *UserMailAuth) ListRegistrations() []*auth.Registration {
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
func (inst *UserMailAuth) Support(r auth.Request) bool {
	a := r.(auth.Authentication)
	want := a.Mechanism()
	have := auth.MechanismEmail
	return want == have
}

// Authenticate ...
func (inst *UserMailAuth) Authenticate(a auth.Authentication) ([]auth.Identity, error) {
	step := a.Step()
	switch step {
	case myStepSendCode:
		return inst.doSendCode(a)
	case myStepAuth:
		return inst.doAuth(a)
	default:
		return inst.doHelp(a)
	}
}

func getLanguage(a auth.Authentication) []i18n.Language {
	return vericode.GetLanguage(a)
}

func (inst *UserMailAuth) doSendCode(a auth.Authentication) ([]auth.Identity, error) {

	vcc, err := inst.getVCC()
	if err != nil {
		return nil, err
	}

	langs := getLanguage(a)
	addr := a.Account()
	acc := vcc.ForAccount(addr, langs...)

	ver, err := acc.SendNewCode()
	if err != nil {
		return nil, err
	}

	nonce := lang.Base64FromBytes(ver.Nonce)
	a.Feedback().Parameters().SetParam("nonce", nonce.String())
	a.Feedback().Parameters().SetParam("account", addr)

	ids := make([]auth.Identity, 0)
	return ids, nil
}

func (inst *UserMailAuth) doAuth(a auth.Authentication) ([]auth.Identity, error) {

	vcc, err := inst.getVCC()
	if err != nil {
		return nil, err
	}

	addr := a.Account()
	code := a.Parameters().GetParam("code")
	nonce := a.Parameters().GetParam("nonce")
	acc := vcc.ForAccount(addr)

	nonce64 := lang.Base64(nonce)

	ver := &vericode.Verification{
		Nonce: nonce64.Bytes(),
		Addr:  addr,
		Code:  vericode.Code(code),
	}

	err = acc.Verify(ver)
	if err != nil {
		return nil, err
	}

	// make result ...
	return inst.makeResultWithEmailAddress(a, addr)
}

func (inst *UserMailAuth) doHelp(a auth.Authentication) ([]auth.Identity, error) {
	sep := ""
	builder := &strings.Builder{}

	// build steps
	builder.Reset()
	steps := []string{myStepHelp, myStepSendCode, myStepAuth}
	for _, step := range steps {
		builder.WriteString(sep)
		builder.WriteString(step)
		sep = ", "
	}
	valueSteps := builder.String()

	// build params
	valueParams := "code,nonce,account"

	// put
	a.Feedback().Parameters().SetParam("steps", valueSteps)
	a.Feedback().Parameters().SetParam("params", valueParams)

	ids := []auth.Identity{}
	return ids, nil
}

func (inst *UserMailAuth) makeResultWithEmailAddress(a auth.Authentication, addr string) ([]auth.Identity, error) {

	info := &rbac.EmailAddressDTO{}
	info.Address = rbac.EmailAddress(addr)

	ident1 := auth.NewEmailIdentity(a, info)

	ids := make([]auth.Identity, 0)
	ids = append(ids, ident1)
	return ids, nil
}
