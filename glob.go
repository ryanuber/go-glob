package glob

import "strings"

// The character which is treated like a glob
const GLOB = "*"

// Compile will prepare a string pattern and return a function for further matching.
func Compile(pattern string) func(subj string) bool {
	// Empty pattern can only match empty subject
	if pattern == "" {
		return func(subj string) bool { return subj == "" }
	}

	// If the pattern _is_ a glob, it matches everything
	if pattern == GLOB {
		return func(subj string) bool { return true }
	}

	parts := strings.Split(pattern, GLOB)

	if len(parts) == 1 {
		// No globs in pattern, so test for equality
		return func(subj string) bool { return subj == pattern }
	}

	leadingGlob := strings.HasPrefix(pattern, GLOB)
	trailingGlob := strings.HasSuffix(pattern, GLOB)
	end := len(parts) - 1

	return func(subj string) bool {
		// Go over the leading parts and ensure they match.
		for i := 0; i < end; i++ {
			idx := strings.Index(subj, parts[i])

			switch i {
			case 0:
				// Check the first section. Requires special handling.
				if !leadingGlob && idx != 0 {
					return false
				}
			default:
				// Check that the middle parts match.
				if idx < 0 {
					return false
				}
			}

			// Trim evaluated text from subj as we loop over the pattern.
			subj = subj[idx+len(parts[i]):]
		}

		// Reached the last section. Requires special handling.
		return trailingGlob || strings.HasSuffix(subj, parts[end])
	}
}

// Glob will test a string pattern, potentially containing globs, against a
// subject string. The result is a simple true/false, determining whether or
// not the glob pattern matched the subject text.
func Glob(pattern, subj string) bool {
	return Compile(pattern)(subj)
}
