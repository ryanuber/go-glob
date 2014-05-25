String globbing in golang
=========================

`go-glob` is a single-function library implementing basic string glob support.

Globs are an extremely user-friendly way of supporting string matching without
requiring knowledge of regular expressions or Go's particular regex engine. Most
people understand that if you put a `*` character somewhere in a string, it is
treated as a wildcard. Surprisingly, this functionality isn't found in Go's
standard library, except for `filepath.Match`, which is intended to be used
while comparing paths (not arbitrary strings), and works on a directory level
instead of a full string level. Globbing in a generic way is useful for
including lots of files under a common directory in any number of subdirectories
in a search, matching files by extension recursively in a directory tree,
matching a portion of text from user input, among many other things.

Example
=======

```
package main

import "github.com/ryanuber/go-glob"

func main() {
    glob.Glob("*World!", "Hello, World!") // true
    glob.Glob("Hello,*", "Hello, World!") // true
    glob.Glob("*ello,*", "Hello, World!") // true
    glob.Glob("World!", "Hello, World!")  // false
    glob.Glob("/home/*", "/home/ryanuber/.bashrc") // true
}
```
