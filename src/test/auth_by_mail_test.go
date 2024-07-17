package main

import (
	"testing"

	"github.com/starter-go/security-gorm/src/test/code/cases"
	"github.com/starter-go/security-gorm/src/test/code/testcom"
	"github.com/starter-go/vlog"
)

func TestHelpByMail(t *testing.T) {

	obj := &testcom.TestAuthWithEmail{}
	vlog.Debug("ref", obj.Units(nil))
	testWithCaseName(cases.HelpByMail, t)
}

func TestSendcodeByMail(t *testing.T) {

	obj := &testcom.TestAuthWithEmail{}
	vlog.Debug("ref", obj.Units(nil))
	testWithCaseName(cases.SendcodeByMail, t)
}

func TestLoginByMail(t *testing.T) {

	obj := &testcom.TestAuthWithEmail{}
	vlog.Debug("ref", obj.Units(nil))
	testWithCaseName(cases.LoginByMail, t)
}
