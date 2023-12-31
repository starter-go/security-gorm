package securitygorm

import (
	"github.com/starter-go/application"
	"github.com/starter-go/security-gorm/gen/main4securitygorm"

	securitygorm "github.com/starter-go/security-gorm"
)

// Module 导出模块 [github.com/starter-go/security-gorm]
func Module() application.Module {
	mb := securitygorm.NewMainModule()
	mb.Components(main4securitygorm.ComForSecurityGorm)
	return mb.Create()
}
