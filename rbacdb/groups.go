package rbacdb

import (
	"context"

	"github.com/starter-go/rbac"
	"gorm.io/gorm"
)

// GroupEntity 表示 Group 的数据库实体
type GroupEntity struct {
	ID rbac.GroupID `gorm:"primaryKey"`

	BaseEntity

	Name rbac.GroupName `gorm:"unique"` // 用户分组名称

	Label       string            // 昵称
	Description string            // 头像 (HTTP-URL)
	Roles       rbac.RoleNameList // 用户的角色
	Enabled     bool              // 启用该项
}

// GroupDAO 是数据库访问对象
type GroupDAO interface {
	Insert(db *gorm.DB, o *GroupEntity) (*GroupEntity, error)
	Update(db *gorm.DB, id rbac.GroupID, updater func(o *GroupEntity)) (*GroupEntity, error)
	Delete(db *gorm.DB, id rbac.GroupID) error

	// FindByAny(db *gorm.DB, text string) (*GroupEntity, error)
	// FindByName(db *gorm.DB, name rbac.GroupName) (*GroupEntity, error)
	// FindByPhone(db *gorm.DB, phone rbac.FullPhoneNumber) (*GroupEntity, error)
	// FindByEmail(db *gorm.DB, email rbac.EmailAddress) (*GroupEntity, error)

	Find(db *gorm.DB, id rbac.GroupID) (*GroupEntity, error)
	List(db *gorm.DB, q *rbac.GroupQuery) ([]*GroupEntity, error)
}

// GroupConvertor 负责 dto <==> entity 的转换
type GroupConvertor interface {
	ConvertE2D(c context.Context, entity *GroupEntity) (*rbac.GroupDTO, error)
	ConvertD2E(c context.Context, dto *rbac.GroupDTO) (*GroupEntity, error)

	ConvertListE2D(c context.Context, entity []*GroupEntity) ([]*rbac.GroupDTO, error)
}
