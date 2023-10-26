package testcase

import (
	"net/http"
	"testing"
)

func TestAdminUsersInsert(t *testing.T) {
	tcl := &testingContextLoader{}
	tcl.run(http.MethodPost, "/api/v1/users")
}

func TestAdminUsersGetAll(t *testing.T) {
	tcl := &testingContextLoader{}
	tcl.run(http.MethodGet, "/api/v1/users")
}

func TestAdminUsersCRUD(t *testing.T) {
	tcl := &testingContextLoader{}
	tcl.run(http.MethodPost, "/api/v1/users/crud")
}
