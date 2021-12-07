package main

import (
	"os"
	"testing"
)

func TestSplitPdf(t *testing.T) {

	fileName := "sample.pdf"
	fileNameWithoutExtension, err := getFileNameWithoutExtension(fileName)
	if err != nil {
		t.Errorf("Error getting file without extension, %v", err)
		t.Fail()
	}
	tempDir, err := createTempDir(fileNameWithoutExtension)
	if err != nil {
		t.Errorf("Error creating tempdir %v", err)
		t.Fail()
	}
	err = splitPdf(fileName, tempDir)
	if err != nil {
		t.Fail()
	}

	err = os.RemoveAll(tempDir)
	if err != nil {
		t.Log("Error deleting the folder : ", err)
	}
}
