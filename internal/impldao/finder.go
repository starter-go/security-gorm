package impldao

import (
	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

type finder struct {
	db         *gorm.DB
	page       *rbac.Pagination
	conditions *rbac.Conditions
	itemModel  any
	listModel  any
}

func (inst *finder) find() error {

	list := inst.listModel
	item := inst.itemModel
	db := inst.db

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
