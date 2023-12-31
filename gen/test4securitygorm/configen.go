package test4securitygorm

import "github.com/starter-go/application"

//starter:configen(version="4")

// ComForSecurityGormTest ...
func ComForSecurityGormTest(cr application.ComponentRegistry) error {
	return registerComponents(cr)
	// return nil
}
