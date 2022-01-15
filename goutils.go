package goutils

import (
	"regexp"
)

func Hello() string {
	return ("Hello from Om Ganesh !!!")
}

// IsAlphanumeric checks if the string contains only letters and numbers.
// Empty string is valid.
func IsAlphanumeric(str string) bool {

	re := regexp.MustCompile("^[a-zA-Z0-9_]*$")
	if IsNull(str) {
		return true
	}
	return re.MatchString(str)
}

// IsNull checks if the string is null.
func IsNull(str string) bool {
	return len(str) == 0
}
