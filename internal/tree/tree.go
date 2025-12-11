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
	PrefixPath      bool // Prefix path for each file?
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
		entryName := entry.Name()
		absolutePath := filepath.Join(root, entryName)
		isLastEntry := i == len(entries)-1
		isDir := entry.IsDir()

		// Show/Skip hidden files
		if !config.ShowHiddenFiles && strings.HasPrefix(entryName, ".") {
			continue
		}

		// Prefix path for each file?
		if config.PrefixPath {
			entryName = absolutePath
		}

		// Print entry
		if isLastEntry {
			fmt.Printf("%s└── %s\n", prefix, entryName)
		} else {
			fmt.Printf("%s├── %s\n", prefix, entryName)
		}

		// Entry a Directory?
		if isDir {
			newPrefix := prefix
			if isLastEntry {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}

			err = traverse(absolutePath, newPrefix, config)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
