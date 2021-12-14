package main

import (
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"
)

const simultaneousFiles = 3

var configFilePath = path.Join(os.Getenv("PWD"), "configfile")

type Anonymous = struct{}

var limiterChannel = make(chan struct{}, simultaneousFiles)

func init() {
	for range make([]struct{}, simultaneousFiles) {
		limiterChannel <- struct{}{}
	}
}

//ocr performs ocr and returns an image with an invisible ocr layer for each file given to it
//
//inputFilePath should be the absolute path to the input jpeg file
//
//outputPath should be the path to the folder where the OCR'd pdf should be stored
func ocr(wg *sync.WaitGroup, inputFilePath string, outputPath string, channel chan Anonymous) {

	defer wg.Done()
	defer func() {
		channel <- Anonymous{}
	}()

	fileNameWithoutFolder := strings.Split(inputFilePath, "/")[len(strings.Split(inputFilePath, "/"))-1]
	fileNameWithoutExtension, _ := getFileNameWithoutExtension(fileNameWithoutFolder)

	cmd := exec.Command("tesseract", inputFilePath, outputPath+"/"+fileNameWithoutExtension, configFilePath)
	log.Println(cmd)
	err := cmd.Run()
	if err != nil {
		log.Println("error while performing ocr on " + fileNameWithoutFolder + " : " + err.Error())
		return
	}
}
