package implconvertor

import (
	"context"
	"fmt"

	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/security/rbac"
)

// PermissionConvertorImpl ...
type PermissionConvertorImpl struct {
	//starter:component
	_as func(rbacdb.PermissionConvertor) //starter:as("#")
}

func (inst *PermissionConvertorImpl) _impl() {
	inst._as(inst)
}

// ConvertE2D ...
func (inst *PermissionConvertorImpl) ConvertE2D(c context.Context, entity *rbacdb.PermissionEntity) (*rbac.PermissionDTO, error) {
	return nil, fmt.Errorf("no impl")
}

// ConvertD2E ...
func (inst *PermissionConvertorImpl) ConvertD2E(c context.Context, dto *rbac.PermissionDTO) (*rbacdb.PermissionEntity, error) {
	return nil, fmt.Errorf("no impl")
}

// ConvertListE2D ...
func (inst *PermissionConvertorImpl) ConvertListE2D(c context.Context, entity []*rbacdb.PermissionEntity) ([]*rbac.PermissionDTO, error) {
	return nil, fmt.Errorf("no impl")
}
