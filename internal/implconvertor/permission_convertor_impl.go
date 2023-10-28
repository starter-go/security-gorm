package implconvertor

import (
	"context"
	"strings"

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

func (inst *PermissionConvertorImpl) computeResourceString(ent *rbacdb.PermissionEntity) {
	method := strings.TrimSpace(ent.Method)
	path := strings.TrimSpace(ent.Path)
	ent.Resource = strings.ToUpper(method) + ":" + path
}

// ConvertE2D ...
func (inst *PermissionConvertorImpl) ConvertE2D(c context.Context, o1 *rbacdb.PermissionEntity) (*rbac.PermissionDTO, error) {
	o2 := &rbac.PermissionDTO{}
	rbacdb.CopyBaseFieldsFromEntityToDTO(&o1.BaseEntity, &o2.BaseDTO)
	o2.ID = o1.ID
	o2.Method = o1.Method
	o2.Path = o1.Path
	o2.AcceptRoles = o1.AcceptRoles
	o2.Enabled = o1.Enabled
	return o2, nil
}

// ConvertD2E ...
func (inst *PermissionConvertorImpl) ConvertD2E(c context.Context, o1 *rbac.PermissionDTO) (*rbacdb.PermissionEntity, error) {
	o2 := &rbacdb.PermissionEntity{}
	rbacdb.CopyBaseFieldsFromDtoToEntity(&o1.BaseDTO, &o2.BaseEntity)
	o2.ID = o1.ID
	o2.Method = strings.ToUpper(o1.Method)
	o2.Path = o1.Path
	o2.AcceptRoles = o1.AcceptRoles.Normalize()
	o2.Enabled = o1.Enabled
	inst.computeResourceString(o2)
	return o2, nil
}

// ConvertListE2D ...
func (inst *PermissionConvertorImpl) ConvertListE2D(c context.Context, src []*rbacdb.PermissionEntity) ([]*rbac.PermissionDTO, error) {
	dst := make([]*rbac.PermissionDTO, 0)
	for _, o1 := range src {
		o2, err := inst.ConvertE2D(c, o1)
		if err != nil {
			return nil, err
		}
		dst = append(dst, o2)
	}
	return dst, nil
}
