package testcom

import (
	"fmt"

	"github.com/starter-go/application"
	"github.com/starter-go/security-gorm/src/test/code/testboot"
	"github.com/starter-go/vlog"
)

// Bootstrap ...
type Bootstrap struct {

	//starter:component
	_as func(application.Lifecycle) //starter:as(".")

	AC     application.Context        //starter:inject("context")
	Boots  []testboot.BootingRegistry //starter:inject(".")
	Method string                     //starter:inject("${test.boot.method}")
	Path   string                     //starter:inject("${test.boot.path}")

}

// Life ...
func (inst *Bootstrap) Life() *application.Life {
	return &application.Life{
		OnLoop: inst.run,
	}
}

func (inst *Bootstrap) getMethodAndPath() (method string, path string) {
	params := inst.AC.GetParameters()
	method = params.GetParamOptional("method", "")
	path = params.GetParamOptional("path", "")
	if method != "" && path != "" {
		return
	}
	return inst.Method, inst.Path
}

func (inst *Bootstrap) run() error {
	method, path := inst.getMethodAndPath()
	key := inst.key(method, path)
	handlers, err := inst.handlers()
	if err != nil {
		return err
	}
	h := handlers[key]
	if h == nil {
		return fmt.Errorf("no testboot.Boot handler for [%s]", key)
	}

	vlog.Warn("do testboot.Boot handle [%s]", key)
	return h.Handler()
}

func (inst *Bootstrap) key(method, path string) string {
	return method + ":" + path
}

func (inst *Bootstrap) handlers() (map[string]*testboot.Boot, error) {
	src := inst.Boots
	dst := make(map[string]*testboot.Boot)
	tmp := make([]*testboot.Boot, 0)

	for _, reg := range src {
		list := reg.Boots()
		tmp = append(tmp, list...)
	}

	for _, b := range tmp {
		key := inst.key(b.Method, b.Path)
		older := dst[key]
		dst[key] = b
		if older != nil {
			return nil, fmt.Errorf("duplicate testboot.Boot handler for [%s]", key)
		}
	}

	return dst, nil
}
