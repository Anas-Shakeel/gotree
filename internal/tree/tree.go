package tree

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/anas-shakeel/gotree/internal/utils"
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
		traverse(dir, "")
	}
}

// traverse recursively traverses and prints directory structure in a tree-like format
func traverse(root, prefix string) {
	// Read root, get all entries
	entries, err := os.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate entries and print recursively
	for i, entry := range entries {
		// Create absolute path to entry
		absolutePath := filepath.Join(root, entry.Name())

		if utils.ShouldSkip(entry.Name()) {
			continue
		}

		// Check if last entry
		if i == len(entries)-1 {
			fmt.Printf("%s└── %s\n", prefix, entry.Name())

			if entry.IsDir() { // Directory?
				traverse(absolutePath, prefix+"    ")
			}

		} else {
			fmt.Printf("%s├── %s\n", prefix, entry.Name())

			if entry.IsDir() { // Directory?
				traverse(absolutePath, prefix+"│    ")
			}
		}
	}

}
