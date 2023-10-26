package gen4securitygorm
import (
    pd2ff24bd1 "github.com/starter-go/security-gorm/internal/auth1/password"
    p047c6784c "github.com/starter-go/security-gorm/internal/implconvertor"
    p8617045c5 "github.com/starter-go/security-gorm/internal/impldao"
    p0d13f39fa "github.com/starter-go/security-gorm/internal/implservice"
    pf5d2c6fae "github.com/starter-go/security-gorm/rbacdb"
    p9621e8b71 "github.com/starter-go/security/random"
    p2dece1e49 "github.com/starter-go/security/rbac"
     "github.com/starter-go/application"
)

// type pd2ff24bd1.RootAuth in package:github.com/starter-go/security-gorm/internal/auth1/password
//
// id:com-d2ff24bd19ca2e13-password-RootAuth
// class:class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle class-9d209f7c2504d33e6054a2c9998e9485-Registry
// alias:
// scope:singleton
//
type pd2ff24bd19_password_RootAuth struct {
}

func (inst* pd2ff24bd19_password_RootAuth) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d2ff24bd19ca2e13-password-RootAuth"
	r.Classes = "class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle class-9d209f7c2504d33e6054a2c9998e9485-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd2ff24bd19_password_RootAuth) new() any {
    return &pd2ff24bd1.RootAuth{}
}

func (inst* pd2ff24bd19_password_RootAuth) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd2ff24bd1.RootAuth)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.Random = inst.getRandom(ie)


    return nil
}


func (inst*pd2ff24bd19_password_RootAuth) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${security.auth-as-root.enabled}")
}


func (inst*pd2ff24bd19_password_RootAuth) getRandom(ie application.InjectionExt)p9621e8b71.Service{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-Service").(p9621e8b71.Service)
}



// type pd2ff24bd1.UsersAuth in package:github.com/starter-go/security-gorm/internal/auth1/password
//
// id:com-d2ff24bd19ca2e13-password-UsersAuth
// class:class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle class-9d209f7c2504d33e6054a2c9998e9485-Registry
// alias:
// scope:singleton
//
type pd2ff24bd19_password_UsersAuth struct {
}

func (inst* pd2ff24bd19_password_UsersAuth) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d2ff24bd19ca2e13-password-UsersAuth"
	r.Classes = "class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle class-9d209f7c2504d33e6054a2c9998e9485-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd2ff24bd19_password_UsersAuth) new() any {
    return &pd2ff24bd1.UsersAuth{}
}

func (inst* pd2ff24bd19_password_UsersAuth) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd2ff24bd1.UsersAuth)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.Random = inst.getRandom(ie)
    com.UserDao = inst.getUserDao(ie)
    com.UserCvt = inst.getUserCvt(ie)


    return nil
}


func (inst*pd2ff24bd19_password_UsersAuth) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${security.auth-as-root.enabled}")
}


func (inst*pd2ff24bd19_password_UsersAuth) getRandom(ie application.InjectionExt)p9621e8b71.Service{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-Service").(p9621e8b71.Service)
}


func (inst*pd2ff24bd19_password_UsersAuth) getUserDao(ie application.InjectionExt)pf5d2c6fae.UserDAO{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-UserDAO").(pf5d2c6fae.UserDAO)
}


func (inst*pd2ff24bd19_password_UsersAuth) getUserCvt(ie application.InjectionExt)pf5d2c6fae.UserConvertor{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-UserConvertor").(pf5d2c6fae.UserConvertor)
}



// type p047c6784c.PermissionConvertorImpl in package:github.com/starter-go/security-gorm/internal/implconvertor
//
// id:com-047c6784c4625f2a-implconvertor-PermissionConvertorImpl
// class:
// alias:alias-f5d2c6fae036814399fa2ed06c0dc99f-PermissionConvertor
// scope:singleton
//
type p047c6784c4_implconvertor_PermissionConvertorImpl struct {
}

func (inst* p047c6784c4_implconvertor_PermissionConvertorImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-047c6784c4625f2a-implconvertor-PermissionConvertorImpl"
	r.Classes = ""
	r.Aliases = "alias-f5d2c6fae036814399fa2ed06c0dc99f-PermissionConvertor"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p047c6784c4_implconvertor_PermissionConvertorImpl) new() any {
    return &p047c6784c.PermissionConvertorImpl{}
}

func (inst* p047c6784c4_implconvertor_PermissionConvertorImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p047c6784c.PermissionConvertorImpl)
	nop(ie, com)

	


    return nil
}



// type p047c6784c.RoleConvertorImpl in package:github.com/starter-go/security-gorm/internal/implconvertor
//
// id:com-047c6784c4625f2a-implconvertor-RoleConvertorImpl
// class:
// alias:alias-f5d2c6fae036814399fa2ed06c0dc99f-RoleConvertor
// scope:singleton
//
type p047c6784c4_implconvertor_RoleConvertorImpl struct {
}

func (inst* p047c6784c4_implconvertor_RoleConvertorImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-047c6784c4625f2a-implconvertor-RoleConvertorImpl"
	r.Classes = ""
	r.Aliases = "alias-f5d2c6fae036814399fa2ed06c0dc99f-RoleConvertor"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p047c6784c4_implconvertor_RoleConvertorImpl) new() any {
    return &p047c6784c.RoleConvertorImpl{}
}

func (inst* p047c6784c4_implconvertor_RoleConvertorImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p047c6784c.RoleConvertorImpl)
	nop(ie, com)

	


    return nil
}



// type p047c6784c.UserConvertorImpl in package:github.com/starter-go/security-gorm/internal/implconvertor
//
// id:com-047c6784c4625f2a-implconvertor-UserConvertorImpl
// class:
// alias:alias-f5d2c6fae036814399fa2ed06c0dc99f-UserConvertor
// scope:singleton
//
type p047c6784c4_implconvertor_UserConvertorImpl struct {
}

func (inst* p047c6784c4_implconvertor_UserConvertorImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-047c6784c4625f2a-implconvertor-UserConvertorImpl"
	r.Classes = ""
	r.Aliases = "alias-f5d2c6fae036814399fa2ed06c0dc99f-UserConvertor"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p047c6784c4_implconvertor_UserConvertorImpl) new() any {
    return &p047c6784c.UserConvertorImpl{}
}

func (inst* p047c6784c4_implconvertor_UserConvertorImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p047c6784c.UserConvertorImpl)
	nop(ie, com)

	


    return nil
}



// type p8617045c5.EmailAddressDaoImpl in package:github.com/starter-go/security-gorm/internal/impldao
//
// id:com-8617045c5b2ce8e3-impldao-EmailAddressDaoImpl
// class:
// alias:alias-f5d2c6fae036814399fa2ed06c0dc99f-EmailAddressDAO
// scope:singleton
//
type p8617045c5b_impldao_EmailAddressDaoImpl struct {
}

func (inst* p8617045c5b_impldao_EmailAddressDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8617045c5b2ce8e3-impldao-EmailAddressDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-f5d2c6fae036814399fa2ed06c0dc99f-EmailAddressDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8617045c5b_impldao_EmailAddressDaoImpl) new() any {
    return &p8617045c5.EmailAddressDaoImpl{}
}

func (inst* p8617045c5b_impldao_EmailAddressDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8617045c5.EmailAddressDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p8617045c5b_impldao_EmailAddressDaoImpl) getAgent(ie application.InjectionExt)pf5d2c6fae.LocalAgent{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-LocalAgent").(pf5d2c6fae.LocalAgent)
}


func (inst*p8617045c5b_impldao_EmailAddressDaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p8617045c5.PermissionDaoImpl in package:github.com/starter-go/security-gorm/internal/impldao
//
// id:com-8617045c5b2ce8e3-impldao-PermissionDaoImpl
// class:
// alias:alias-f5d2c6fae036814399fa2ed06c0dc99f-PermissionDAO
// scope:singleton
//
type p8617045c5b_impldao_PermissionDaoImpl struct {
}

func (inst* p8617045c5b_impldao_PermissionDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8617045c5b2ce8e3-impldao-PermissionDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-f5d2c6fae036814399fa2ed06c0dc99f-PermissionDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8617045c5b_impldao_PermissionDaoImpl) new() any {
    return &p8617045c5.PermissionDaoImpl{}
}

func (inst* p8617045c5b_impldao_PermissionDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8617045c5.PermissionDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p8617045c5b_impldao_PermissionDaoImpl) getAgent(ie application.InjectionExt)pf5d2c6fae.LocalAgent{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-LocalAgent").(pf5d2c6fae.LocalAgent)
}


func (inst*p8617045c5b_impldao_PermissionDaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p8617045c5.PhoneNumberDaoImpl in package:github.com/starter-go/security-gorm/internal/impldao
//
// id:com-8617045c5b2ce8e3-impldao-PhoneNumberDaoImpl
// class:
// alias:alias-f5d2c6fae036814399fa2ed06c0dc99f-PhoneNumberDAO
// scope:singleton
//
type p8617045c5b_impldao_PhoneNumberDaoImpl struct {
}

func (inst* p8617045c5b_impldao_PhoneNumberDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8617045c5b2ce8e3-impldao-PhoneNumberDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-f5d2c6fae036814399fa2ed06c0dc99f-PhoneNumberDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8617045c5b_impldao_PhoneNumberDaoImpl) new() any {
    return &p8617045c5.PhoneNumberDaoImpl{}
}

func (inst* p8617045c5b_impldao_PhoneNumberDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8617045c5.PhoneNumberDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p8617045c5b_impldao_PhoneNumberDaoImpl) getAgent(ie application.InjectionExt)pf5d2c6fae.LocalAgent{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-LocalAgent").(pf5d2c6fae.LocalAgent)
}


func (inst*p8617045c5b_impldao_PhoneNumberDaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p8617045c5.RoleDaoImpl in package:github.com/starter-go/security-gorm/internal/impldao
//
// id:com-8617045c5b2ce8e3-impldao-RoleDaoImpl
// class:
// alias:alias-f5d2c6fae036814399fa2ed06c0dc99f-RoleDAO
// scope:singleton
//
type p8617045c5b_impldao_RoleDaoImpl struct {
}

func (inst* p8617045c5b_impldao_RoleDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8617045c5b2ce8e3-impldao-RoleDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-f5d2c6fae036814399fa2ed06c0dc99f-RoleDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8617045c5b_impldao_RoleDaoImpl) new() any {
    return &p8617045c5.RoleDaoImpl{}
}

func (inst* p8617045c5b_impldao_RoleDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8617045c5.RoleDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p8617045c5b_impldao_RoleDaoImpl) getAgent(ie application.InjectionExt)pf5d2c6fae.LocalAgent{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-LocalAgent").(pf5d2c6fae.LocalAgent)
}


func (inst*p8617045c5b_impldao_RoleDaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p8617045c5.UserDaoImpl in package:github.com/starter-go/security-gorm/internal/impldao
//
// id:com-8617045c5b2ce8e3-impldao-UserDaoImpl
// class:
// alias:alias-f5d2c6fae036814399fa2ed06c0dc99f-UserDAO
// scope:singleton
//
type p8617045c5b_impldao_UserDaoImpl struct {
}

func (inst* p8617045c5b_impldao_UserDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8617045c5b2ce8e3-impldao-UserDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-f5d2c6fae036814399fa2ed06c0dc99f-UserDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8617045c5b_impldao_UserDaoImpl) new() any {
    return &p8617045c5.UserDaoImpl{}
}

func (inst* p8617045c5b_impldao_UserDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8617045c5.UserDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.EmailAddressDAO = inst.getEmailAddressDAO(ie)
    com.PhoneNumberDAO = inst.getPhoneNumberDAO(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p8617045c5b_impldao_UserDaoImpl) getAgent(ie application.InjectionExt)pf5d2c6fae.LocalAgent{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-LocalAgent").(pf5d2c6fae.LocalAgent)
}


func (inst*p8617045c5b_impldao_UserDaoImpl) getEmailAddressDAO(ie application.InjectionExt)pf5d2c6fae.EmailAddressDAO{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-EmailAddressDAO").(pf5d2c6fae.EmailAddressDAO)
}


func (inst*p8617045c5b_impldao_UserDaoImpl) getPhoneNumberDAO(ie application.InjectionExt)pf5d2c6fae.PhoneNumberDAO{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-PhoneNumberDAO").(pf5d2c6fae.PhoneNumberDAO)
}


func (inst*p8617045c5b_impldao_UserDaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p0d13f39fa.PermissionCacheImpl in package:github.com/starter-go/security-gorm/internal/implservice
//
// id:com-0d13f39fa52fea3f-implservice-PermissionCacheImpl
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-PermissionCache
// scope:singleton
//
type p0d13f39fa5_implservice_PermissionCacheImpl struct {
}

func (inst* p0d13f39fa5_implservice_PermissionCacheImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0d13f39fa52fea3f-implservice-PermissionCacheImpl"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-PermissionCache"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0d13f39fa5_implservice_PermissionCacheImpl) new() any {
    return &p0d13f39fa.PermissionCacheImpl{}
}

func (inst* p0d13f39fa5_implservice_PermissionCacheImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0d13f39fa.PermissionCacheImpl)
	nop(ie, com)

	
    com.PermissionService = inst.getPermissionService(ie)


    return nil
}


func (inst*p0d13f39fa5_implservice_PermissionCacheImpl) getPermissionService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}



// type p0d13f39fa.PermissionServiceImpl in package:github.com/starter-go/security-gorm/internal/implservice
//
// id:com-0d13f39fa52fea3f-implservice-PermissionServiceImpl
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-PermissionService
// scope:singleton
//
type p0d13f39fa5_implservice_PermissionServiceImpl struct {
}

func (inst* p0d13f39fa5_implservice_PermissionServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0d13f39fa52fea3f-implservice-PermissionServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-PermissionService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0d13f39fa5_implservice_PermissionServiceImpl) new() any {
    return &p0d13f39fa.PermissionServiceImpl{}
}

func (inst* p0d13f39fa5_implservice_PermissionServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0d13f39fa.PermissionServiceImpl)
	nop(ie, com)

	
    com.PermissionDao = inst.getPermissionDao(ie)
    com.PermissionConvertor = inst.getPermissionConvertor(ie)
    com.PermissionCache = inst.getPermissionCache(ie)


    return nil
}


func (inst*p0d13f39fa5_implservice_PermissionServiceImpl) getPermissionDao(ie application.InjectionExt)pf5d2c6fae.PermissionDAO{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-PermissionDAO").(pf5d2c6fae.PermissionDAO)
}


func (inst*p0d13f39fa5_implservice_PermissionServiceImpl) getPermissionConvertor(ie application.InjectionExt)pf5d2c6fae.PermissionConvertor{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-PermissionConvertor").(pf5d2c6fae.PermissionConvertor)
}


func (inst*p0d13f39fa5_implservice_PermissionServiceImpl) getPermissionCache(ie application.InjectionExt)p2dece1e49.PermissionCache{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionCache").(p2dece1e49.PermissionCache)
}



// type p0d13f39fa.RoleServiceImpl in package:github.com/starter-go/security-gorm/internal/implservice
//
// id:com-0d13f39fa52fea3f-implservice-RoleServiceImpl
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-RoleService
// scope:singleton
//
type p0d13f39fa5_implservice_RoleServiceImpl struct {
}

func (inst* p0d13f39fa5_implservice_RoleServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0d13f39fa52fea3f-implservice-RoleServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-RoleService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0d13f39fa5_implservice_RoleServiceImpl) new() any {
    return &p0d13f39fa.RoleServiceImpl{}
}

func (inst* p0d13f39fa5_implservice_RoleServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0d13f39fa.RoleServiceImpl)
	nop(ie, com)

	
    com.RoleDao = inst.getRoleDao(ie)
    com.RoleConvertor = inst.getRoleConvertor(ie)


    return nil
}


func (inst*p0d13f39fa5_implservice_RoleServiceImpl) getRoleDao(ie application.InjectionExt)pf5d2c6fae.RoleDAO{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-RoleDAO").(pf5d2c6fae.RoleDAO)
}


func (inst*p0d13f39fa5_implservice_RoleServiceImpl) getRoleConvertor(ie application.InjectionExt)pf5d2c6fae.RoleConvertor{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-RoleConvertor").(pf5d2c6fae.RoleConvertor)
}



// type p0d13f39fa.UserServiceImpl in package:github.com/starter-go/security-gorm/internal/implservice
//
// id:com-0d13f39fa52fea3f-implservice-UserServiceImpl
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-UserService
// scope:singleton
//
type p0d13f39fa5_implservice_UserServiceImpl struct {
}

func (inst* p0d13f39fa5_implservice_UserServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0d13f39fa52fea3f-implservice-UserServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-UserService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0d13f39fa5_implservice_UserServiceImpl) new() any {
    return &p0d13f39fa.UserServiceImpl{}
}

func (inst* p0d13f39fa5_implservice_UserServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0d13f39fa.UserServiceImpl)
	nop(ie, com)

	
    com.UserDao = inst.getUserDao(ie)
    com.UserConvertor = inst.getUserConvertor(ie)


    return nil
}


func (inst*p0d13f39fa5_implservice_UserServiceImpl) getUserDao(ie application.InjectionExt)pf5d2c6fae.UserDAO{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-UserDAO").(pf5d2c6fae.UserDAO)
}


func (inst*p0d13f39fa5_implservice_UserServiceImpl) getUserConvertor(ie application.InjectionExt)pf5d2c6fae.UserConvertor{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-UserConvertor").(pf5d2c6fae.UserConvertor)
}


