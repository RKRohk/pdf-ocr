build:
	CGO=0 go build -o pdf-ocr

run:
	./pdf-ocr

all: build run