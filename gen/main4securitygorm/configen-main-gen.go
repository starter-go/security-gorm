package main4securitygorm

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p047c6784c4_implconvertor_GroupConvertorImpl{})
    inst.register(&p047c6784c4_implconvertor_PermissionConvertorImpl{})
    inst.register(&p047c6784c4_implconvertor_RegionConvertorImpl{})
    inst.register(&p047c6784c4_implconvertor_RoleConvertorImpl{})
    inst.register(&p047c6784c4_implconvertor_UserConvertorImpl{})
    inst.register(&p0d13f39fa5_implservice_GroupServiceImpl{})
    inst.register(&p0d13f39fa5_implservice_PermissionCacheImpl{})
    inst.register(&p0d13f39fa5_implservice_PermissionServiceImpl{})
    inst.register(&p0d13f39fa5_implservice_RegionServiceImpl{})
    inst.register(&p0d13f39fa5_implservice_RoleServiceImpl{})
    inst.register(&p0d13f39fa5_implservice_UserServiceImpl{})
    inst.register(&p8617045c5b_impldao_EmailAddressDaoImpl{})
    inst.register(&p8617045c5b_impldao_GroupDaoImpl{})
    inst.register(&p8617045c5b_impldao_PermissionDaoImpl{})
    inst.register(&p8617045c5b_impldao_PhoneNumberDaoImpl{})
    inst.register(&p8617045c5b_impldao_RegionDaoImpl{})
    inst.register(&p8617045c5b_impldao_RoleDaoImpl{})
    inst.register(&p8617045c5b_impldao_UserDaoImpl{})
    inst.register(&pd2ff24bd19_password_RootAuth{})
    inst.register(&pd2ff24bd19_password_UsersAuth{})
    inst.register(&pf5d2c6fae0_rbacdb_LocalAgentImpl{})
    inst.register(&pf5d2c6fae0_rbacdb_ThisGroupRegistry{})


    return nil
}
