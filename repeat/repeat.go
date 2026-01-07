package repeat

import (
	"strings"
)

func Repeat(char string) string {
	var characters strings.Builder
	for i := 0; i < 5; i++ {
		characters.WriteString(char)
	}
	return characters.String()
}
