// This package defines a Counter type to count number of directories and files
package counter

import "fmt"

type Counter struct {
	Dirs  int
	Files int
}

// Register registers an entry based on isDir param.
//
// If isDir is true, it increments Directories... and files for false
func (c *Counter) Register(isDir bool) {
	if isDir {
		c.Dirs++
	} else {
		c.Files++
	}
}

// Output outputs dir/file count as a formatted string.
func (c *Counter) Output() string {
	return fmt.Sprintf("%d directories, %d files", c.Dirs, c.Files)
}
