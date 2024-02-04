package impldao

import (
	"fmt"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

// GroupDaoImpl ...
type GroupDaoImpl struct {

	//starter:component
	_as func(rbacdb.GroupDAO) //starter:as("#")

	Agent       rbacdb.LocalAgent  //starter:inject("#")
	UUIDService random.UUIDService //starter:inject("#")

}

func (inst *GroupDaoImpl) _impl(a rbacdb.GroupDAO) {
	a = inst
}

func (inst *GroupDaoImpl) model() *rbacdb.GroupEntity {
	return &rbacdb.GroupEntity{}
}

func (inst *GroupDaoImpl) modelList() []*rbacdb.GroupEntity {
	return make([]*rbacdb.GroupEntity, 0)
}

func (inst *GroupDaoImpl) makeResult(ent *rbacdb.GroupEntity, res *gorm.DB) (*rbacdb.GroupEntity, error) {
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
func (inst *GroupDaoImpl) Insert(db *gorm.DB, o *rbacdb.GroupEntity) (*rbacdb.GroupEntity, error) {

	o.ID = 0
	o.UUID = inst.UUIDService.Build().Class("rbacdb.GroupEntity").Generate()

	db = inst.Agent.DB(db)
	res := db.Create(o)
	return inst.makeResult(o, res)
}

// Update ...
func (inst *GroupDaoImpl) Update(db *gorm.DB, id rbac.GroupID, updater func(*rbacdb.GroupEntity)) (*rbacdb.GroupEntity, error) {
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
func (inst *GroupDaoImpl) Delete(db *gorm.DB, id rbac.GroupID) error {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.Delete(m, id)
	return res.Error
}

// Find ...
func (inst *GroupDaoImpl) Find(db *gorm.DB, id rbac.GroupID) (*rbacdb.GroupEntity, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.First(m, id)
	return inst.makeResult(m, res)
}

// List ...
func (inst *GroupDaoImpl) List(db *gorm.DB, q *rbac.GroupQuery) ([]*rbacdb.GroupEntity, error) {

	if q == nil {
		q = &rbac.GroupQuery{}
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

// FindByName ...
func (inst *GroupDaoImpl) FindByName(db *gorm.DB, name rbac.GroupName) (*rbacdb.GroupEntity, error) {
	mo := inst.model()
	db = inst.Agent.DB(db)
	res := db.Where("name = ?", name).First(mo)
	return inst.makeResult(mo, res)
}
