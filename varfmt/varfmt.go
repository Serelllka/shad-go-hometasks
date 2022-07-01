package varfmt

import (
	"fmt"
)

func Sprintf(format string, args ...interface{}) (answer string) {
	fmt.Println(len(args))
	runes := []rune(format)
	var ans []rune
	f := true

	for _, letter := range runes {
		switch letter {
		case '{':
			f = false
			continue
		case '}':
			f = true
			continue
		}
		if f {
			ans = append(ans, letter)
			continue
		}
	}
	return ""
}
