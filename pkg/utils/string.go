package utils

import (
	"strings"
)

func UnwrapString(str, prefix, suffix string) (string, bool) {
	if strings.HasPrefix(str, prefix) && strings.HasSuffix(str, suffix) {
		return strings.TrimPrefix(strings.TrimSuffix(str, suffix), prefix), true
	}
	return str, false
}
