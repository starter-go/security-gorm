package testcom

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/starter-go/application"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/src/test/code/cases"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

// TestAdminUsers ... 测试 /api/admin.v1/Users/*
type TestAdminUsers struct {

	//starter:component
	_as func(units.Units) //starter:as(".")

	AC          application.Context //starter:inject("context")
	Userservice rbac.UserService    //starter:inject("#")
}

func (inst *TestAdminUsers) _impl() units.Units {
	return inst
}

// Units ...
func (inst *TestAdminUsers) Units(list []*units.Registration) []*units.Registration {

	list = append(list, &units.Registration{
		Name:    cases.GetUserList,
		Enabled: true,
		Test:    inst.doTestGetList,
	})
	list = append(list, &units.Registration{
		Name:    cases.InsertUser,
		Enabled: true,
		Test:    inst.doTestInsert,
	})
	list = append(list, &units.Registration{
		Name:    cases.DoUserCRUD,
		Enabled: true,
		Test:    inst.doTestCRUD,
	})

	return list
}

// // Boots ...
// func (inst *TestAdminUsers) Boots() []*testboot.Boot {
// 	bl := testboot.BootList{}
// 	bl.Handle(http.MethodGet, "/api/v1/users", inst.doTestGetList)
// 	bl.Handle(http.MethodPost, "/api/v1/users", inst.doTestInsert)
// 	bl.Handle(http.MethodPost, "/api/v1/users/crud", inst.doTestCRUD)
// 	return bl.List()
// }

func (inst *TestAdminUsers) makeUserName() rbac.UserName {
	now := lang.Now()
	str := strconv.FormatInt(now.Int(), 10)
	return rbac.UserName("user-" + str)
}

func (inst *TestAdminUsers) doTestInsert() error {
	ctx := inst.AC
	name := inst.makeUserName()
	o1 := &rbac.UserDTO{
		Name: name,
	}
	o2, err := inst.Userservice.Insert(ctx, o1)
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

func (inst *TestAdminUsers) doTestUpdate() error {
	return fmt.Errorf("no impl")
}

func (inst *TestAdminUsers) doTestDelete() error {
	return fmt.Errorf("no impl")
}

func (inst *TestAdminUsers) doTestGetOne() error {
	ctx := inst.AC
	const id = 1
	o, err := inst.Userservice.Find(ctx, id)
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

func (inst *TestAdminUsers) doTestGetList() error {
	ctx := inst.AC
	ser := inst.Userservice
	list, err := ser.List(ctx, nil)
	if err != nil {
		return err
	}
	inst.log("list:", list)
	return nil
}

func (inst *TestAdminUsers) doTestCRUD() error {

	ctx := inst.AC
	ser := inst.Userservice
	name := inst.makeUserName()

	o1 := &rbac.UserDTO{
		Name: name,
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

func (inst *TestAdminUsers) log(format string, o any) {
	data, err := json.MarshalIndent(o, "", "\t")
	if err == nil {
		str := string(data)
		vlog.Info(format, str)
	}
}
