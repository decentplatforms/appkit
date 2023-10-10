package testhelp

import (
	"strings"
)

// MapOptions maps test option keys to values.
// Keys are strings separated by . symbols, denoting more
// specific options. When using GetTestOption(TestOptionMap[T], key), the most specific
// match is chosen, so a query for a.b.c will match a.b if a.b.c does not exist.
// If no key matches, returns def.
type TestOptionMap[T any] map[string]T

// GetTestOption gets an option from a map[string]T (TestOptionMap[T]).
// Keys are strings separated by . symbols, denoting more
// specific options. When using GetTestOption(TestOptionMap[T], key), the most specific
// match is chosen, so a query for a.b.c will match a.b if a.b.c does not exist.
// If no key matches, returns def.
func GetTestOption[T any](tom TestOptionMap[T], key string, def T) T {
	toks := strings.Split(key, ".")
	for i := len(toks); i > 0; i-- {
		subkey := strings.Join(toks[:i], ".")
		if v, ok := tom[subkey]; ok {
			return v
		}
	}
	return def
}
