package impldao

import (
	"fmt"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/security/random"
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
	inst.normalizeEntity(o)
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
		inst.normalizeEntity(m)
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

	if q == nil {
		q = &rbac.PhoneNumberQuery{}
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

func (inst *PhoneNumberDaoImpl) normalizeEntity(ent *rbacdb.PhoneNumberEntity) {
	part1 := ent.Region.String()
	part2 := ent.SimpleNumber.String()
	pure := rbac.PurePhoneNumber(part1 + part2)
	ent.FullNumber = pure.Normalize()
}
