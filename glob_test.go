package glob

import "testing"

func testGlobMatch(t *testing.T, pattern, subj string) {
	if !Glob(pattern, subj) {
		t.Fatalf("%s should match %s", pattern, subj)
	}
}

func testGlobNoMatch(t *testing.T, pattern, subj string) {
	if Glob(pattern, subj) {
		t.Fatalf("%s should not match %s", pattern, subj)
	}
}

func TestEmptyPattern(t *testing.T) {
	testGlobMatch(t, "", "")
	testGlobNoMatch(t, "", "test")
}

func TestPatternWithoutGlobs(t *testing.T) {
	testGlobMatch(t, "test", "test")
}

func TestGlob(t *testing.T) {
	for _, pattern := range []string{
		"*test",
		"this*",
		"*is *",
		"*is*a*",
	} {
		testGlobMatch(t, pattern, "this is a test")
	}

	for _, pattern := range []string{
		"test*",
		"*is",
		"*no*",
	} {
		testGlobNoMatch(t, pattern, "this is a test")
	}
}
