package rbacdb

import (
	"context"

	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// RegionEntity 表示区域的实体
type RegionEntity struct {
	ID rbac.RegionID

	BaseEntity

	DisplayName string
	SimpleName  string
	FullName    string
	FlagURL     string
	Code2       rbac.RegionCode2
	Code3       rbac.RegionCode3
	PhoneCode   rbac.RegionPhoneCode
}

// RegionDAO 是数据库访问对象
type RegionDAO interface {
	Insert(db *gorm.DB, o *RegionEntity) (*RegionEntity, error)
	Update(db *gorm.DB, id rbac.RegionID, updater func(*RegionEntity)) (*RegionEntity, error)
	Delete(db *gorm.DB, id rbac.RegionID) error

	Find(db *gorm.DB, id rbac.RegionID) (*RegionEntity, error)
	List(db *gorm.DB, q *rbac.RegionQuery) ([]*RegionEntity, error)
	ListAll(db *gorm.DB) ([]*RegionEntity, error)
}

// RegionConvertor 负责 dto <==> entity 的转换
type RegionConvertor interface {
	ConvertE2D(c context.Context, entity *RegionEntity) (*rbac.RegionDTO, error)
	ConvertD2E(c context.Context, dto *rbac.RegionDTO) (*RegionEntity, error)

	ConvertListE2D(c context.Context, entity []*RegionEntity) ([]*rbac.RegionDTO, error)
}
