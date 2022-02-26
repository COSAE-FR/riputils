package common

import (
	"os"
)

// IsDirectory determines if a file represented
// by `path` is a directory or not
func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}
