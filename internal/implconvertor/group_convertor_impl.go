package implconvertor

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/rbacdb"
)

// GroupConvertorImpl ...
type GroupConvertorImpl struct {
	//starter:component
	_as func(rbacdb.GroupConvertor) //starter:as("#")
}

func (inst *GroupConvertorImpl) _impl() {
	inst._as(inst)
}

// ConvertE2D ...
func (inst *GroupConvertorImpl) ConvertE2D(c context.Context, o1 *rbacdb.GroupEntity) (*rbac.GroupDTO, error) {
	o2 := &rbac.GroupDTO{}
	rbacdb.CopyBaseFieldsFromEntityToDTO(&o1.BaseEntity, &o2.BaseDTO)
	o2.ID = o1.ID
	o2.Label = o1.Label
	o2.Name = o1.Name
	o2.Description = o1.Description
	o2.Roles = o1.Roles.Normalize()
	o2.Enabled = o1.Enabled
	return o2, nil
}

// ConvertD2E ...
func (inst *GroupConvertorImpl) ConvertD2E(c context.Context, o1 *rbac.GroupDTO) (*rbacdb.GroupEntity, error) {
	o2 := &rbacdb.GroupEntity{}
	rbacdb.CopyBaseFieldsFromDtoToEntity(&o1.BaseDTO, &o2.BaseEntity)
	o2.ID = o1.ID
	o2.Label = o1.Label
	o2.Name = o1.Name
	o2.Description = o1.Description
	o2.Roles = o1.Roles.Normalize()
	o2.Enabled = o1.Enabled
	return o2, nil
}

// ConvertListE2D ...
func (inst *GroupConvertorImpl) ConvertListE2D(c context.Context, src []*rbacdb.GroupEntity) ([]*rbac.GroupDTO, error) {
	dst := make([]*rbac.GroupDTO, 0)
	for _, o1 := range src {
		o2, err := inst.ConvertE2D(c, o1)
		if err == nil {
			dst = append(dst, o2)
		} else {
			return nil, err
		}
	}
	return dst, nil
}
