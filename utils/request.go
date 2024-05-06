package utils

import (
	"net/http"
	"strings"
)

func GetLanguageFromRequest(req *http.Request) string {
	lang := req.Header.Get("Accept-Language")
	lang = strings.ToLower(lang)
	if lang == "en" || lang == "eng" || lang == "us" || lang == "english" {
		return "en"
	} else {
		return "en"
	}
}
