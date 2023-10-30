package impldao

import (
	"fmt"

	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/security/random"
	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// EmailAddressDaoImpl ...
type EmailAddressDaoImpl struct {

	//starter:component
	_as func(rbacdb.EmailAddressDAO) //starter:as("#")

	Agent       rbacdb.LocalAgent  //starter:inject("#")
	UUIDService random.UUIDService //starter:inject("#")
}

func (inst *EmailAddressDaoImpl) _impl(a rbacdb.EmailAddressDAO) {
	a = inst
}

func (inst *EmailAddressDaoImpl) model() *rbacdb.EmailAddressEntity {
	return &rbacdb.EmailAddressEntity{}
}

func (inst *EmailAddressDaoImpl) modelList() []*rbacdb.EmailAddressEntity {
	return make([]*rbacdb.EmailAddressEntity, 0)
}

func (inst *EmailAddressDaoImpl) makeResult(ent *rbacdb.EmailAddressEntity, err error) (*rbacdb.EmailAddressEntity, error) {
	if err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, fmt.Errorf("the result entity is nil")
	}
	return ent, nil
}

// Insert ...
func (inst *EmailAddressDaoImpl) Insert(db *gorm.DB, o *rbacdb.EmailAddressEntity) (*rbacdb.EmailAddressEntity, error) {

	o.ID = 0
	o.UUID = inst.UUIDService.Build().Class("rbacdb.EmailAddressEntity").Generate()

	db = inst.Agent.DB(db)
	res := db.Create(o)
	return inst.makeResult(o, res.Error)
}

// Update ...
func (inst *EmailAddressDaoImpl) Update(db *gorm.DB, id rbac.EmailAddressID, updater func(*rbacdb.EmailAddressEntity)) (*rbacdb.EmailAddressEntity, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.Find(m, id)
	if res.Error == nil {
		updater(m)
		res = db.Save(m)
	}
	return inst.makeResult(m, res.Error)
}

// Delete ...
func (inst *EmailAddressDaoImpl) Delete(db *gorm.DB, id rbac.EmailAddressID) error {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.Delete(m, id)
	return res.Error
}

// Find ...
func (inst *EmailAddressDaoImpl) Find(db *gorm.DB, id rbac.EmailAddressID) (*rbacdb.EmailAddressEntity, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.Find(m, id)
	return inst.makeResult(m, res.Error)
}

// FindByAddress ...
func (inst *EmailAddressDaoImpl) FindByAddress(db *gorm.DB, address rbac.EmailAddress) (*rbacdb.EmailAddressEntity, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.Where("address = ?", address).First(m)
	return inst.makeResult(m, res.Error)
}

// List ...
func (inst *EmailAddressDaoImpl) List(db *gorm.DB, q *rbac.EmailAddressQuery) ([]*rbacdb.EmailAddressEntity, error) {

	if q == nil {
		q = &rbac.EmailAddressQuery{}
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
