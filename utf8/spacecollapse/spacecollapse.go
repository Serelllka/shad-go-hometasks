package spacecollapse

import "strings"

func CollapseSpaces(input string) string {
	runes := []rune(input)
	ind := 0
	s, w := false, true
	for i := 0; i < len(runes); i++ {
		if strings.Contains("\t\n\r", string(runes[i])) {
			runes[i] = ' '
		}
		w = !(s && runes[i] == ' ')
		s = runes[i] == ' '
		if w {
			runes[ind] = runes[i]
			ind++
		}
	}
	return string(runes[:ind])
}
