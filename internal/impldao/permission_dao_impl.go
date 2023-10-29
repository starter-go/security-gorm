package impldao

import (
	"fmt"

	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/security/random"
	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// PermissionDaoImpl ...
type PermissionDaoImpl struct {

	//starter:component
	_as func(rbacdb.PermissionDAO) //starter:as("#")

	Agent       rbacdb.LocalAgent  //starter:inject("#")
	UUIDService random.UUIDService //starter:inject("#")
}

func (inst *PermissionDaoImpl) _impl(a rbacdb.PermissionDAO) {
	a = inst
}

func (inst *PermissionDaoImpl) model() *rbacdb.PermissionEntity {
	return &rbacdb.PermissionEntity{}
}

func (inst *PermissionDaoImpl) modelList() []*rbacdb.PermissionEntity {
	return make([]*rbacdb.PermissionEntity, 0)
}

func (inst *PermissionDaoImpl) makeResult(ent *rbacdb.PermissionEntity, res *gorm.DB) (*rbacdb.PermissionEntity, error) {
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
func (inst *PermissionDaoImpl) Insert(db *gorm.DB, o *rbacdb.PermissionEntity) (*rbacdb.PermissionEntity, error) {

	o.ID = 0
	o.UUID = inst.UUIDService.Build().Class("rbacdb.PermissionEntity").Generate()

	db = inst.Agent.DB(db)
	res := db.Create(o)
	return inst.makeResult(o, res)
}

// Update ...
func (inst *PermissionDaoImpl) Update(db *gorm.DB, id rbac.PermissionID, updater func(*rbacdb.PermissionEntity)) (*rbacdb.PermissionEntity, error) {
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
func (inst *PermissionDaoImpl) Delete(db *gorm.DB, id rbac.PermissionID) error {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.Delete(m, id)
	return res.Error
}

// Find ...
func (inst *PermissionDaoImpl) Find(db *gorm.DB, id rbac.PermissionID) (*rbacdb.PermissionEntity, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.First(m, id)
	return inst.makeResult(m, res)
}

// List ...
func (inst *PermissionDaoImpl) List(db *gorm.DB, q *rbac.PermissionQuery) ([]*rbacdb.PermissionEntity, error) {

	db = inst.Agent.DB(db)

	// query
	if q == nil {
		q = &rbac.PermissionQuery{}
	}

	// page
	page := q.Pagination
	if page.Size > 0 {

		var cnt int64
		db.Model(inst.model()).Count(&cnt)
		q.Pagination.Total = cnt

		db = db.Offset(int(page.Offset()))
		db = db.Limit(int(page.Limit()))
	}

	// find
	list := inst.modelList()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

// ListAll ...
func (inst *PermissionDaoImpl) ListAll(db *gorm.DB) ([]*rbacdb.PermissionEntity, error) {
	db = inst.Agent.DB(db).Model(inst.model())
	list := inst.modelList()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}
