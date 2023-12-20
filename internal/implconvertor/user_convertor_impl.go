package implconvertor

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/rbacdb"
)

// UserConvertorImpl ...
type UserConvertorImpl struct {
	//starter:component
	_as func(rbacdb.UserConvertor) //starter:as("#")
}

func (inst *UserConvertorImpl) _impl() {
	inst._as(inst)
}

// ConvertE2D ...
func (inst *UserConvertorImpl) ConvertE2D(c context.Context, o1 *rbacdb.UserEntity) (*rbac.UserDTO, error) {
	o2 := &rbac.UserDTO{}
	rbacdb.CopyBaseFieldsFromEntityToDTO(&o1.BaseEntity, &o2.BaseDTO)
	o2.ID = o1.ID
	o2.Avatar = o1.Avatar
	o2.Name = o1.Name
	o2.NickName = o1.Nickname
	o2.Roles = o1.Roles
	o2.Enabled = o1.Enabled
	return o2, nil
}

// ConvertD2E ...
func (inst *UserConvertorImpl) ConvertD2E(c context.Context, o1 *rbac.UserDTO) (*rbacdb.UserEntity, error) {
	o2 := &rbacdb.UserEntity{}
	rbacdb.CopyBaseFieldsFromDtoToEntity(&o1.BaseDTO, &o2.BaseEntity)
	o2.ID = o1.ID
	o2.Avatar = o1.Avatar
	o2.Name = o1.Name
	o2.Nickname = o1.NickName
	o2.Roles = o1.Roles.Normalize()
	o2.Enabled = o1.Enabled
	return o2, nil
}

// ConvertListE2D ...
func (inst *UserConvertorImpl) ConvertListE2D(c context.Context, src []*rbacdb.UserEntity) ([]*rbac.UserDTO, error) {
	dst := make([]*rbac.UserDTO, 0)
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
