package common

import (
	"os"
	"path/filepath"
	"strings"
)

func GetJsonFilesFromPath(paths []string) []string {
	fileList := []string{}
	for _, path := range paths {
		filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
			if strings.HasSuffix(path, ".json") {
				fileList = append(fileList, path)
			}
			return nil
		})
	}
	return fileList
}
