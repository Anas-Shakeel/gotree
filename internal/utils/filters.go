package utils

import (
	"os"
	"strings"

	"github.com/anas-shakeel/gotree/internal/config"
)

// Filters out entries based on config. Returns the filtered entries
func FilterEntries(entries *[]os.DirEntry, config *config.Config) *[]os.DirEntry {
	// Filter out entries
	filtered := make([]os.DirEntry, 0, len(*entries))
	for _, entry := range *entries {
		// Show hidden files?
		if !config.ShowHiddenFiles && strings.HasPrefix(entry.Name(), ".") {
			continue // No!
		}

		// Show directories only?
		if config.DirsOnly && !entry.IsDir() {
			continue // No!
		}

		filtered = append(filtered, entry)
	}

	return &filtered
}
