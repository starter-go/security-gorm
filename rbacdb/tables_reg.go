package rbacdb

import "github.com/starter-go/libgorm"

// ThisGroupRegistry ... 注册这一组(rbac.db)数据表
type ThisGroupRegistry struct {

	//starter:component
	_as func(libgorm.GroupRegistry) //starter:as(".")

	Enabled bool   //starter:inject("${datagroup.rbac.enabled}")
	Prefix  string //starter:inject("${datagroup.rbac.table-name-prefix}")
	Source  string //starter:inject("${datagroup.rbac.datasource}")

}

func (inst *ThisGroupRegistry) _impl(a libgorm.GroupRegistry) {
	a = inst
}

// Groups ...
func (inst *ThisGroupRegistry) Groups() []*libgorm.GroupRegistration {
	g1 := getGroup(inst.Prefix)
	r1 := &libgorm.GroupRegistration{
		Alias:   "rbac",
		Enabled: inst.Enabled,
		Prefix:  inst.Prefix,
		Source:  inst.Source,
		URI:     "uri:datagroup:rbac",
		Group:   g1,
	}
	return []*libgorm.GroupRegistration{r1}
}

////////////////////////////////////////////////////////////////////////////////

// Group 导出数据表分组（rbac.db）对象
func getGroup(prefix string) libgorm.Group {
	if (prefix != "") && (theTableNamePrefix == "") {
		theTableNamePrefix = prefix
	}
	return &group{}
}

type group struct {
}

func (inst *group) Prototypes() []any {
	return entities()
}

////////////////////////////////////////////////////////////////////////////////
