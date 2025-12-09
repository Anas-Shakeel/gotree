// Gotree implements the tree command that displays directory structures
// in a hierarchical format.
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(".")
	printTree(".", "")
}

// printTree prints directory structure in a tree-like format
func printTree(root, prefix string) {
	// Read root, get all entries
	entries, err := os.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate entries and print recursively
	for i, entry := range entries {
		// Create absolute path to entry
		absolutePath := filepath.Join(root, entry.Name())

		if shouldSkip(entry.Name()) {
			continue
		}

		// Check if last entry
		if i == len(entries)-1 {
			fmt.Printf("%s└── %s\n", prefix, entry.Name())

			if entry.IsDir() { // Directory?
				printTree(absolutePath, prefix+"    ")
			}

		} else {
			fmt.Printf("%s├── %s\n", prefix, entry.Name())

			if entry.IsDir() { // Directory?
				printTree(absolutePath, prefix+"│    ")
			}
		}
	}

}

// Returns true if entry should be skipped, false otherwise
func shouldSkip(entry string) bool {
	// Skip hidden files?
	if strings.HasPrefix(entry, ".") {
		return true
	}

	return false
}
