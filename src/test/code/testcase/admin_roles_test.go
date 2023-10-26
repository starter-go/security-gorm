package testcase

import (
	"net/http"
	"testing"
)

func TestAdminRolesGetList(t *testing.T) {
	tcl := &testingContextLoader{}
	tcl.run(http.MethodGet, "/api/v1/roles")
}

func TestAdminRolesInsert(t *testing.T) {
	tcl := &testingContextLoader{}
	tcl.run(http.MethodPost, "/api/v1/roles")
}

func TestAdminRolesCRUD(t *testing.T) {
	tcl := &testingContextLoader{}
	tcl.run(http.MethodPost, "/api/v1/roles/crud")
}
