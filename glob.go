package glob

import "unicode/utf8"

// The character which is treated like a glob (matches any sequence).
const GLOB = "*"

// single matches any single character (one Unicode code point).
const single = '?'

// Glob will test a string pattern, potentially containing globs, against a
// subject string. The result is a simple true/false, determining whether or
// not the glob pattern matched the subject text.
//
// '*' matches any sequence of characters (including empty).
// '?' matches any single character.
func Glob(pattern, subj string) bool {
	return match(pattern, subj)
}

func match(pattern, subj string) bool {
	for len(pattern) > 0 {
		switch pattern[0] {
		case '*':
			// Collapse consecutive *'s.
			for len(pattern) > 1 && pattern[1] == '*' {
				pattern = pattern[1:]
			}
			pattern = pattern[1:]
			if pattern == "" {
				return true
			}
			// Try the rest of the pattern against every suffix of subj.
			for i := 0; i <= len(subj); {
				if match(pattern, subj[i:]) {
					return true
				}
				if i == len(subj) {
					break
				}
				_, size := utf8.DecodeRuneInString(subj[i:])
				if size < 1 {
					size = 1
				}
				i += size
			}
			return false
		case '?':
			if subj == "" {
				return false
			}
			_, size := utf8.DecodeRuneInString(subj)
			if size < 1 {
				size = 1
			}
			pattern = pattern[1:]
			subj = subj[size:]
		default:
			if subj == "" || pattern[0] != subj[0] {
				return false
			}
			pattern = pattern[1:]
			subj = subj[1:]
		}
	}
	return subj == ""
}
