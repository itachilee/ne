package utils

import "regexp"

func CompressStr(str string) string {
	if str == "" {
		return ""
	}
	return regexp.MustCompile(`\\s+`).ReplaceAllString(str, "")
}
