package securitygorm

import (
	"embed"

	"github.com/starter-go/application"
)

const (
	theModuleName     = "github.com/starter-go/security-gorm"
	theModuleVersion  = "v1.0.54.3"
	theModuleRevision = 27
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

	return mb
}

// NewTestModule 创建模块 [github.com/starter-go/security-gorm#test]
func NewTestModule() *application.ModuleBuilder {

	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)

	return mb
}
