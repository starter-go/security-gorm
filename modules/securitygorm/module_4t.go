package securitygorm

import (
	"github.com/starter-go/application"
	securitygorm1 "github.com/starter-go/security-gorm"
	"github.com/starter-go/security-gorm/gen/test4securitygorm"
)

// ModuleForTest 导出模块 [github.com/starter-go/security-gorm#test]
func ModuleForTest() application.Module {
	mb := securitygorm1.NewTestModule()
	mb.Components(test4securitygorm.ComForSecurityGormTest)
	mb.Depend(Module())
	return mb.Create()
}
