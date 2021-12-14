package main

import (
	"fmt"
	"os"
	"path"
	"sync"
	"testing"
)

func TestJpegToPdfConversion(t *testing.T) {

	pwd := os.Getenv("PWD")
	testImagePath := path.Join(pwd, "files", "testfile.jpg")

	outputPath, err := os.MkdirTemp(os.TempDir(), "testjpgtopdfconversion")
	if err != nil {
		t.Error("Error creating temp directory: ", err)
		t.Fail()
	}

	defer os.RemoveAll(outputPath)

	var wg sync.WaitGroup
	limiterChannel := make(chan Anonymous, 1)
	wg.Add(1)
	go ocr(&wg, testImagePath, outputPath, limiterChannel)
	wg.Wait()

	//Check if pdf exists
	fmt.Println("output path is :", outputPath)
	if files, err := os.ReadDir(outputPath); err != nil {
		t.Fail()
	} else {
		if !(len(files) > 0 && files[0].Name() == "testfile.pdf") {
			t.Log("length of files was more than 1 or file was not found")
			t.Fail()
		} else {
			joinPDF(outputPath, path.Join(pwd, "output.pdf"))
		}
	}
}
