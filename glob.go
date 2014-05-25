package main

import (
	"fmt"
	"strings"
)

func Glob(pattern, subj string) bool {
	parts := strings.Split(pattern, "*")

	switch len(parts) {
	case 0:
		return false
	case 1:
		return subj == pattern
	default:
		leadingGlob := strings.HasPrefix(pattern, "*")
		trailingGlob := strings.HasSuffix(pattern, "*")
		end := len(parts) - 1

		for i, part := range parts {
			switch {
			case i == 0 && leadingGlob:
				continue
			case i == 1 && leadingGlob:
				if !strings.Contains(subj, part) {
					return false
				}
			case i == end:
				if len(subj) > 0 {
					return trailingGlob
				}
				return true
			default:
				if !strings.HasPrefix(subj, part) {
					return false
				}
			}
			idx := strings.Index(subj, part) + len(part) + 1
			subj = subj[idx:]
		}
	}

	return true
}

func main() {
	fmt.Printf("%#v\n", Glob("*e*t*", "test123"))
}
