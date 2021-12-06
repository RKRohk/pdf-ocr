package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

const simultaneousFiles = 3

var limiterChannel = make(chan struct{}, simultaneousFiles)

func init() {
	for range make([]struct{}, simultaneousFiles) {
		limiterChannel <- struct{}{}
	}
}

//ocr performs ocr and returns an image with an invisible ocr layer for each file given to it
func ocr(wg *sync.WaitGroup, filePath string, channel chan struct{}) {
	fileNameWithoutFolder := strings.Split(filePath, "/")[len(strings.Split(filePath, "/"))-1]
	fileNameWithoutExtension := strings.Split(fileNameWithoutFolder, ".")[0]
	pwd := os.Getenv("PWD")
	sourceFile := filePath
	outputFile := pwd + "/" + "output" + "/" + fileNameWithoutExtension
	cmd := exec.Command("tesseract", sourceFile, outputFile, pwd+"/configfile")
	err := cmd.Run()
	if err != nil {
		fmt.Println("error while performing ocr on " + fileNameWithoutFolder + ": " + err.Error())
	}
	wg.Done()
	channel <- struct{}{}
}
