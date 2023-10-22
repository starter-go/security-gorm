package securitygorm

import (
	"github.com/starter-go/application"
	securitygorm "github.com/starter-go/security-gorm"
	"github.com/starter-go/security-gorm/gen/gen4securitygorm"
)

// Module ... 导出模块
func Module() application.Module {
	mb := securitygorm.ModuleT()
	mb.Components(gen4securitygorm.ComForSecurityGorm)
	return mb.Create()
}
