package securitygorm

import (
	"github.com/starter-go/application"
	"github.com/starter-go/module-gorm-mysql/modules/mysql"
	"github.com/starter-go/module-gorm-sqlserver/modules/sqlserver"
	securitygorm1 "github.com/starter-go/security-gorm"
	"github.com/starter-go/security-gorm/gen/test4securitygorm"
	"github.com/starter-go/units/modules/units"
)

// ModuleForTest 导出模块 [github.com/starter-go/security-gorm#test]
func ModuleForTest() application.Module {
	mb := securitygorm1.NewTestModule()
	mb.Components(test4securitygorm.ComForSecurityGormTest)

	mb.Depend(Module())
	mb.Depend(units.Module())
	mb.Depend(sqlserver.Module())
	mb.Depend(mysql.Module())

	return mb.Create()
}
