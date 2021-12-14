package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gen2brain/go-fitz"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

//Has shown to give good results
const simultaneousSplit = 10

func getFileNameWithoutExtension(fileName string) (string, error) {
	if len(strings.Split(fileName, ".")) < 2 {
		return "", fmt.Errorf("invalid filename")
	}
	fileNameWithoutExtension := strings.Split(fileName, ".")[0]
	return fileNameWithoutExtension, nil
}

func createTempDir(fileName string) (string, error) {
	tempDir, err := os.MkdirTemp(os.TempDir(), fileName)
	return tempDir, err
}

//Splits the pdf into multiple pdfs
//
//fileName: the absolute path to the pdf file
//
//tempDir: the absolute path to the temp directory
func splitPdf(fileName string, tempDir string) error {
	fileNameWithoutExtension, _ := getFileNameWithoutExtension(fileName)
	doc, err := fitz.New(fileName)
	if err != nil {
		return err
	}
	defer doc.Close()
	if err != nil {
		return err
	}

	log.Println("temp dir is : ", tempDir)

	channel := make(chan int, simultaneousSplit)
	for c := 0; c < simultaneousSplit; c++ {
		channel <- 0
	}

	var wg sync.WaitGroup
	for i := 0; i < doc.NumPage(); i++ {

		wg.Add(1)
		go convertPageToJPEG(channel, i, doc, tempDir, fileNameWithoutExtension, &wg)

	}

	log.Println("Waiting for all goroutines to finish")
	wg.Wait()

	return nil

}

//Converts a page to jpeg and saves it to the temp directory
func convertPageToJPEG(channel chan int, i int, doc *fitz.Document, tempDir string, fileNameWithoutExtension string, wg *sync.WaitGroup) error {
	<-channel

	//Ensure workgroup is marked as done even in a panic
	defer func() {
		wg.Done()
	}()

	//Free up channel
	defer func() {
		channel <- 0
	}()

	log.Println("Processing file number : ", i)
	img, err := doc.Image(i)
	if err != nil {
		log.Println("Error :", err)
	}
	f, err := os.Create(filepath.Join(tempDir, fmt.Sprintf("%s:%03d.jpg", fileNameWithoutExtension, i)))
	if err != nil {
		log.Println("Error :", err)
	}
	defer f.Close()

	err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	if err != nil {
		log.Println("Error :", err)
	}
	return nil
}

//Joins the pdfs in the temp directory into one pdf
//
//outputDir : the absolute path to the output directory
//
//outputFile: absolute path to the output file
func joinPDF(outputDir string, outputFile string) {
	files, err := os.ReadDir(outputDir)

	if err != nil {
		log.Panic("Error reading files from temp directory: ", err)
	}

	configuration := api.LoadConfiguration()

	filesWithFullPath := make([]string, len(files))
	for index, fileName := range files {
		filesWithFullPath[index] = filepath.Join(outputDir, fileName.Name())
	}

	err = api.MergeAppendFile(filesWithFullPath, outputFile, configuration)
	if err != nil {
		log.Println("Error merging files : ", err)
	}
}
