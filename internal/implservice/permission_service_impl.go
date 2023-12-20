package implservice

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/rbacdb"
)

// PermissionServiceImpl ...
type PermissionServiceImpl struct {

	//starter:component
	_as func(rbac.PermissionService) //starter:as("#")

	PermissionDao       rbacdb.PermissionDAO       //starter:inject("#")
	PermissionConvertor rbacdb.PermissionConvertor //starter:inject("#")
	PermissionCache     rbac.PermissionCache       //starter:inject("#")
}

func (inst *PermissionServiceImpl) _impl() {
	inst._as(inst)
}

// GetCache ...
func (inst *PermissionServiceImpl) GetCache() rbac.PermissionCache {
	return inst.PermissionCache
}

// Insert ...
func (inst *PermissionServiceImpl) Insert(c context.Context, o *rbac.PermissionDTO) (*rbac.PermissionDTO, error) {
	o2, err := inst.PermissionConvertor.ConvertD2E(c, o)
	if err != nil {
		return nil, err
	}
	o3, err := inst.PermissionDao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}
	return inst.PermissionConvertor.ConvertE2D(c, o3)
}

// Update ...
func (inst *PermissionServiceImpl) Update(c context.Context, id rbac.PermissionID, o1 *rbac.PermissionDTO) (*rbac.PermissionDTO, error) {
	o2, err := inst.PermissionConvertor.ConvertD2E(c, o1)
	if err != nil {
		return nil, err
	}
	o3, err := inst.PermissionDao.Update(nil, id, func(ent *rbacdb.PermissionEntity) {
		// 不支持修改这些字段：method, path, resource
		// ent.Method = o.Method
		// ent.Path = o.Path
		// ent.Resource = o.Method + ":" + o.Path
		ent.AcceptRoles = o2.AcceptRoles
		ent.Enabled = o2.Enabled
	})
	if err != nil {
		return nil, err
	}
	return inst.PermissionConvertor.ConvertE2D(c, o3)
}

// Delete ...
func (inst *PermissionServiceImpl) Delete(c context.Context, id rbac.PermissionID) error {
	return inst.PermissionDao.Delete(nil, id)
}

// Find ...
func (inst *PermissionServiceImpl) Find(c context.Context, id rbac.PermissionID) (*rbac.PermissionDTO, error) {
	ent, err := inst.PermissionDao.Find(nil, id)
	if err != nil {
		return nil, err
	}
	return inst.PermissionConvertor.ConvertE2D(c, ent)
}

// List ...
func (inst *PermissionServiceImpl) List(c context.Context, q *rbac.PermissionQuery) ([]*rbac.PermissionDTO, error) {
	list, err := inst.PermissionDao.List(nil, q)
	if err != nil {
		return nil, err
	}
	return inst.PermissionConvertor.ConvertListE2D(c, list)
}

// ListAll ...
func (inst *PermissionServiceImpl) ListAll(c context.Context) ([]*rbac.PermissionDTO, error) {
	list, err := inst.PermissionDao.ListAll(nil)
	if err != nil {
		return nil, err
	}
	return inst.PermissionConvertor.ConvertListE2D(c, list)
}
