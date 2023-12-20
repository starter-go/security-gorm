package implconvertor

import (
	"context"
	"strings"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/rbacdb"
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
func (inst *RoleConvertorImpl) ConvertE2D(c context.Context, o1 *rbacdb.RoleEntity) (*rbac.RoleDTO, error) {
	o2 := &rbac.RoleDTO{}
	rbacdb.CopyBaseFieldsFromEntityToDTO(&o1.BaseEntity, &o2.BaseDTO)
	o2.ID = o1.ID
	o2.Name = o1.Name
	o2.Description = o1.Description
	return o2, nil
}

// ConvertD2E ...
func (inst *RoleConvertorImpl) ConvertD2E(c context.Context, o1 *rbac.RoleDTO) (*rbacdb.RoleEntity, error) {

	name := o1.Name.String()
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)

	o2 := &rbacdb.RoleEntity{}
	rbacdb.CopyBaseFieldsFromDtoToEntity(&o1.BaseDTO, &o2.BaseEntity)
	o2.ID = o1.ID
	o2.Name = rbac.RoleName(name)
	o2.Description = o1.Description
	return o2, nil
}

// ConvertListE2D ...
func (inst *RoleConvertorImpl) ConvertListE2D(c context.Context, src []*rbacdb.RoleEntity) ([]*rbac.RoleDTO, error) {
	dst := make([]*rbac.RoleDTO, 0)
	for _, o1 := range src {
		o2, err := inst.ConvertE2D(c, o1)
		if err != nil {
			return nil, err
		}
		dst = append(dst, o2)
	}
	return dst, nil
}
