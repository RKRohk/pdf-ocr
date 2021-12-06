package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gen2brain/go-fitz"
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
	if err != nil {
		return "", err
	}
	return tempDir, nil

}

func splitPdf(fileName string, tempDir string) error {
	fileNameWithoutExtension, _ := getFileNameWithoutExtension(fileName)
	outputDir := os.TempDir() + "/" + fileNameWithoutExtension
	os.Mkdir(outputDir, 0700)
	doc, err := fitz.New("testfile.pdf")
	if err != nil {
		return err
	}
	defer doc.Close()
	if err != nil {
		return err
	}

	fmt.Println("temp dir is : ", tempDir)

	channel := make(chan int, simultaneousSplit)
	for c := 0; c < simultaneousSplit; c++ {
		channel <- 0
	}

	var wg sync.WaitGroup
	for i := 0; i < doc.NumPage(); i++ {

		wg.Add(1)
		go convertPageToJPEG(channel, i, doc, tempDir, fileNameWithoutExtension, &wg)

	}

	fmt.Println("Waiting for all goroutines to finish")
	wg.Wait()

	return nil

}

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

	fmt.Println("Processing file number : ", i)
	img, err := doc.Image(i)
	if err != nil {
		fmt.Println("Error :", err)
	}
	f, err := os.Create(filepath.Join(tempDir, fmt.Sprintf("%s:%03d.jpg", fileNameWithoutExtension, i)))
	if err != nil {
		fmt.Println("Error :", err)
	}
	defer f.Close()

	err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	if err != nil {
		fmt.Println("Error :", err)
	}
	return nil
}
