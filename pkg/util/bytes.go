package util

import "fmt"

func BytesLiteral(b []byte) string {
	s := fmt.Sprintf("%#v", b)[7:]
	return s[:len(s)-1]
}
