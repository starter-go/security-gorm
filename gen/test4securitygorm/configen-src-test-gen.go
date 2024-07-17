package test4securitygorm
import (
    p0ef6f2938 "github.com/starter-go/application"
    p24287f458 "github.com/starter-go/rbac"
    pe31a5ae46 "github.com/starter-go/security-gorm/src/test/code/testcom"
     "github.com/starter-go/application"
)

// type pe31a5ae46.TestAdminPermissions in package:github.com/starter-go/security-gorm/src/test/code/testcom
//
// id:com-e31a5ae465e830bc-testcom-TestAdminPermissions
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pe31a5ae465_testcom_TestAdminPermissions struct {
}

func (inst* pe31a5ae465_testcom_TestAdminPermissions) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e31a5ae465e830bc-testcom-TestAdminPermissions"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe31a5ae465_testcom_TestAdminPermissions) new() any {
    return &pe31a5ae46.TestAdminPermissions{}
}

func (inst* pe31a5ae465_testcom_TestAdminPermissions) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe31a5ae46.TestAdminPermissions)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)
    com.Permissionservice = inst.getPermissionservice(ie)


    return nil
}


func (inst*pe31a5ae465_testcom_TestAdminPermissions) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*pe31a5ae465_testcom_TestAdminPermissions) getPermissionservice(ie application.InjectionExt)p24287f458.PermissionService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-PermissionService").(p24287f458.PermissionService)
}



// type pe31a5ae46.TestAdminRoles in package:github.com/starter-go/security-gorm/src/test/code/testcom
//
// id:com-e31a5ae465e830bc-testcom-TestAdminRoles
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pe31a5ae465_testcom_TestAdminRoles struct {
}

func (inst* pe31a5ae465_testcom_TestAdminRoles) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e31a5ae465e830bc-testcom-TestAdminRoles"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe31a5ae465_testcom_TestAdminRoles) new() any {
    return &pe31a5ae46.TestAdminRoles{}
}

func (inst* pe31a5ae465_testcom_TestAdminRoles) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe31a5ae46.TestAdminRoles)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)
    com.RoleService = inst.getRoleService(ie)


    return nil
}


func (inst*pe31a5ae465_testcom_TestAdminRoles) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*pe31a5ae465_testcom_TestAdminRoles) getRoleService(ie application.InjectionExt)p24287f458.RoleService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-RoleService").(p24287f458.RoleService)
}



// type pe31a5ae46.TestAdminUsers in package:github.com/starter-go/security-gorm/src/test/code/testcom
//
// id:com-e31a5ae465e830bc-testcom-TestAdminUsers
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pe31a5ae465_testcom_TestAdminUsers struct {
}

func (inst* pe31a5ae465_testcom_TestAdminUsers) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e31a5ae465e830bc-testcom-TestAdminUsers"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe31a5ae465_testcom_TestAdminUsers) new() any {
    return &pe31a5ae46.TestAdminUsers{}
}

func (inst* pe31a5ae465_testcom_TestAdminUsers) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe31a5ae46.TestAdminUsers)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)
    com.Userservice = inst.getUserservice(ie)


    return nil
}


func (inst*pe31a5ae465_testcom_TestAdminUsers) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*pe31a5ae465_testcom_TestAdminUsers) getUserservice(ie application.InjectionExt)p24287f458.UserService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-UserService").(p24287f458.UserService)
}



// type pe31a5ae46.TestAuthWithEmail in package:github.com/starter-go/security-gorm/src/test/code/testcom
//
// id:com-e31a5ae465e830bc-testcom-TestAuthWithEmail
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pe31a5ae465_testcom_TestAuthWithEmail struct {
}

func (inst* pe31a5ae465_testcom_TestAuthWithEmail) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e31a5ae465e830bc-testcom-TestAuthWithEmail"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe31a5ae465_testcom_TestAuthWithEmail) new() any {
    return &pe31a5ae46.TestAuthWithEmail{}
}

func (inst* pe31a5ae465_testcom_TestAuthWithEmail) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe31a5ae46.TestAuthWithEmail)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)
    com.AuthService = inst.getAuthService(ie)


    return nil
}


func (inst*pe31a5ae465_testcom_TestAuthWithEmail) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*pe31a5ae465_testcom_TestAuthWithEmail) getAuthService(ie application.InjectionExt)p24287f458.AuthService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-AuthService").(p24287f458.AuthService)
}



// type pe31a5ae46.TestAuthWithSMS in package:github.com/starter-go/security-gorm/src/test/code/testcom
//
// id:com-e31a5ae465e830bc-testcom-TestAuthWithSMS
// class:class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle
// alias:
// scope:singleton
//
type pe31a5ae465_testcom_TestAuthWithSMS struct {
}

func (inst* pe31a5ae465_testcom_TestAuthWithSMS) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e31a5ae465e830bc-testcom-TestAuthWithSMS"
	r.Classes = "class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe31a5ae465_testcom_TestAuthWithSMS) new() any {
    return &pe31a5ae46.TestAuthWithSMS{}
}

func (inst* pe31a5ae465_testcom_TestAuthWithSMS) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe31a5ae46.TestAuthWithSMS)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)


    return nil
}


func (inst*pe31a5ae465_testcom_TestAuthWithSMS) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}



// type pe31a5ae46.TestAuthWithUserPassword in package:github.com/starter-go/security-gorm/src/test/code/testcom
//
// id:com-e31a5ae465e830bc-testcom-TestAuthWithUserPassword
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pe31a5ae465_testcom_TestAuthWithUserPassword struct {
}

func (inst* pe31a5ae465_testcom_TestAuthWithUserPassword) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e31a5ae465e830bc-testcom-TestAuthWithUserPassword"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe31a5ae465_testcom_TestAuthWithUserPassword) new() any {
    return &pe31a5ae46.TestAuthWithUserPassword{}
}

func (inst* pe31a5ae465_testcom_TestAuthWithUserPassword) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe31a5ae46.TestAuthWithUserPassword)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)
    com.RbacAuthService = inst.getRbacAuthService(ie)


    return nil
}


func (inst*pe31a5ae465_testcom_TestAuthWithUserPassword) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*pe31a5ae465_testcom_TestAuthWithUserPassword) getRbacAuthService(ie application.InjectionExt)p24287f458.AuthService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-AuthService").(p24287f458.AuthService)
}


