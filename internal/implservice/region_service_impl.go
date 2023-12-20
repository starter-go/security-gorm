package implservice

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/rbacdb"
)

// RegionServiceImpl ...
type RegionServiceImpl struct {

	//starter:component
	_as func(rbac.RegionService) //starter:as("#")

	RegionDao       rbacdb.RegionDAO       //starter:inject("#")
	RegionConvertor rbacdb.RegionConvertor //starter:inject("#")
}

func (inst *RegionServiceImpl) _impl() {
	inst._as(inst)
}

// Insert ...
func (inst *RegionServiceImpl) Insert(c context.Context, o *rbac.RegionDTO) (*rbac.RegionDTO, error) {
	o2, err := inst.RegionConvertor.ConvertD2E(c, o)
	if err != nil {
		return nil, err
	}
	o3, err := inst.RegionDao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}
	return inst.RegionConvertor.ConvertE2D(c, o3)
}

// Update ...
func (inst *RegionServiceImpl) Update(c context.Context, id rbac.RegionID, o *rbac.RegionDTO) (*rbac.RegionDTO, error) {
	o3, err := inst.RegionDao.Update(nil, id, func(ent *rbacdb.RegionEntity) {
		// todo ...

		ent.FlagURL = o.FlagURL
		ent.DisplayName = o.DisplayName
		ent.FullName = o.FullName
		ent.SimpleName = o.SimpleName
		ent.Code2 = o.Code2
		ent.Code3 = o.Code3
		ent.PhoneCode = o.PhoneCode

	})
	if err != nil {
		return nil, err
	}
	return inst.RegionConvertor.ConvertE2D(c, o3)
}

// Delete ...
func (inst *RegionServiceImpl) Delete(c context.Context, id rbac.RegionID) error {
	return inst.RegionDao.Delete(nil, id)
}

// Find ...
func (inst *RegionServiceImpl) Find(c context.Context, id rbac.RegionID) (*rbac.RegionDTO, error) {
	ent, err := inst.RegionDao.Find(nil, id)
	if err != nil {
		return nil, err
	}
	return inst.RegionConvertor.ConvertE2D(c, ent)
}

// List ...
func (inst *RegionServiceImpl) List(c context.Context, q *rbac.RegionQuery) ([]*rbac.RegionDTO, error) {
	list, err := inst.RegionDao.List(nil, q)
	if err != nil {
		return nil, err
	}
	return inst.RegionConvertor.ConvertListE2D(c, list)
}
