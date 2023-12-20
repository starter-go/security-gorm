package impldao

import (
	"fmt"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

// RegionDaoImpl ...
type RegionDaoImpl struct {

	//starter:component
	_as func(rbacdb.RegionDAO) //starter:as("#")

	Agent       rbacdb.LocalAgent  //starter:inject("#")
	UUIDService random.UUIDService //starter:inject("#")
}

func (inst *RegionDaoImpl) _impl(a rbacdb.RegionDAO) {
	a = inst
}

func (inst *RegionDaoImpl) model() *rbacdb.RegionEntity {
	return &rbacdb.RegionEntity{}
}

func (inst *RegionDaoImpl) modelList() []*rbacdb.RegionEntity {
	return make([]*rbacdb.RegionEntity, 0)
}

func (inst *RegionDaoImpl) makeResult(ent *rbacdb.RegionEntity, res *gorm.DB) (*rbacdb.RegionEntity, error) {
	err := res.Error
	rows := res.RowsAffected
	if err != nil {
		return nil, err
	}
	if rows != 1 {
		return nil, fmt.Errorf("db(result).RowsAffected  is %d", rows)
	}
	if ent == nil {
		return nil, fmt.Errorf("the result entity is nil")
	}
	return ent, nil
}

// Insert ...
func (inst *RegionDaoImpl) Insert(db *gorm.DB, o *rbacdb.RegionEntity) (*rbacdb.RegionEntity, error) {

	o.ID = 0
	o.UUID = inst.UUIDService.Build().Class("rbacdb.RegionEntity").Generate()

	db = inst.Agent.DB(db)
	res := db.Create(o)
	return inst.makeResult(o, res)
}

// Update ...
func (inst *RegionDaoImpl) Update(db *gorm.DB, id rbac.RegionID, updater func(*rbacdb.RegionEntity)) (*rbacdb.RegionEntity, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.First(m, id)
	if res.Error == nil {
		updater(m)
		res = db.Save(m)
	}
	return inst.makeResult(m, res)
}

// Delete ...
func (inst *RegionDaoImpl) Delete(db *gorm.DB, id rbac.RegionID) error {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.Delete(m, id)
	return res.Error
}

// Find ...
func (inst *RegionDaoImpl) Find(db *gorm.DB, id rbac.RegionID) (*rbacdb.RegionEntity, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.First(m, id)
	return inst.makeResult(m, res)
}

// List ...
func (inst *RegionDaoImpl) List(db *gorm.DB, q *rbac.RegionQuery) ([]*rbacdb.RegionEntity, error) {

	if q == nil {
		q = &rbac.RegionQuery{}
	}
	list := inst.modelList()
	item := inst.model()

	f := finder{}
	f.all = q.All
	f.listModel = &list
	f.itemModel = item
	f.page = &q.Pagination
	f.conditions = &q.Conditions
	f.db = inst.Agent.DB(db)

	err := f.find()
	if err != nil {
		return nil, err
	}
	return list, nil
}

// ListAll ...
func (inst *RegionDaoImpl) ListAll(db *gorm.DB) ([]*rbacdb.RegionEntity, error) {
	db = inst.Agent.DB(db)
	list := inst.modelList()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}
