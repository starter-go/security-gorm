package gen4securitygorm
import (
    p512a30914 "github.com/starter-go/libgorm"
    pf5d2c6fae "github.com/starter-go/security-gorm/rbacdb"
     "github.com/starter-go/application"
)

// type pf5d2c6fae.LocalAgentImpl in package:github.com/starter-go/security-gorm/rbacdb
//
// id:com-f5d2c6fae0368143-rbacdb-LocalAgentImpl
// class:
// alias:alias-f5d2c6fae036814399fa2ed06c0dc99f-LocalAgent
// scope:singleton
//
type pf5d2c6fae0_rbacdb_LocalAgentImpl struct {
}

func (inst* pf5d2c6fae0_rbacdb_LocalAgentImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f5d2c6fae0368143-rbacdb-LocalAgentImpl"
	r.Classes = ""
	r.Aliases = "alias-f5d2c6fae036814399fa2ed06c0dc99f-LocalAgent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf5d2c6fae0_rbacdb_LocalAgentImpl) new() any {
    return &pf5d2c6fae.LocalAgentImpl{}
}

func (inst* pf5d2c6fae0_rbacdb_LocalAgentImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf5d2c6fae.LocalAgentImpl)
	nop(ie, com)

	
    com.DataSourceMan = inst.getDataSourceMan(ie)
    com.DataSourceAlias = inst.getDataSourceAlias(ie)


    return nil
}


func (inst*pf5d2c6fae0_rbacdb_LocalAgentImpl) getDataSourceMan(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}


func (inst*pf5d2c6fae0_rbacdb_LocalAgentImpl) getDataSourceAlias(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.rbac.datasource}")
}



// type pf5d2c6fae.ThisGroupRegistry in package:github.com/starter-go/security-gorm/rbacdb
//
// id:com-f5d2c6fae0368143-rbacdb-ThisGroupRegistry
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:
// scope:singleton
//
type pf5d2c6fae0_rbacdb_ThisGroupRegistry struct {
}

func (inst* pf5d2c6fae0_rbacdb_ThisGroupRegistry) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f5d2c6fae0368143-rbacdb-ThisGroupRegistry"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf5d2c6fae0_rbacdb_ThisGroupRegistry) new() any {
    return &pf5d2c6fae.ThisGroupRegistry{}
}

func (inst* pf5d2c6fae0_rbacdb_ThisGroupRegistry) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf5d2c6fae.ThisGroupRegistry)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.Prefix = inst.getPrefix(ie)
    com.Source = inst.getSource(ie)


    return nil
}


func (inst*pf5d2c6fae0_rbacdb_ThisGroupRegistry) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${datagroup.rbac.enabled}")
}


func (inst*pf5d2c6fae0_rbacdb_ThisGroupRegistry) getPrefix(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.rbac.table-name-prefix}")
}


func (inst*pf5d2c6fae0_rbacdb_ThisGroupRegistry) getSource(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.rbac.datasource}")
}


