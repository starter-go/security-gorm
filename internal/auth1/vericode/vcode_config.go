package vericode

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/keyvalues"
	"github.com/starter-go/mails"
	"github.com/starter-go/mails/templates"
	"github.com/starter-go/security/random"
)

// Config 用于表示 Context 的配置
type Config struct {

	// services
	KVS              keyvalues.Service
	Mails            mails.Service
	Random           random.Service
	MessageTemplates templates.TemplateManager

	// params
	NS                  keyvalues.NS
	Alias               keyvalues.ClassAlias
	MaxAge              lang.Milliseconds
	MessageTemplateName string
}
