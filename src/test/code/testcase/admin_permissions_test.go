package testcase

import (
	"net/http"
	"testing"
)

func TestAdminPermissionsInsert(t *testing.T) {
	tcl := &testingContextLoader{}
	tcl.run(http.MethodPost, "/api/v1/permissions")
}

func TestAdminPermissionsGetAll(t *testing.T) {
	tcl := &testingContextLoader{}
	tcl.run(http.MethodGet, "/api/v1/permissions")
}

func TestAdminPermissionsCRUD(t *testing.T) {
	tcl := &testingContextLoader{}
	tcl.run(http.MethodPost, "/api/v1/permissions/crud")
}
