package rbacdb

import (
	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// Finder  提供一套通用的查询工具
type Finder interface {
	Init(db *gorm.DB) Finder
	AsList(list any) Finder
	AsEntity(entity any) Finder
	All(all bool) Finder
	WithPagination(page *rbac.Pagination) Finder
	WithConditions(cond *rbac.Conditions) Finder
	Find() error
}

// NewFinder 新建一个查询工具
func NewFinder() Finder {
	return &finderFacade{}
}

////////////////////////////////////////////////////////////////////////////////

type finderFacade struct {
	inner finder
}

func (inst *finderFacade) Init(db *gorm.DB) Finder {
	inst.inner.db = db
	return inst
}

func (inst *finderFacade) AsList(list any) Finder {
	inst.inner.listModel = list
	return inst
}

func (inst *finderFacade) AsEntity(entity any) Finder {
	inst.inner.itemModel = entity
	return inst
}

func (inst *finderFacade) All(all bool) Finder {
	inst.inner.all = all
	return inst
}

func (inst *finderFacade) WithPagination(page *rbac.Pagination) Finder {
	if page == nil {
		page = &rbac.Pagination{}
	}
	inst.inner.page = page
	return inst
}

func (inst *finderFacade) WithConditions(cond *rbac.Conditions) Finder {
	if cond == nil {
		cond = &rbac.Conditions{}
	}
	inst.inner.conditions = cond
	return inst
}

func (inst *finderFacade) Find() error {
	return inst.inner.find()
}

////////////////////////////////////////////////////////////////////////////////

type finder struct {
	db         *gorm.DB
	page       *rbac.Pagination
	conditions *rbac.Conditions
	itemModel  any
	listModel  any
	all        bool // list-all-items
}

func (inst *finder) find() error {

	list := inst.listModel
	item := inst.itemModel
	db := inst.db

	if inst.all {
		res := db.Find(list)
		return res.Error
	}

	// Conditions
	if inst.hasConditions() {
		q := inst.conditions.Query
		args1 := inst.conditions.Args
		args2 := make([]any, 0)
		for _, item := range args1 {
			args2 = append(args2, item)
		}
		db = db.Where(q, args2...)
	}

	if inst.hasPage() {
		// count
		var count int64
		db.Model(item).Count(&count)
		inst.page.Total = count

		// page
		page := inst.page
		if page.Page < 1 {
			page.Page = 1
		}
		db = db.Limit(int(page.Limit()))
		db = db.Offset(int(page.Offset()))
	}

	// fetch
	res := db.Find(list)
	return res.Error
}

func (inst *finder) hasPage() bool {
	if inst.page == nil {
		return false
	}
	size := inst.page.Size
	if size < 0 {
		return false
	} else if size == 0 {
		inst.page.Size = 10
	} else {
		inst.page.Size = size
	}
	return true
}

func (inst *finder) hasConditions() bool {
	if inst.conditions == nil {
		return false
	}
	q := inst.conditions.Query
	return (q != "")
}
