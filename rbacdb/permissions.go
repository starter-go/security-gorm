package rbacdb

import (
	"context"

	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// PermissionEntity 表示 Permission 的数据库实体
type PermissionEntity struct {
	ID rbac.PermissionID `gorm:"primaryKey"`

	BaseEntity

	Method      string
	Path        string
	Resource    string            `gorm:"unique"` // like 'method + ":" + path'
	AcceptRoles rbac.RoleNameList `gorm:"column:roles"`
}

// PermissionDAO 是数据库访问对象
type PermissionDAO interface {
	Insert(db *gorm.DB, o *PermissionEntity) (*PermissionEntity, error)
	Update(db *gorm.DB, id rbac.PermissionID, updater func(*PermissionEntity)) (*PermissionEntity, error)
	Delete(db *gorm.DB, id rbac.PermissionID) error

	Find(db *gorm.DB, id rbac.PermissionID) (*PermissionEntity, error)
	List(db *gorm.DB, q *rbac.PermissionQuery) ([]*PermissionEntity, error)
	ListAll(db *gorm.DB) ([]*PermissionEntity, error)
}

// PermissionConvertor 负责 dto <==> entity 的转换
type PermissionConvertor interface {
	ConvertE2D(c context.Context, entity *PermissionEntity) (*rbac.PermissionDTO, error)
	ConvertD2E(c context.Context, dto *rbac.PermissionDTO) (*PermissionEntity, error)

	ConvertListE2D(c context.Context, entity []*PermissionEntity) ([]*rbac.PermissionDTO, error)
}
