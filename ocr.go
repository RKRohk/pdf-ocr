package main

import (
	"fmt"
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

func performOCR(inputFilePath string, filename string, id string) {
	var wg sync.WaitGroup

	outputDir, err := createTempDir(filename) //TODO(): handle error
	if err != nil {
		log.Println("error creating temp dir for file " + inputFilePath + " : " + err.Error())
	}

	//split pdf
	err = splitPdf(inputFilePath, outputDir)

	if err != nil {
		log.Println("error splitting pdf " + inputFilePath + " : " + err.Error())
	}
	//do ocr on every page

	channel := db[id]

	counter := 0

	pages, _ := os.ReadDir(outputDir)
	for _, page := range pages {
		wg.Add(1)
		<-limiterChannel
		log.Println("Performing ocr on page", page.Name())
		log.Println(page)
		channel <- fmt.Sprintf("Processing page %d of %d", counter, len(pages))
		counter++
		go ocr(&wg, path.Join(outputDir, page.Name()), outputDir, limiterChannel)
	}

	wg.Wait()
	//merge

	joinPDF(outputDir, path.Join(os.Getenv("PWD"), "output", fmt.Sprintf("%s.pdf", id)))

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
