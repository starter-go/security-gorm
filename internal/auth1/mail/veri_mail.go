package mail

////////////////////////////////////////////////////////////////////////////////

// type veriMailClass struct {
// 	class  keyvalues.Class
// 	maxAge time.Duration
// }

// func (inst *veriMailClass) init(ser keyvalues.Service, maxage lang.Milliseconds) error {
// 	const (
// 		ns    = "github.com/starter-go/security-gorm/internal/auth1/mail"
// 		alias = "VeriMail"
// 	)
// 	cl, err := ser.GetClassNS(ns, alias)
// 	if err != nil {
// 		return err
// 	}
// 	inst.maxAge = maxage.Duration()
// 	inst.class = cl
// 	return nil
// }

// func (inst *veriMailClass) put(addr string, item *VeriMail) error {

// 	if !strings.Contains(addr, "@") {
// 		return fmt.Errorf("bad email address")
// 	}

// 	if item == nil {
// 		return fmt.Errorf("item is nil")
// 	}

// 	opt := &keyvalues.Options{MaxAge: inst.maxAge}
// 	ent := inst.class.GetJSONEntry(addr)
// 	return ent.Put(item, opt)
// }

// func (inst *veriMailClass) get(addr string) (*VeriMail, error) {

// 	if !strings.Contains(addr, "@") {
// 		return nil, fmt.Errorf("bad email address")
// 	}

// 	item := &VeriMail{}
// 	ent := inst.class.GetJSONEntry(addr)
// 	err := ent.Get(item)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return item, nil
// }
