package implservice

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/rbac"
)

// PermissionCacheImpl ...
type PermissionCacheImpl struct {

	//starter:component
	_as func(rbac.PermissionCache) //starter:as("#")

	PermissionService rbac.PermissionService //starter:inject("#")

	snapshot *permissionCacheSnapshot
	mutex    sync.Mutex
}

func (inst *PermissionCacheImpl) _impl() rbac.PermissionCache {
	return inst
}

// Clear ...
func (inst *PermissionCacheImpl) Clear() {
	inst.snapshot = nil
}

// Find ...
func (inst *PermissionCacheImpl) Find(c context.Context, want *rbac.PermissionDTO) (*rbac.PermissionDTO, error) {

	inst.mutex.Lock()
	defer inst.mutex.Unlock()

	snap := inst.snapshot
	if snap != nil {
		if snap.isExpired() {
			snap = nil
		}
	}

	if snap == nil {
		snap = &permissionCacheSnapshot{}
		snap.init()
		snap.load(c, inst.PermissionService)
		inst.snapshot = snap
	}

	return snap.find(want)
}

////////////////////////////////////////////////////////////////////////////////

type permissionCacheSnapshot struct {
	expiredAt lang.Time
	table     map[string]*rbac.PermissionDTO
}

func (inst *permissionCacheSnapshot) init() {
	inst.table = make(map[string]*rbac.PermissionDTO)
	inst.expiredAt = lang.Now() + (1000 * 600) // max-age:10min
}

func (inst *permissionCacheSnapshot) keyFor(o *rbac.PermissionDTO) string {
	return strings.ToUpper(o.Method) + ":" + o.Path
}

func (inst *permissionCacheSnapshot) load(c context.Context, ser rbac.PermissionService) {
	tab := make(map[string]*rbac.PermissionDTO)
	query := &rbac.PermissionQuery{All: true}
	list, err := ser.List(c, query)
	if err == nil {
		for _, item := range list {
			key := inst.keyFor(item)
			tab[key] = item
		}
	}
	inst.table = tab
}

func (inst *permissionCacheSnapshot) isExpired() bool {
	now := lang.Now()
	return inst.expiredAt < now
}

func (inst *permissionCacheSnapshot) find(want *rbac.PermissionDTO) (*rbac.PermissionDTO, error) {
	key := inst.keyFor(want)
	item := inst.table[key]
	if item == nil {
		return nil, fmt.Errorf("no permission: [%s]", key)
	}
	item2 := &rbac.PermissionDTO{}
	*item2 = *item
	return item2, nil
}
