package securitygorm

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/libgorm/modules/libgorm"
	modulegormmysql "github.com/starter-go/module-gorm-mysql"
	modulegormsqlserver "github.com/starter-go/module-gorm-sqlserver"
	"github.com/starter-go/security/modules/security"
)

const (
	theModuleName        = "github.com/starter-go/security-gorm"
	theModuleVersion     = "v0.0.17"
	theModuleRevision    = 17
	theModuleResPath     = "src/main/resources"
	theTestModuleResPath = "src/test/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

// ModuleT ...
func ModuleT() *application.ModuleBuilder {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theModuleResFS, theModuleResPath)

	mb.Depend(security.Module())
	mb.Depend(libgorm.Module())
	mb.Depend(modulegormsqlserver.Module())
	mb.Depend(modulegormmysql.Module())

	return mb
}

// TestModuleT ...
func TestModuleT() *application.ModuleBuilder {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)
	return mb
}
