package glob

import "unicode/utf8"

// The character which is treated like a glob
const GLOB = "*"

// The character which is treated like a single-character wildcard
const SINGLE = "?"

// Glob will test a string pattern, potentially containing globs, against a
// subject string. The result is a simple true/false, determining whether or
// not the glob pattern matched the subject text.
// The '*' character matches zero or more characters.
// The '?' character matches exactly one character (rune).
func Glob(pattern, subj string) bool {
	// Empty pattern can only match empty subject
	if pattern == "" {
		return subj == ""
	}

	// If the pattern is a single glob, it matches everything
	if pattern == GLOB {
		return true
	}

	// Fast path for exact match
	if pattern == subj {
		return true
	}

	var (
		pIdx    int
		sIdx    int
		starIdx = -1
		sTmpIdx = -1
		pLen    = len(pattern)
		sLen    = len(subj)
	)

	for sIdx < sLen {
		if pIdx < pLen {
			pByte := pattern[pIdx]
			if pByte == '*' {
				for pIdx < pLen && pattern[pIdx] == '*' {
					pIdx++
				}
				if pIdx == pLen {
					return true
				}
				starIdx = pIdx - 1
				sTmpIdx = sIdx
				continue
			}

			if pByte == '?' {
				_, sSize := utf8.DecodeRuneInString(subj[sIdx:])
				pIdx++
				sIdx += sSize
				continue
			}

			if pByte < utf8.RuneSelf && subj[sIdx] < utf8.RuneSelf {
				if pByte == subj[sIdx] {
					pIdx++
					sIdx++
					continue
				}
			} else {
				pRune, pSize := utf8.DecodeRuneInString(pattern[pIdx:])
				sRune, sSize := utf8.DecodeRuneInString(subj[sIdx:])
				if pRune == sRune {
					pIdx += pSize
					sIdx += sSize
					continue
				}
			}
		}

		if starIdx != -1 {
			_, sSize := utf8.DecodeRuneInString(subj[sTmpIdx:])
			sTmpIdx += sSize
			sIdx = sTmpIdx
			pIdx = starIdx + 1
			continue
		}

		return false
	}

	for pIdx < pLen && pattern[pIdx] == '*' {
		pIdx++
	}

	return pIdx == pLen
}
