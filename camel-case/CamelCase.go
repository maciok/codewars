package sum_of_digits

import (
	"strings"
)

var replacer = strings.NewReplacer("-", " ", "_", " ")

func ToCamelCase(s string) string {
	if s == "" {
		return ""
	}

	result := strings.Title(replacer.Replace(s))
	result = s[:1] + result[1:]
	result = strings.Replace(result, " ", "", -1)
	return result
}
