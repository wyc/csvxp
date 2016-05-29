package main

import "strings"

func quote(s string) string {
	return strings.Replace(s, `"`, `\"`, -1)
}

func unquote(s string) string {
	return strings.Replace(s, `\"`, `"`, -1)
}
