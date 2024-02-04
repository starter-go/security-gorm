package implservice

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/rbacdb"
)

// GroupServiceImpl ...
type GroupServiceImpl struct {

	//starter:component
	_as func(rbac.GroupService) //starter:as("#")

	GroupDao       rbacdb.GroupDAO       //starter:inject("#")
	GroupConvertor rbacdb.GroupConvertor //starter:inject("#")
}

func (inst *GroupServiceImpl) _impl() {
	inst._as(inst)
}

// Insert ...
func (inst *GroupServiceImpl) Insert(c context.Context, o *rbac.GroupDTO) (*rbac.GroupDTO, error) {
	o2, err := inst.GroupConvertor.ConvertD2E(c, o)
	if err != nil {
		return nil, err
	}
	o3, err := inst.GroupDao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}
	return inst.GroupConvertor.ConvertE2D(c, o3)
}

// Update ...
func (inst *GroupServiceImpl) Update(c context.Context, id rbac.GroupID, o *rbac.GroupDTO) (*rbac.GroupDTO, error) {
	o3, err := inst.GroupDao.Update(nil, id, func(ent *rbacdb.GroupEntity) {
		// todo ...
		ent.Label = o.Label
		ent.Description = o.Description
		ent.Roles = o.Roles
		ent.Enabled = o.Enabled
	})
	if err != nil {
		return nil, err
	}
	return inst.GroupConvertor.ConvertE2D(c, o3)
}

// Delete ...
func (inst *GroupServiceImpl) Delete(c context.Context, id rbac.GroupID) error {
	return inst.GroupDao.Delete(nil, id)
}

// Find ...
func (inst *GroupServiceImpl) Find(c context.Context, id rbac.GroupID) (*rbac.GroupDTO, error) {
	ent, err := inst.GroupDao.Find(nil, id)
	if err != nil {
		return nil, err
	}
	return inst.GroupConvertor.ConvertE2D(c, ent)
}

// List ...
func (inst *GroupServiceImpl) List(c context.Context, q *rbac.GroupQuery) ([]*rbac.GroupDTO, error) {
	list, err := inst.GroupDao.List(nil, q)
	if err != nil {
		return nil, err
	}
	return inst.GroupConvertor.ConvertListE2D(c, list)
}
