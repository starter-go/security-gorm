package impldao

import (
	"fmt"

	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/security/random"
	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// RoleDaoImpl ...
type RoleDaoImpl struct {

	//starter:component
	_as func(rbacdb.RoleDAO) //starter:as("#")

	Agent       rbacdb.LocalAgent  //starter:inject("#")
	UUIDService random.UUIDService //starter:inject("#")
}

func (inst *RoleDaoImpl) _impl(a rbacdb.RoleDAO) {
	a = inst
}

func (inst *RoleDaoImpl) model() *rbacdb.RoleEntity {
	return &rbacdb.RoleEntity{}
}

func (inst *RoleDaoImpl) modelList() []*rbacdb.RoleEntity {
	return make([]*rbacdb.RoleEntity, 0)
}

func (inst *RoleDaoImpl) makeResult(ent *rbacdb.RoleEntity, res *gorm.DB) (*rbacdb.RoleEntity, error) {
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
func (inst *RoleDaoImpl) Insert(db *gorm.DB, o *rbacdb.RoleEntity) (*rbacdb.RoleEntity, error) {

	o.ID = 0
	o.UUID = inst.UUIDService.Build().Class("rbacdb.RoleEntity").Generate()

	db = inst.Agent.DB(db)
	res := db.Create(o)
	return inst.makeResult(o, res)
}

// Update ...
func (inst *RoleDaoImpl) Update(db *gorm.DB, id rbac.RoleID, updater func(*rbacdb.RoleEntity)) (*rbacdb.RoleEntity, error) {
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
func (inst *RoleDaoImpl) Delete(db *gorm.DB, id rbac.RoleID) error {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.Delete(m, id)
	return res.Error
}

// Find ...
func (inst *RoleDaoImpl) Find(db *gorm.DB, id rbac.RoleID) (*rbacdb.RoleEntity, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.First(m, id)
	return inst.makeResult(m, res)
}

// List ...
func (inst *RoleDaoImpl) List(db *gorm.DB, q *rbac.RoleQuery) ([]*rbacdb.RoleEntity, error) {

	if q == nil {
		q = &rbac.RoleQuery{}
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
func (inst *RoleDaoImpl) ListAll(db *gorm.DB) ([]*rbacdb.RoleEntity, error) {
	db = inst.Agent.DB(db)
	list := inst.modelList()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}
