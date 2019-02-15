package ip_validation

import (
	"strings"
	"strconv"
)

func Is_valid_ip(ip string) bool {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false
	}

	for _, part := range parts {
		if isIncorrect(part) {
			return false
		}
	}

	return true
}

func isIncorrect(s string) bool {
	num, e := strconv.Atoi(s)

	if e != nil {
		return true
	}

	switch len(s) {
	case 3:
		return num < 100 || num > 255
	case 2:
		return num < 10 || num > 99
	case 1:
		return num < 0 || num > 9
	default:
		return true
	}
}
