package rotate_kata

import (
	"strconv"
)

func MaxRot(n int64) int64 {
	s := []byte(strconv.FormatInt(n, 10))
	rot := len(s) - 1

	biggest := n

	for i := 0; i < rot; i++ {
		tmp := s[0:i]
		s := append(tmp, rotate(s[i:])...)
		if int := toInt(s); biggest < int {
			biggest = int
		}
	}

	return biggest
}

func rotate(b []byte) []byte {
	return append(b[1:], b[0])
}

func toInt(b []byte) int64 {
	i, _ := strconv.ParseInt(string(b), 10, 64)

	return i
}
