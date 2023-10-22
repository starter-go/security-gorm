package implconvertor

import (
	"context"
	"fmt"

	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/security/rbac"
)

// RoleConvertorImpl ...
type RoleConvertorImpl struct {
	//starter:component
	_as func(rbacdb.RoleConvertor) //starter:as("#")
}

func (inst *RoleConvertorImpl) _impl() {
	inst._as(inst)
}

// ConvertE2D ...
func (inst *RoleConvertorImpl) ConvertE2D(c context.Context, entity *rbacdb.RoleEntity) (*rbac.RoleDTO, error) {
	return nil, fmt.Errorf("no impl")
}

// ConvertD2E ...
func (inst *RoleConvertorImpl) ConvertD2E(c context.Context, dto *rbac.RoleDTO) (*rbacdb.RoleEntity, error) {
	return nil, fmt.Errorf("no impl")
}

// ConvertListE2D ...
func (inst *RoleConvertorImpl) ConvertListE2D(c context.Context, entity []*rbacdb.RoleEntity) ([]*rbac.RoleDTO, error) {
	return nil, fmt.Errorf("no impl")
}
