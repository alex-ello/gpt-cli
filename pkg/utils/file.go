package utils

import (
	"os"
	"strings"
)

func GetFileList(dir string, separator string, max int) (string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	var fileList []string
	i := 0
	for _, file := range files {
		fileList = append(fileList, file.Name())
		if i > max {
			break
		}
		i++
	}

	return strings.Join(fileList, separator), nil
}
