// Gotree implements the tree command that displays directory structures
// in a hierarchical format.
package main

import (
	"flag"
	"fmt"

	"github.com/anas-shakeel/gotree/internal/tree"
)

func main() {
	// Command line arguments/flags
	allFiles := flag.Bool("a", false, "show all files")

	// Set custon usage message
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), getCustomUsage())
		flag.PrintDefaults()
	}

	// Parse command-line arguments
	flag.Parse()
	pArgs := flag.Args()

	var directories []string // Multiple directories to print
	if len(pArgs) > 0 {
		directories = append(directories, pArgs...)
	} else {
		directories = append(directories, ".")
	}

	// Configurations for tree output
	config := tree.Config{
		ShowHiddenFiles: *allFiles,
	}

	tree.PrintTree(directories, &config)
}

// Returns the custom usage text to be used in help message.
func getCustomUsage() string {
	return `Usage: gotree [OPTIONS] <path(s)...>

Options:`
}
