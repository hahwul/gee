package gee

import "regexp"

// Regex is regex
func Regex(pattern, str string) bool {
	match, _ := regexp.MatchString(pattern, str)
	return match
}

// RegexV is RegexV
func RegexV(pattern, str string) bool {
	match, _ := regexp.MatchString(pattern, str)
	return !match
}
