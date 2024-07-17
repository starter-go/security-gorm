package main

import (
	"testing"

	"github.com/starter-go/security-gorm/src/test/code/cases"
)

func TestHelpBySMS(t *testing.T) {
	testWithCaseName(cases.HelpBySMS, t)
}

func TestLoginBySMS(t *testing.T) {
	testWithCaseName(cases.LoginBySMS, t)
}

func TestSendcodeBySMS(t *testing.T) {
	testWithCaseName(cases.SendcodeBySMS, t)
}
