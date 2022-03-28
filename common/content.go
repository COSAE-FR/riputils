package common

import (
	"io/ioutil"
	"os"
)

func FilePathToContent(filePath string) string {
	if len(filePath) == 0 {
		return filePath
	}
	if stat, err := os.Stat(filePath); err == nil {
		if !stat.IsDir() {
			content, err := ioutil.ReadFile(filePath)
			if err == nil && len(content) > 0 {
				return string(content)
			}
		}
	}
	return filePath
}
