package implservice

import (
	"context"

	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/security/rbac"
)

// UserServiceImpl ...
type UserServiceImpl struct {

	//starter:component
	_as func(rbac.UserService) //starter:as("#")

	UserDao       rbacdb.UserDAO       //starter:inject("#")
	UserConvertor rbacdb.UserConvertor //starter:inject("#")
}

func (inst *UserServiceImpl) _impl() {
	inst._as(inst)
}

// Insert ...
func (inst *UserServiceImpl) Insert(c context.Context, o *rbac.UserDTO) (*rbac.UserDTO, error) {
	o2, err := inst.UserConvertor.ConvertD2E(c, o)
	if err != nil {
		return nil, err
	}
	o3, err := inst.UserDao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}
	return inst.UserConvertor.ConvertE2D(c, o3)
}

// Update ...
func (inst *UserServiceImpl) Update(c context.Context, id rbac.UserID, o *rbac.UserDTO) (*rbac.UserDTO, error) {
	o3, err := inst.UserDao.Update(nil, id, func(ent *rbacdb.UserEntity) {
		// todo ...
		ent.Avatar = o.Avatar
		ent.Nickname = o.NickName
	})
	if err != nil {
		return nil, err
	}
	return inst.UserConvertor.ConvertE2D(c, o3)
}

// Delete ...
func (inst *UserServiceImpl) Delete(c context.Context, id rbac.UserID) error {
	return inst.UserDao.Delete(nil, id)
}

// Find ...
func (inst *UserServiceImpl) Find(c context.Context, id rbac.UserID) (*rbac.UserDTO, error) {
	ent, err := inst.UserDao.Find(nil, id)
	if err != nil {
		return nil, err
	}
	return inst.UserConvertor.ConvertE2D(c, ent)
}

// List ...
func (inst *UserServiceImpl) List(c context.Context, q *rbac.UserQuery) ([]*rbac.UserDTO, error) {
	list, err := inst.UserDao.List(nil, q)
	if err != nil {
		return nil, err
	}
	return inst.UserConvertor.ConvertListE2D(c, list)
}
