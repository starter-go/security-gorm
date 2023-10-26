package rbacdb

import (
	"context"

	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// RoleEntity 表示 Role 的数据库实体
type RoleEntity struct {
	ID rbac.RoleID `gorm:"primaryKey"`

	BaseEntity

	Name rbac.RoleName `gorm:"unique"`
}

// RoleDAO 是数据库访问对象
type RoleDAO interface {
	Insert(db *gorm.DB, o *RoleEntity) (*RoleEntity, error)
	Update(db *gorm.DB, id rbac.RoleID, updater func(*RoleEntity)) (*RoleEntity, error)
	Delete(db *gorm.DB, id rbac.RoleID) error

	Find(db *gorm.DB, id rbac.RoleID) (*RoleEntity, error)
	List(db *gorm.DB, q *rbac.RoleQuery) ([]*RoleEntity, error)
	ListAll(db *gorm.DB) ([]*RoleEntity, error)
}

// RoleConvertor 负责 dto <==> entity 的转换
type RoleConvertor interface {
	ConvertE2D(c context.Context, entity *RoleEntity) (*rbac.RoleDTO, error)
	ConvertD2E(c context.Context, dto *rbac.RoleDTO) (*RoleEntity, error)

	ConvertListE2D(c context.Context, entity []*RoleEntity) ([]*rbac.RoleDTO, error)
}
