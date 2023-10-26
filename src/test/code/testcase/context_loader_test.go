package testcase

import (
	"net/http"
	"testing"
)

func TestContextLoader(t *testing.T) {
	tcl := &testingContextLoader{}
	tcl.run(http.MethodOptions, "///")
}
