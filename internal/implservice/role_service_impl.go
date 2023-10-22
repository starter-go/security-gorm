package implservice

import (
	"context"

	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/security/rbac"
)

// RoleServiceImpl ...
type RoleServiceImpl struct {

	//starter:component
	_as func(rbac.RoleService) //starter:as("#")

	RoleDao       rbacdb.RoleDAO       //starter:inject("#")
	RoleConvertor rbacdb.RoleConvertor //starter:inject("#")
}

func (inst *RoleServiceImpl) _impl() {
	inst._as(inst)
}

// Insert ...
func (inst *RoleServiceImpl) Insert(c context.Context, o *rbac.RoleDTO) (*rbac.RoleDTO, error) {
	o2, err := inst.RoleConvertor.ConvertD2E(c, o)
	if err != nil {
		return nil, err
	}
	o3, err := inst.RoleDao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}
	return inst.RoleConvertor.ConvertE2D(c, o3)
}

// Update ...
func (inst *RoleServiceImpl) Update(c context.Context, id rbac.RoleID, o *rbac.RoleDTO) (*rbac.RoleDTO, error) {
	o3, err := inst.RoleDao.Update(nil, id, func(ent *rbacdb.RoleEntity) {
		// todo ...
		ent.Name = o.Name
	})
	if err != nil {
		return nil, err
	}
	return inst.RoleConvertor.ConvertE2D(c, o3)
}

// Delete ...
func (inst *RoleServiceImpl) Delete(c context.Context, id rbac.RoleID) error {
	return inst.RoleDao.Delete(nil, id)
}

// Find ...
func (inst *RoleServiceImpl) Find(c context.Context, id rbac.RoleID) (*rbac.RoleDTO, error) {
	ent, err := inst.RoleDao.Find(nil, id)
	if err != nil {
		return nil, err
	}
	return inst.RoleConvertor.ConvertE2D(c, ent)
}

// List ...
func (inst *RoleServiceImpl) List(c context.Context, q *rbac.RoleQuery) ([]*rbac.RoleDTO, error) {
	list, err := inst.RoleDao.List(nil, q)
	if err != nil {
		return nil, err
	}
	return inst.RoleConvertor.ConvertListE2D(c, list)
}
