package impldao

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/rbacdb"
	"gorm.io/gorm"
)

// use: rbacdb.NewFinder()

type finder struct {
	db         *gorm.DB
	page       *rbac.Pagination
	conditions *rbac.Conditions
	itemModel  any
	listModel  any
	all        bool // list-all-items
}

func (inst *finder) find() error {
	f := rbacdb.NewFinder()
	f.Init(inst.db)
	f.WithConditions(inst.conditions).WithPagination(inst.page)
	f.AsEntity(inst.itemModel).AsList(inst.listModel)
	f.All(inst.all)
	return f.Find()
}
