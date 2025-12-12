// This package defines the Config struct type
package config

// Config struct for options
type Config struct {
	ShowHiddenFiles bool // Show hidden files as well?
	PrefixPath      bool // Prefix path for each file?
	DirsOnly        bool // Show directories only?
}
