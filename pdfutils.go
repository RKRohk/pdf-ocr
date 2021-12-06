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

	if a := recover(); a != nil {
		fmt.Println("Recover : ", a)
		return fmt.Errorf("error processing file")
	}

	var wg sync.WaitGroup
	for i := 0; i < doc.NumPage(); i++ {

		wg.Add(1)
		go func(i int) error {
			<-channel
			fmt.Println("Processing file number : ", i)
			img, err := doc.Image(i)
			if err != nil {
				channel <- 0
				return err
			}
			f, err := os.Create(filepath.Join(tempDir, fmt.Sprintf("%s:%03d.jpg", fileNameWithoutExtension, i)))
			if err != nil {
				channel <- 0
				return err
			}
			err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
			if err != nil {
				channel <- 0
				return err
			}
			f.Close()
			channel <- 0
			wg.Done()
			return nil
		}(i)

	}
	wg.Wait()
	return nil

}
