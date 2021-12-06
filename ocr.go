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

func ocr(wg *sync.WaitGroup, fileName string, channel chan struct{}) {
	fileNameWithoutFolder := strings.Split(fileName, "/")[len(strings.Split(fileName, "/"))-1]
	fileNameWithoutExtension := strings.Split(fileNameWithoutFolder, ".")[0]
	pwd := os.Getenv("PWD")
	sourceFile := pwd + "/" + fileName
	outputFile := pwd + "/" + "output" + "/" + fileNameWithoutExtension
	cmd := exec.Command("tesseract", sourceFile, outputFile, pwd+"/configfile")
	err := cmd.Run()
	if err != nil {
		fmt.Println("error while performing ocr on " + fileName + ": " + err.Error())
	}
	wg.Done()
	channel <- struct{}{}
}
