package testboot

////////////////////////////////////////////////////////////////////////////////

// Boot ...
type Boot struct {
	Method  string
	Path    string
	Handler func() error
}

////////////////////////////////////////////////////////////////////////////////

// BootingRegistry ...
type BootingRegistry interface {
	Boots() []*Boot
}

////////////////////////////////////////////////////////////////////////////////

// BootList ...
type BootList struct {
	items []*Boot
}

// Handle ...
func (inst *BootList) Handle(method string, path string, h func() error) {
	b := &Boot{
		Method:  method,
		Path:    path,
		Handler: h,
	}
	inst.items = append(inst.items, b)
}

// List ...
func (inst *BootList) List() []*Boot {
	return inst.items
}

////////////////////////////////////////////////////////////////////////////////
