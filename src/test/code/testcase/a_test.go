package testcase

import (
	"sort"
	"testing"

	"github.com/starter-go/security-gorm/src/test/code/testboot"
	"github.com/starter-go/security-gorm/src/test/code/testcom"
)

func Test1(t *testing.T) {
	p := &testcom.BootParams{}
	p.Callback = func(b *testcom.Bootstrap) {
		src := b.AC.ListComponentIDs()
		ids := make([]string, 0)
		for _, id := range src {
			ids = append(ids, id.String())
		}
		sort.Strings(ids)
		for i, id := range ids {
			t.Logf(" component[%d].id:%s", i, id)
		}
	}
	testboot.Boot(p)
}
