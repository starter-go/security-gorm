package impldao

import (
	"fmt"
	"strings"

	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/security/random"
	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// UserDaoImpl ...
type UserDaoImpl struct {

	//starter:component
	_as func(rbacdb.UserDAO) //starter:as("#")

	Agent           rbacdb.LocalAgent      //starter:inject("#")
	EmailAddressDAO rbacdb.EmailAddressDAO //starter:inject("#")
	PhoneNumberDAO  rbacdb.PhoneNumberDAO  //starter:inject("#")
	UUIDService     random.UUIDService     //starter:inject("#")

}

func (inst *UserDaoImpl) _impl(a rbacdb.UserDAO) {
	a = inst
}

func (inst *UserDaoImpl) model() *rbacdb.UserEntity {
	return &rbacdb.UserEntity{}
}

func (inst *UserDaoImpl) modelList() []*rbacdb.UserEntity {
	return make([]*rbacdb.UserEntity, 0)
}

func (inst *UserDaoImpl) makeResult(ent *rbacdb.UserEntity, res *gorm.DB) (*rbacdb.UserEntity, error) {
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
func (inst *UserDaoImpl) Insert(db *gorm.DB, o *rbacdb.UserEntity) (*rbacdb.UserEntity, error) {

	o.ID = 0
	o.UUID = inst.UUIDService.Build().Class("rbacdb.UserEntity").Generate()

	db = inst.Agent.DB(db)
	res := db.Create(o)
	return inst.makeResult(o, res)
}

// Update ...
func (inst *UserDaoImpl) Update(db *gorm.DB, id rbac.UserID, updater func(*rbacdb.UserEntity)) (*rbacdb.UserEntity, error) {
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
func (inst *UserDaoImpl) Delete(db *gorm.DB, id rbac.UserID) error {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.Delete(m, id)
	return res.Error
}

// Find ...
func (inst *UserDaoImpl) Find(db *gorm.DB, id rbac.UserID) (*rbacdb.UserEntity, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.First(m, id)
	return inst.makeResult(m, res)
}

// List ...
func (inst *UserDaoImpl) List(db *gorm.DB, q *rbac.UserQuery) ([]*rbacdb.UserEntity, error) {

	if q == nil {
		q = &rbac.UserQuery{}
	}
	list := inst.modelList()
	item := inst.model()

	f := finder{}
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

// FindByAny ...
func (inst *UserDaoImpl) FindByAny(db *gorm.DB, text string) (*rbacdb.UserEntity, error) {

	user, err := inst.FindByName(db, rbac.UserName(text))
	if err == nil && user != nil {
		return user, nil
	}

	num, ok := inst.parsePhoneNumber(text)
	if ok {
		user, err := inst.FindByPhone(db, num)
		if err == nil && user != nil {
			return user, nil
		}
	}

	addr, ok := inst.parseEmailAddress(text)
	if ok {
		user, err := inst.FindByEmail(db, addr)
		if err == nil && user != nil {
			return user, nil
		}
	}

	return nil, fmt.Errorf("no user by text:" + text)
}

// FindByName ...
func (inst *UserDaoImpl) FindByName(db *gorm.DB, name rbac.UserName) (*rbacdb.UserEntity, error) {
	mo := inst.model()
	db = inst.Agent.DB(db)
	res := db.Where("name = ?", name).First(mo)
	return inst.makeResult(mo, res)
}

// FindByPhone ...
func (inst *UserDaoImpl) FindByPhone(db *gorm.DB, num rbac.FullPhoneNumber) (*rbacdb.UserEntity, error) {
	ent, err := inst.PhoneNumberDAO.FindByNumber(db, num)
	if err != nil {
		return nil, err
	}
	return inst.Find(db, ent.Owner)
}

// FindByEmail ...
func (inst *UserDaoImpl) FindByEmail(db *gorm.DB, addr rbac.EmailAddress) (*rbacdb.UserEntity, error) {
	ent, err := inst.EmailAddressDAO.FindByAddress(db, addr)
	if err != nil {
		return nil, err
	}
	return inst.Find(db, ent.Owner)
}

func (inst *UserDaoImpl) parseEmailAddress(text string) (rbac.EmailAddress, bool) {
	list := strings.Split(text, "@")
	if len(list) == 2 {
		user := strings.TrimSpace(list[0])
		host := strings.TrimSpace(list[1])
		if user != "" && host != "" {
			if text == user+"@"+host {
				return rbac.EmailAddress(text), true
			}
		}
	}
	return "", false
}

func (inst *UserDaoImpl) parsePhoneNumber(text string) (rbac.FullPhoneNumber, bool) {
	chs := []rune(text)
	b := strings.Builder{}
	for _, ch := range chs {
		if '0' <= ch && ch <= '9' {
			b.WriteRune(ch)
		} else if ch == ' ' || ch == '-' || ch == '+' {
			continue
		} else {
			return "", false
		}
	}
	str := b.String()
	return rbac.FullPhoneNumber(str), true
}
