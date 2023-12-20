package implconvertor

import (
	"context"
	"strings"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/rbacdb"
)

// RegionConvertorImpl ...
type RegionConvertorImpl struct {
	//starter:component
	_as func(rbacdb.RegionConvertor) //starter:as("#")
}

func (inst *RegionConvertorImpl) _impl() {
	inst._as(inst)
}

// ConvertE2D ...
func (inst *RegionConvertorImpl) ConvertE2D(c context.Context, o1 *rbacdb.RegionEntity) (*rbac.RegionDTO, error) {
	o2 := &rbac.RegionDTO{}
	rbacdb.CopyBaseFieldsFromEntityToDTO(&o1.BaseEntity, &o2.BaseDTO)
	o2.ID = o1.ID

	o2.FlagURL = o1.FlagURL
	o2.DisplayName = o1.DisplayName
	o2.SimpleName = o1.SimpleName
	o2.Code2 = o1.Code2
	o2.Code3 = o1.Code3
	o2.PhoneCode = o1.PhoneCode
	o2.FullName = o1.FullName

	return o2, nil
}

// ConvertD2E ...
func (inst *RegionConvertorImpl) ConvertD2E(c context.Context, o1 *rbac.RegionDTO) (*rbacdb.RegionEntity, error) {
	o2 := &rbacdb.RegionEntity{}
	rbacdb.CopyBaseFieldsFromDtoToEntity(&o1.BaseDTO, &o2.BaseEntity)
	o2.ID = o1.ID

	o2.FlagURL = strings.TrimSpace(o1.FlagURL)
	o2.DisplayName = strings.TrimSpace(o1.DisplayName)
	o2.SimpleName = strings.TrimSpace(o1.SimpleName)
	o2.FullName = strings.TrimSpace(o1.FullName)

	o2.PhoneCode = o1.PhoneCode.Normalize()
	o2.Code2 = o1.Code2.Normalize()
	o2.Code3 = o1.Code3.Normalize()

	return o2, nil
}

// ConvertListE2D ...
func (inst *RegionConvertorImpl) ConvertListE2D(c context.Context, src []*rbacdb.RegionEntity) ([]*rbac.RegionDTO, error) {
	dst := make([]*rbac.RegionDTO, 0)
	for _, o1 := range src {
		o2, err := inst.ConvertE2D(c, o1)
		if err != nil {
			return nil, err
		}
		dst = append(dst, o2)
	}
	return dst, nil
}
