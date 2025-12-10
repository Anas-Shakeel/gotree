package utils

import (
	"strings"
)

// Decides whether to skip entry or not!
//
// TODO: how does it decide??? explain please.
func ShouldSkip(entry string) bool {
	// Skip hidden files?
	if strings.HasPrefix(entry, ".") {
		return true
	}

	return false
}
