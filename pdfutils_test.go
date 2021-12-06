package main

import (
	"os"
	"testing"
)

func TestSplitPdf(t *testing.T) {

	fileName := "testfile.pdf"
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
	err = splitPdf("testfile.pdf", tempDir)
	if err != nil {
		t.Fail()
	}

	t.Cleanup(func() {
		os.RemoveAll(tempDir)
	})
}
