package tree

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Config struct for options
type Config struct {
	ShowHiddenFiles bool // Show hidden files as well?
}

// PrintTree calls the internal traversing code (function)
//
// Acting as a middleman between cli and traversing logic.
func PrintTree(directories []string, config *Config) {
	for _, dir := range directories {
		fmt.Println(dir)
		err := traverse(dir, "", config)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// traverse recursively traverses and prints directory
// structure in a tree-like format
func traverse(root, prefix string, config *Config) error {
	// Read root, get all entries
	entries, err := os.ReadDir(root)
	if err != nil {
		return errors.New("error opening directory")
	}

	// Iterate entries and print recursively
	for i, entry := range entries {
		// Create absolute path to entry
		absolutePath := filepath.Join(root, entry.Name())
		isLastEntry := i == len(entries)-1 // ??

		// Show/Skip hidden files
		if !config.ShowHiddenFiles && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		newPrefix := prefix

		// Print entry
		if isLastEntry {
			fmt.Printf("%s└── %s\n", prefix, entry.Name())
			newPrefix += "    "
		} else {
			fmt.Printf("%s├── %s\n", prefix, entry.Name())
			newPrefix += "│   "
		}

		// Entry a Directory?
		if entry.IsDir() {
			err = traverse(absolutePath, newPrefix, config)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
