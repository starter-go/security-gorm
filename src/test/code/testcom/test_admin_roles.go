package testcom

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/starter-go/application"
	"github.com/starter-go/security-gorm/src/test/code/testboot"
	"github.com/starter-go/security/rbac"
	"github.com/starter-go/vlog"
)

// TestAdminRoles ... 测试 /api/admin.v1/roles/*
type TestAdminRoles struct {

	//starter:component
	_as func(testboot.BootingRegistry) //starter:as(".")

	AC          application.Context //starter:inject("context")
	RoleService rbac.RoleService    //starter:inject("#")
}

func (inst *TestAdminRoles) _impl() {
	inst._as(inst)
}

// Boots ...
func (inst *TestAdminRoles) Boots() []*testboot.Boot {
	bl := testboot.BootList{}
	bl.Handle(http.MethodGet, "/api/v1/roles", inst.doTestGetList)
	bl.Handle(http.MethodPost, "/api/v1/roles", inst.doTestInsert)
	bl.Handle(http.MethodPost, "/api/v1/roles/crud", inst.doTestCRUD)
	return bl.List()
}

func (inst *TestAdminRoles) doTestInsert() error {
	ctx := inst.AC
	o1 := &rbac.RoleDTO{
		Name: "test",
	}
	o2, err := inst.RoleService.Insert(ctx, o1)
	if err != nil {
		return err
	}
	str, err := json.MarshalIndent(o2, "", "\t")
	if err != nil {
		return err
	}
	vlog.Debug("insert result:  %s", str)
	return nil
}

func (inst *TestAdminRoles) doTestUpdate() error {
	return fmt.Errorf("no impl")
}

func (inst *TestAdminRoles) doTestDelete() error {
	return fmt.Errorf("no impl")
}

func (inst *TestAdminRoles) doTestGetOne() error {
	ctx := inst.AC
	const id = 1
	o, err := inst.RoleService.Find(ctx, id)
	if err != nil {
		return err
	}
	str, err := json.MarshalIndent(o, "", "\t")
	if err != nil {
		return err
	}
	vlog.Debug("find result:  %s", str)
	return nil
}

func (inst *TestAdminRoles) doTestGetList() error {
	return fmt.Errorf("no impl")
}

func (inst *TestAdminRoles) doTestCRUD() error {

	ctx := inst.AC
	ser := inst.RoleService

	o1 := &rbac.RoleDTO{
		Name: "test",
	}

	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}

	id := o2.ID
	o3, err := ser.Update(ctx, id, o2)
	if err != nil {
		return err
	}

	o4, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}

	list, err := ser.List(ctx, nil)
	if err != nil {
		return err
	}

	inst.log("o1: %s", o1)
	inst.log("o2: %s", o2)
	inst.log("o3: %s", o3)
	inst.log("o4: %s", o4)
	inst.log("list:", list)

	return ser.Delete(ctx, id)
}

func (inst *TestAdminRoles) log(format string, o any) {
	data, err := json.MarshalIndent(o, "", "\t")
	if err == nil {
		str := string(data)
		vlog.Info(format, str)
	}
}
