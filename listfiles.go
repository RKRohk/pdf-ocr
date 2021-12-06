package main

import (
	"fmt"
	"io/ioutil"
)

func listFiles() []string {
	fileInfo, _ := ioutil.ReadDir("HRM_MEGA_GOD MODE_compressed")

	fileNames := make([]string, len(fileInfo))

	for index, file := range fileInfo {
		fileNames[index] = fmt.Sprintf("HRM_MEGA_GOD MODE_compressed/%s", file.Name())
	}

	return fileNames
}
