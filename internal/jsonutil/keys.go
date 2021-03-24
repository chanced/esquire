package jsonutil

import "strings"

// TODO: either modify sjson & gjson or come up with a different approach to setting keys

func EscapeKey(key string) string {
	return strings.ReplaceAll(strings.ReplaceAll(key, `.`, `\.`), `:`, `\:`)
}
