package ast

import "strings"

func printRaw(s string) string {
	return strings.ReplaceAll(s, "\n", "\\n")
}
