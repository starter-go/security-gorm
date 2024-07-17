package vericode

import (
	"time"

	"github.com/starter-go/i18n"
	"github.com/starter-go/keyvalues"
	"github.com/starter-go/mails"
	"github.com/starter-go/mails/templates"
	"github.com/starter-go/security/random"
)

// Context 用于实现 sms|email 等验证码机制
type Context struct {

	// services

	sender     mails.Service
	random     random.Service
	mailsTempl templates.TemplateManager

	// params

	messageTemplateName string
	class               keyvalues.Class
	maxAge              time.Duration
	maxTTL              int
}

// ForAccount ...
func (inst *Context) ForAccount(addr string, langs ...i18n.Language) *Account {
	ent := inst.class.GetJSONEntry(addr)
	acc := &Account{
		vc:        inst,
		addr:      addr,
		entry:     ent,
		languages: langs,
	}
	return acc
}

// NewMessage ...
func (inst *Context) NewMessage(code Code, langs ...i18n.Language) (*mails.Message, error) {

	name := inst.messageTemplateName
	// lang2 := i18n.Language(language)
	templ, err := inst.mailsTempl.Find(name, langs...)
	if err != nil {
		return nil, err
	}

	props := map[string]string{
		"code": string(code),
	}

	return templ.NewMessage(props)

	// old := inst.msgReplacer
	// new := string(code)
	// src := &inst.msgTemplate
	// dst := &mails.Message{}
	// *dst = *src
	// text := string(src.Content)
	// text = strings.ReplaceAll(text, old, new)
	// dst.Content = []byte(text)
	// return dst
}

////////////////////////////////////////////////////////////////////////////////

// New 新建一个 vericode.Context
func New(cfg *Config) (*Context, error) {

	kvs := cfg.KVS
	ns := cfg.NS
	alias := cfg.Alias

	cl, err := kvs.GetClassNS(ns, alias)
	if err != nil {
		return nil, err
	}

	ctx := &Context{
		class:               cl,
		maxAge:              cfg.MaxAge.Duration(),
		sender:              cfg.Mails,
		random:              cfg.Random,
		mailsTempl:          cfg.MessageTemplates,
		messageTemplateName: cfg.MessageTemplateName,
	}
	return ctx, nil
}
