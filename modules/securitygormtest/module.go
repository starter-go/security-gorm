package securitygormtest

import (
	"github.com/starter-go/application"
	securitygorm1 "github.com/starter-go/security-gorm"
	"github.com/starter-go/security-gorm/gen/gen4securitygormtest"
	securitygorm2 "github.com/starter-go/security-gorm/modules/securitygorm"
)

// ModuleForTest ... 导出模块
func ModuleForTest() application.Module {
	mb := securitygorm1.TestModuleT()
	mb.Components(gen4securitygormtest.ComForSecurityGormTest)
	mb.Depend(securitygorm2.Module())
	return mb.Create()
}
