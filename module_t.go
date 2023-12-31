package securitygorm

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/module-gorm-mysql/modules/mysql"
	"github.com/starter-go/module-gorm-sqlserver/modules/sqlserver"
	"github.com/starter-go/security/modules/security"
)

const (
	theModuleName     = "github.com/starter-go/security-gorm"
	theModuleVersion  = "v1.0.44"
	theModuleRevision = 18
)

////////////////////////////////////////////////////////////////////////////////

const (
	theMainModuleResPath = "src/main/resources"
	theTestModuleResPath = "src/test/resources"
)

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

// NewMainModule 创建模块 [github.com/starter-go/security-gorm]
func NewMainModule() *application.ModuleBuilder {

	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)

	mb.Depend(security.Module())
	mb.Depend(libgorm.Module())

	return mb
}

// NewTestModule 创建模块 [github.com/starter-go/security-gorm#test]
func NewTestModule() *application.ModuleBuilder {

	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)

	mb.Depend(sqlserver.Module())
	mb.Depend(mysql.Module())

	return mb
}
