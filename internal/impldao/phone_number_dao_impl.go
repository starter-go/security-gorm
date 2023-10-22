package impldao

import (
	"fmt"

	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/security/random"
	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// PhoneNumberDaoImpl ...
type PhoneNumberDaoImpl struct {

	//starter:component
	_as func(rbacdb.PhoneNumberDAO) //starter:as("#")

	Agent       rbacdb.LocalAgent  //starter:inject("#")
	UUIDService random.UUIDService //starter:inject("#")
}

func (inst *PhoneNumberDaoImpl) _impl(a rbacdb.PhoneNumberDAO) {
	a = inst
}

func (inst *PhoneNumberDaoImpl) model() *rbacdb.PhoneNumberEntity {
	return &rbacdb.PhoneNumberEntity{}
}

func (inst *PhoneNumberDaoImpl) modelList() []*rbacdb.PhoneNumberEntity {
	return make([]*rbacdb.PhoneNumberEntity, 0)
}

func (inst *PhoneNumberDaoImpl) makeResult(ent *rbacdb.PhoneNumberEntity, err error) (*rbacdb.PhoneNumberEntity, error) {
	if err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, fmt.Errorf("the result entity is nil")
	}
	return ent, nil
}

// Insert ...
func (inst *PhoneNumberDaoImpl) Insert(db *gorm.DB, o *rbacdb.PhoneNumberEntity) (*rbacdb.PhoneNumberEntity, error) {

	o.ID = 0
	o.UUID = inst.UUIDService.Build().Class("rbacdb.PhoneNumberEntity").Generate()

	db = inst.Agent.DB(db)
	res := db.Create(o)
	return inst.makeResult(o, res.Error)
}

// Update ...
func (inst *PhoneNumberDaoImpl) Update(db *gorm.DB, id rbac.PhoneNumberID, updater func(*rbacdb.PhoneNumberEntity)) (*rbacdb.PhoneNumberEntity, error) {
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
func (inst *PhoneNumberDaoImpl) Delete(db *gorm.DB, id rbac.PhoneNumberID) error {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.Delete(m, id)
	return res.Error
}

// Find ...
func (inst *PhoneNumberDaoImpl) Find(db *gorm.DB, id rbac.PhoneNumberID) (*rbacdb.PhoneNumberEntity, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.Find(m, id)
	return inst.makeResult(m, res.Error)
}

// FindByNumber ...
func (inst *PhoneNumberDaoImpl) FindByNumber(db *gorm.DB, num rbac.FullPhoneNumber) (*rbacdb.PhoneNumberEntity, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.Where("full_number = ?", num).First(m)
	return inst.makeResult(m, res.Error)
}

// List ...
func (inst *PhoneNumberDaoImpl) List(db *gorm.DB, q *rbac.PhoneNumberQuery) ([]*rbacdb.PhoneNumberEntity, error) {

	db = inst.Agent.DB(db).Model(inst.model())

	// page
	page := q.Pagination
	if page.Size > 0 {

		var cnt int64
		db.Count(&cnt)

		db = db.Offset(int(page.Offset()))
		db = db.Limit(int(page.Limit()))

		q.Pagination.Total = cnt
	}

	// find
	list := inst.modelList()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}
