package rbacdb

import (
	"context"

	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// PhoneNumberEntity ...
type PhoneNumberEntity struct {
	ID rbac.PhoneNumberID `gorm:"primaryKey"`

	BaseEntity

	FullNumber rbac.FullPhoneNumber `gorm:"unique"`

	Region       rbac.RegionPhoneCode
	SimpleNumber rbac.SimplePhoneNumber
	Verified     bool
}

// PhoneNumberDAO ...
type PhoneNumberDAO interface {
	Insert(db *gorm.DB, o *PhoneNumberEntity) (*PhoneNumberEntity, error)
	Update(db *gorm.DB, id rbac.PhoneNumberID, updater func(*PhoneNumberEntity)) (*PhoneNumberEntity, error)
	Delete(db *gorm.DB, id rbac.PhoneNumberID) error

	Find(db *gorm.DB, id rbac.PhoneNumberID) (*PhoneNumberEntity, error)
	FindByNumber(db *gorm.DB, num rbac.FullPhoneNumber) (*PhoneNumberEntity, error)
	List(db *gorm.DB, q *rbac.PhoneNumberQuery) ([]*PhoneNumberEntity, error)
}

// PhoneNumberConvertor ...
type PhoneNumberConvertor interface {
	ConvertE2D(c context.Context, entity *PhoneNumberEntity) (*rbac.PhoneNumberDTO, error)
	ConvertD2E(c context.Context, dto *rbac.PhoneNumberDTO) (*PhoneNumberEntity, error)

	ConvertListE2D(c context.Context, entity []*PhoneNumberEntity) ([]*rbac.PhoneNumberDTO, error)
}
