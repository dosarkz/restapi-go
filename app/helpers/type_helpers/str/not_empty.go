package str

import "strings"

func NotEmpty(s string) bool {
	if strings.TrimSpace(s) != "" {
		return true
	}
	return false
}
