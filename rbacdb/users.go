package rbacdb

import (
	"context"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// UserEntity 表示 User 的数据库实体
type UserEntity struct {
	ID rbac.UserID

	BaseEntity

	Name rbac.UserName `gorm:"unique"` // 用户名

	Nickname string              // 昵称
	Avatar   string              // 头像 (HTTP-URL)
	Phone    rbac.PhoneNumberID  // 主要的手机号
	Email    rbac.EmailAddressID // 主要的 e-mail 地址
	Roles    rbac.RoleNameList   // 用户的角色

	Password lang.Hex // 用户当前的密码
	Salt     lang.Hex // 跟密码相关的盐
}

// UserDAO 是数据库访问对象
type UserDAO interface {
	Insert(db *gorm.DB, o *UserEntity) (*UserEntity, error)
	Update(db *gorm.DB, id rbac.UserID, updater func(o *UserEntity)) (*UserEntity, error)
	Delete(db *gorm.DB, id rbac.UserID) error

	FindByAny(db *gorm.DB, text string) (*UserEntity, error)
	FindByName(db *gorm.DB, name rbac.UserName) (*UserEntity, error)
	FindByPhone(db *gorm.DB, phone rbac.FullPhoneNumber) (*UserEntity, error)
	FindByEmail(db *gorm.DB, email rbac.EmailAddress) (*UserEntity, error)

	Find(db *gorm.DB, id rbac.UserID) (*UserEntity, error)
	List(db *gorm.DB, q *rbac.UserQuery) ([]*UserEntity, error)
}

// UserConvertor 负责 dto <==> entity 的转换
type UserConvertor interface {
	ConvertE2D(c context.Context, entity *UserEntity) (*rbac.UserDTO, error)
	ConvertD2E(c context.Context, dto *rbac.UserDTO) (*UserEntity, error)

	ConvertListE2D(c context.Context, entity []*UserEntity) ([]*rbac.UserDTO, error)
}
