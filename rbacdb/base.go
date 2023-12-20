package rbacdb

import (
	"time"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"gorm.io/gorm"
)

// BaseEntity 是基本的数据库实体
type BaseEntity struct {
	UUID lang.UUID `gorm:"unique"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Owner   rbac.UserID
	Creator rbac.UserID
	Updater rbac.UserID
}
