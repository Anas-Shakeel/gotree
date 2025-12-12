package tree

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/anas-shakeel/gotree/internal/config"
	"github.com/anas-shakeel/gotree/internal/utils"
)

// PrintTree calls the internal traversing code (function)
//
// Acting as a middleman between cli and traversing logic.
func PrintTree(directories []string, config *config.Config) {
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
func traverse(root, prefix string, config *config.Config) error {
	// Read root, get all entries
	entries, err := os.ReadDir(root)
	if err != nil {
		return errors.New("error opening directory")
	}

	// Filter out entries
	filtered := utils.FilterEntries(&entries, config)

	// Iterate entries and print recursively
	for i, entry := range *filtered {
		// Create absolute path to entry
		entryName := entry.Name()
		absolutePath := filepath.Join(root, entryName)
		isLastEntry := i == len(*filtered)-1
		isDir := entry.IsDir()

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
