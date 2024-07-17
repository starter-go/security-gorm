package vericode

import (
	"strings"

	"github.com/starter-go/i18n"
	"github.com/starter-go/security/auth"
)

// GetLanguage ...
func GetLanguage(a auth.Authentication) []i18n.Language {
	dst := make([]i18n.Language, 0)
	str := a.Parameters().GetParam("language")
	src := strings.Split(str, ",")
	for _, item := range src {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		dst = append(dst, i18n.Language(item))
	}
	return dst
}
