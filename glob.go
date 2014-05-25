package main

import (
	"fmt"
	"strings"
)

const GLOB = "*"

func Glob(pattern, subj string) bool {
	parts := strings.Split(pattern, GLOB)

	switch len(parts) {
	case 0:
		return false
	case 1:
		return subj == pattern
	default:
		leadingGlob := strings.HasPrefix(pattern, GLOB)
		trailingGlob := strings.HasSuffix(pattern, GLOB)
		end := len(parts) - 1

		for i, part := range parts {
			switch {
			case i == 0 && leadingGlob:
				end--
				continue
			case i == 1 && leadingGlob:
				if !strings.Contains(subj, part) {
					return false
				}
			case i == end:
				if len(subj) > 0 {
					return trailingGlob || strings.HasSuffix(subj, part)
				}
				return true
			default:
				if !strings.Contains(subj, part) {
					return false
				}
			}
			idx := strings.Index(subj, part) + len(part)
			subj = subj[idx:]
		}
	}

	return true
}

func main() {
	fmt.Printf("%#v\n", Glob("this is*long* test*woo", "this is a long test-sentence with many globswoo"))
}
