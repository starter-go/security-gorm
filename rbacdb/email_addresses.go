package rbacdb

import (
	"context"

	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// EmailAddressEntity ...
type EmailAddressEntity struct {
	ID rbac.EmailAddressID

	BaseEntity

	Address  rbac.EmailAddress `gorm:"unique"`
	Verified bool
}

// EmailAddressDAO ...
type EmailAddressDAO interface {
	Insert(db *gorm.DB, o *EmailAddressEntity) (*EmailAddressEntity, error)
	Update(db *gorm.DB, id rbac.EmailAddressID, updater func(*EmailAddressEntity)) (*EmailAddressEntity, error)
	Delete(db *gorm.DB, id rbac.EmailAddressID) error

	Find(db *gorm.DB, id rbac.EmailAddressID) (*EmailAddressEntity, error)
	List(db *gorm.DB, q *rbac.EmailAddressQuery) ([]*EmailAddressEntity, error)

	FindByAddress(db *gorm.DB, address rbac.EmailAddress) (*EmailAddressEntity, error)
}

// EmailAddressConvertor ...
type EmailAddressConvertor interface {
	ConvertE2D(c context.Context, entity *EmailAddressEntity) (*rbac.EmailAddressDTO, error)
	ConvertD2E(c context.Context, dto *rbac.EmailAddressDTO) (*EmailAddressEntity, error)

	ConvertListE2D(c context.Context, entity []*EmailAddressEntity) ([]*rbac.EmailAddressDTO, error)
}
