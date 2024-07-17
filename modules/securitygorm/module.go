package securitygorm

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/mails/modules/mails"
	"github.com/starter-go/security-gorm/gen/main4securitygorm"
	"github.com/starter-go/security/modules/security"

	securitygorm "github.com/starter-go/security-gorm"
)

// Module 导出模块 [github.com/starter-go/security-gorm]
func Module() application.Module {
	mb := securitygorm.NewMainModule()
	mb.Components(main4securitygorm.ComForSecurityGorm)

	mb.Depend(security.Module())
	mb.Depend(libgorm.Module())
	mb.Depend(mails.LibModule())

	return mb.Create()
}
