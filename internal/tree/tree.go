package tree

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/anas-shakeel/gotree/internal/config"
	"github.com/anas-shakeel/gotree/internal/counter"
	"github.com/anas-shakeel/gotree/internal/utils"
)

// PrintTree calls the internal traversing code (function)
//
// Acting as a middleman between cli and traversing logic.
func PrintTree(directories []string, config *config.Config) {
	var count counter.Counter

	for _, dir := range directories {
		fmt.Println(dir)
		err := traverse(dir, "", config, &count)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Increment for root dir (only for successful ones)
		count.Dirs++
	}

	// Print count
	if config.DirsOnly {
		fmt.Printf("\n%d directories\n", count.Dirs)
	} else {
		fmt.Printf("\n%d directories, %d files\n", count.Dirs, count.Files)
	}
}

// traverse recursively traverses and prints directory
// structure in a tree-like format
func traverse(root, prefix string, config *config.Config, count *counter.Counter) error {
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

		// Count files
		if !isDir {
			count.Files++
		}

		// Entry a Directory?
		if isDir {
			count.Dirs++ // Count dirs
			newPrefix := prefix
			if isLastEntry {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}

			err = traverse(absolutePath, newPrefix, config, count)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
