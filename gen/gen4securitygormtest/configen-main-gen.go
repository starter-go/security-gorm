package gen4securitygormtest

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

    
    inst.register(&pe31a5ae465_testcom_Bootstrap{})
    inst.register(&pe31a5ae465_testcom_TestAuthWithEmail{})
    inst.register(&pe31a5ae465_testcom_TestAuthWithSMS{})
    inst.register(&pe31a5ae465_testcom_TestAuthWithUserPassword{})


    return nil
}
