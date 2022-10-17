package main

import (
	"regexp"
	"strconv"
	"strings"
)

// Convert string to a number
func ToNumber(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

// Check whether string contains a substring (based on a regex)
func StringContains(regex string, src string) bool {
	return FindInText(regex, src) != ""
}

// Find substring in the text
func FindInText(regex string, s string) string {
	match := regexp.MustCompile("(?i)" + regex).FindString(s)
	return match
}

// Find all matches in the text
func FindAllMatchesInText(regex string, s string) []string {
	matches := regexp.MustCompile(regex).FindAllString(s, -1)

	if matches == nil {
		return []string{}
	} else {
		return matches
	}
}

// Trim spaces in a string
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// Replace the matched pattern in a string with a replacement
func ReplaceInText(regex string, src string, replacement string) string {
	return regexp.MustCompile(regex).ReplaceAllString(src, replacement)
}