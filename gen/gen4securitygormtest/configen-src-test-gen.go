package gen4securitygormtest
import (
    p0ef6f2938 "github.com/starter-go/application"
    pe31a5ae46 "github.com/starter-go/security-gorm/src/test/code/testcom"
    p2dece1e49 "github.com/starter-go/security/rbac"
     "github.com/starter-go/application"
)

// type pe31a5ae46.Bootstrap in package:github.com/starter-go/security-gorm/src/test/code/testcom
//
// id:com-e31a5ae465e830bc-testcom-Bootstrap
// class:class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle
// alias:
// scope:singleton
//
type pe31a5ae465_testcom_Bootstrap struct {
}

func (inst* pe31a5ae465_testcom_Bootstrap) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e31a5ae465e830bc-testcom-Bootstrap"
	r.Classes = "class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe31a5ae465_testcom_Bootstrap) new() any {
    return &pe31a5ae46.Bootstrap{}
}

func (inst* pe31a5ae465_testcom_Bootstrap) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe31a5ae46.Bootstrap)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)
    com.RbacAuthService = inst.getRbacAuthService(ie)


    return nil
}


func (inst*pe31a5ae465_testcom_Bootstrap) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*pe31a5ae465_testcom_Bootstrap) getRbacAuthService(ie application.InjectionExt)p2dece1e49.AuthService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-AuthService").(p2dece1e49.AuthService)
}



// type pe31a5ae46.TestAuthWithEmail in package:github.com/starter-go/security-gorm/src/test/code/testcom
//
// id:com-e31a5ae465e830bc-testcom-TestAuthWithEmail
// class:class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle
// alias:
// scope:singleton
//
type pe31a5ae465_testcom_TestAuthWithEmail struct {
}

func (inst* pe31a5ae465_testcom_TestAuthWithEmail) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e31a5ae465e830bc-testcom-TestAuthWithEmail"
	r.Classes = "class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle"
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


    return nil
}


func (inst*pe31a5ae465_testcom_TestAuthWithEmail) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
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
// class:class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle
// alias:
// scope:singleton
//
type pe31a5ae465_testcom_TestAuthWithUserPassword struct {
}

func (inst* pe31a5ae465_testcom_TestAuthWithUserPassword) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e31a5ae465e830bc-testcom-TestAuthWithUserPassword"
	r.Classes = "class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle"
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
    com.Enabled = inst.getEnabled(ie)


    return nil
}


func (inst*pe31a5ae465_testcom_TestAuthWithUserPassword) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*pe31a5ae465_testcom_TestAuthWithUserPassword) getRbacAuthService(ie application.InjectionExt)p2dece1e49.AuthService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-AuthService").(p2dece1e49.AuthService)
}


func (inst*pe31a5ae465_testcom_TestAuthWithUserPassword) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${testcase.enable.auth-with-password}")
}


