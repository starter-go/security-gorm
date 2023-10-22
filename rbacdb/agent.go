package rbacdb

import (
	"github.com/starter-go/libgorm"
	"gorm.io/gorm"
)

// LocalAgent 是 rbac.db 的本地代理
type LocalAgent interface {
	libgorm.Agent
}

////////////////////////////////////////////////////////////////////////////////

// LocalAgentImpl ...
type LocalAgentImpl struct {

	//starter:component
	_as func(LocalAgent) //starter:as("#")

	DataSourceMan   libgorm.DataSourceManager //starter:inject("#")
	DataSourceAlias string                    //starter:inject("${datagroup.rbac.datasource}")

	cache libgorm.DataSourceAgent
}

func (inst *LocalAgentImpl) _impl() {
	inst._as(inst)
}

// DB ...
func (inst *LocalAgentImpl) DB(db *gorm.DB) *gorm.DB {
	c := &inst.cache
	if !c.Ready() {
		c.Init(inst.DataSourceMan, inst.DataSourceAlias)
	}
	return c.DB(db)
}
